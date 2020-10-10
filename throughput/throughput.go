// Package throughput provides an interface and implementations of that
// interface to read the throughput of network devices.
package throughput

import (
	"time"
)

// Reporter is the interface implemented by types that report basic network
// device throughput stats.
type Reporter interface {
	// Doc comment for Report
	Report(devices []string) []Stat
}

type Stat struct {
	Name    string // Name is blahblah
	BitsIn  int
	BitsOut int
	Elapsed time.Duration
}
