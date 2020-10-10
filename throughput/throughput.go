// Package throughput provides an interface and implementations of that
// interface to read the throughput of network devices.
package throughput

import (
	"time"
)

// Reporter is the interface implemented by types that report basic network
// device throughput stats.
type Reporter interface {
	// Report should return a slice of Stats for the devices listed in the given
	// slice of device names. If the input slice is empty, Stats for all
	// discovered devices should be returned.
	Report(devices []string) []Stat
}

// Stat reports data on how much traffic has passed through network devices.
type Stat struct {
	// Name is device name.
	Name string

	// BitsIn is the count of inbound bits that passed through the interface since
	// the last innvocation of Report().
	BitsIn int

	// BitsOut is the count of outbound bits that passed through the interface since
	// the last innvocation of Report().
	BitsOut int

	// Elapsed is the amount of time that has elapsed since the last invocation of
	// Report().
	Elapsed time.Duration
}
