package throughput

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/kr/pretty"
)

const netDevCounterSize float64 = 32

// Linux implements the Reporter interface for linux systems.
type Linux struct {
	devices     map[string]*deviceData
	counterSize float64
	maxCounter  uint64
}

// deviceData is the persistent data held in a Linux struct. When Report() is
// called, it updates deviceData for all devices and returns []Stats from that
// data.
type deviceData struct {
	lastTime     time.Time
	lastBytesIn  uint64
	lastBytesOut uint64

	currentTime     time.Time
	currentBytesIn  uint64
	currentBytesOut uint64

	rawText [shiftSize]string
}

// singleRead is the struct representing the bytesIn and bytesOut of a device at
// one point in time.
type singleRead struct {
	name     string
	bytesIn  uint64
	bytesOut uint64
	rawText  string
}

// NewLinux returns a pointer to an initialized Linux.
func NewLinux() *Linux {
	return &Linux{
		devices:     map[string]*deviceData{},
		counterSize: netDevCounterSize,
		maxCounter:  uint64(math.Pow(2, netDevCounterSize)),
	}
}

// Report reads /proc/net/dev, updates its internal state with the latest
// counters, and returns a slice of Stats for all network devices.
func (l *Linux) Report() (*Stats, error) {
	p, err := os.Open("/proc/net/dev")
	if err != nil {
		return &Stats{}, err
	}

	srs, err := parseNetDev(p)
	if err != nil {
		return &Stats{}, err
	}

	l.update(srs, time.Now())
	stats := l.stats()

	// Look for anything that should trigger a raw data dump to debug bad data rates.
	for _, device := range stats.Devices() {
		in, out, err := stats.Avg(device, Kbps)
		if err != nil {
			return &Stats{}, fmt.Errorf("could not get average from %s", device)
		}

		trigger := false
		switch {
		// 1 Tbps is faster than any hardware as of 2020.
		case in > 1000000000 || out > 1000000000:
			glog.Warningf("Triggering data dump because very large rate found")
			trigger = true
		case in < 0 || out < 0:
			glog.Warningf("Triggering data dump because negative rate found")
			trigger = true
		}

		if trigger {
			glog.Infof("  device=%s", device)
			glog.Infof("  in=%f", in)
			glog.Infof("  out=%f", out)
			glog.Infof("  stats:\n%s", pretty.Sprint(stats))
			l.dumpRawText(device)
			glog.Flush()
		}
	}

	return l.stats(), nil
}

// dumpRawText logs the raw text from /proc/net/dev for each device.
func (l *Linux) dumpRawText(device string) {
	glog.Infof("Dumping /prov/net/dev data for %s", device)

	data, ok := l.devices[device]
	if !ok {
		glog.Errorf("dumpRawText on device that does not exist: %s", device)
		return
	}

	for i, t := range data.rawText {
		glog.Infof("  [%d] %s", i, t)
	}
}

// update l.devices with info from a slice of singleReads.
func (l *Linux) update(srs []*singleRead, now time.Time) {
	for _, sr := range srs {
		glog.V(1).Infof("updating state for %s", sr.name)

		if _, exists := l.devices[sr.name]; !exists {
			l.devices[sr.name] = &deviceData{}
		}

		d := l.devices[sr.name]

		d.lastTime = d.currentTime
		d.lastBytesIn = d.currentBytesIn
		d.lastBytesOut = d.currentBytesOut

		d.currentTime = now
		d.currentBytesIn = sr.bytesIn
		d.currentBytesOut = sr.bytesOut

		d.shiftRawText()
		d.rawText[0] = sr.rawText
	}
}

// stats returns a pointer to Stats from the data in l.devices.
func (l *Linux) stats() *Stats {
	stats := &Stats{
		devices: map[string]*stat{},
	}

	for device, data := range l.devices {
		inDiff := data.currentBytesIn - data.lastBytesIn
		outDiff := data.currentBytesOut - data.lastBytesOut

		// Correct for counter rollover
		if data.currentBytesIn < data.lastBytesIn {
			glog.V(1).Infof("Counter rollover for %s (in)", device)
			glog.V(2).Infof("max = %d", l.maxCounter)
			inDiff = l.maxCounter - data.lastBytesIn + data.currentBytesIn
		}
		if data.currentBytesOut < data.lastBytesOut {
			glog.V(1).Infof("Counter rollover for %s (out)", device)
			glog.V(2).Infof("max = %d", l.maxCounter)
			outDiff = l.maxCounter - data.lastBytesOut + data.currentBytesOut
		}

		s := &stat{
			elapsed:  data.currentTime.Sub(data.lastTime),
			bytesIn:  inDiff,
			bytesOut: outDiff,
		}
		stats.devices[device] = s
	}

	return stats
}

// parseNetDev expects the input byte slice to match the format in
// /proc/net/dev. It returns a slice of singleReads, populating each element
// with each device found. If the input byte slice does not match the correct
// format, an empty slice will be returned.
func parseNetDev(i io.Reader) ([]*singleRead, error) {
	srs := []*singleRead{}

	s := bufio.NewScanner(i)
	for s.Scan() {
		fields := strings.Fields(strings.TrimSpace(s.Text()))
		if len(fields) == 17 {
			dev := strings.TrimSuffix(fields[0], ":")
			glog.V(2).Infof("found device %s", dev)

			bytesRecvStr := fields[1]
			bytesTransStr := fields[9]

			glog.V(2).Infof("bytesRecvStr for %s: %s", dev, bytesRecvStr)
			glog.V(2).Infof("bytesTransStr for %s: %s", dev, bytesTransStr)

			bytesRecv, err := strconv.ParseUint(bytesRecvStr, 10, 64)
			if err != nil {
				return []*singleRead{}, err
			}
			bytesTrans, err := strconv.ParseUint(bytesTransStr, 10, 64)
			if err != nil {
				return []*singleRead{}, err
			}

			sr := &singleRead{
				name:     dev,
				bytesIn:  uint64(bytesRecv),
				bytesOut: uint64(bytesTrans),
				rawText:  s.Text(),
			}
			srs = append(srs, sr)
		}
	}

	if err := s.Err(); err != nil {
		return []*singleRead{}, err
	}

	return srs, nil
}

const shiftSize = 4

// Shift raw data: 0 is latest, shiftSize is the earliest.
func (d *deviceData) shiftRawText() {
	for i := shiftSize - 1; i > 0; i-- {
		d.rawText[i] = d.rawText[i-1]
	}
}
