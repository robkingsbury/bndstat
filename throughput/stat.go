package throughput

import (
	"math"
	"time"
)

// A unit is used to format the value returned by Stat.Avg().
type Unit int

const (
	Bps  Unit = 1
	Kbps Unit = 10
	Mbps Unit = 20
	Gbps Unit = 30
	Tbps Unit = 40
)

// String implements Stringer for a Unit.
func (u Unit) String() string {
	switch u {
	case Kbps:
		return "kbps"
	case Mbps:
		return "mbps"
	case Gbps:
		return "gbps"
	case Tbps:
		return "tbps"
	default:
		return "bps"
	}
}

// stat reports data on how much traffic has passed through a network device.
type stat struct {
	// bytesIn is the count of inbound Bytes that passed through the interface
	// since the last innvocation of Report().
	bytesIn uint64

	// bytesOut is the count of outbound Bytes that passed through the interface
	// since the last innvocation of Report().
	bytesOut uint64

	// elapsed is the amount of time that has elapsed since the last invocation of
	// Report().
	elapsed time.Duration
}

// Stats contains a Stat for each network device.
type Stats struct {
	devices map[string]*stat
}

// Devices returns a slice of device names that Stats has information on.
func (s *Stats) Devices() []string {
	devices := []string{}
	for k := range s.devices {
		devices = append(devices, k)
	}
	return devices
}

// Avg returns the average throughput for the device, in the units specified.
// If the device does not exist, zeros are returned.
func (s *Stats) Avg(device string, unit Unit) (in float64, out float64) {
	stat, ok := s.devices[device]
	if !ok {
		return 0, 0
	}

	div := math.Pow(2, float64(unit))
	in = (float64(stat.bytesIn) / div) / stat.elapsed.Seconds()
	out = (float64(stat.bytesOut) / div) / stat.elapsed.Seconds()
	return in, out
}
