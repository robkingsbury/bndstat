// Package throughput provides an interface and implementations of that
// interface to read the throughput of network devices.
package throughput

import (
	"time"
)

type Reporter interface {
	Report() []Stat
}

type Stat struct {
	Name    string
	BitsIn  int
	BitsOut int
	Elapsed time.Duration
}
