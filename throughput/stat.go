package throughput

import (
	"math"
	"time"
)

// stat reports data on how much traffic has passed through a network device.
//
// TODO: add methods to extract timestamp and elapsed time, maybe include both
//       start and end timestamps?
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

// Stats contains throughput data for each network device.
type Stats struct {
	devices map[string]*stat // All the stats.
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
//
// TODO: test needed, return error when no device found (don't be lazy)
func (s *Stats) Avg(device string, unit Unit) (in float64, out float64) {
	stat, ok := s.devices[device]
	if !ok {
		return 0, 0
	}

	div := math.Pow(2, float64(unit))
	in = (float64(stat.bytesIn*8) / div) / stat.elapsed.Seconds()
	out = (float64(stat.bytesOut*8) / div) / stat.elapsed.Seconds()
	return in, out
}
