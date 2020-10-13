package throughput

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

// Linux implements the Reporter interface for linux systems.
type Linux struct {
	devices map[string]*deviceData
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
}

// singleRead is the struct representing the bitsIn and bitsOut of a device at
// one point in time.
type singleRead struct {
	name     string
	bytesIn  uint64
	bytesOut uint64
}

// NewLinux returns a pointer to an initialized Linux.
func NewLinux() *Linux {
	return &Linux{devices: map[string]*deviceData{}}
}

// Report reads /proc/net/dev, updates its internal state with the latest
// counters, and returns a slice of Stats for all network devices.
func (l *Linux) Report() (*Stats, error) {
	p, err := os.Open("/proc/net/dev")
	if err != nil {
		return &Stats{}, err
	}

	srs, err := l.parseNetDev(p)
	if err != nil {
		return &Stats{}, err
	}

	l.update(srs, time.Now())
	return l.stats(), nil
}

// update l.devices with info from a slice of singleReads.
func (l *Linux) update(srs []*singleRead, now time.Time) {
	for _, sr := range srs {
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
	}
}

// stats returns a pointer to Stats from the data in l.devices.
func (l *Linux) stats() *Stats {
	stats := &Stats{
		devices: map[string]*stat{},
	}

	for device, data := range l.devices {
		s := &stat{
			elapsed:  data.currentTime.Sub(data.lastTime),
			bytesIn:  data.currentBytesIn - data.lastBytesIn,
			bytesOut: data.currentBytesOut - data.lastBytesOut,
		}
		stats.devices[device] = s
	}

	return stats
}

// parseNetDev expects the input byte slice to match the format in
// /proc/net/dev. It returns a slice of singleReads, populating each element
// with each device found. If the input byte slice does not match the correct
// format, an empty slice will be returned.
func (l *Linux) parseNetDev(i io.Reader) ([]*singleRead, error) {
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
			}
			srs = append(srs, sr)
		}
	}

	if err := s.Err(); err != nil {
		return []*singleRead{}, err
	}

	return srs, nil
}
