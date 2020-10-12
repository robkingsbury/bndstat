// Package throughput provides an interface and implementations of that
// interface to read the throughput of network devices.
package throughput

import (
	"fmt"
	"runtime"
	"time"

	"github.com/golang/glog"
)

// Reporter is the interface implemented by types that report basic network
// device throughput stats.
type Reporter interface {
	// Report should return a slice of Stats for all discovered network devices.
	Report() ([]*Stat, error)
}

// Stat reports data on how much traffic has passed through network devices.
type Stat struct {
	// Name is device name.
	Name string

	// BytesIn is the count of inbound Bytes that passed through the interface
	// since the last innvocation of Report().
	BytesIn uint64

	// BytesOut is the count of outbound Bytes that passed through the interface
	// since the last innvocation of Report().
	BytesOut uint64

	// Elapsed is the amount of time that has elapsed since the last invocation of
	// Report().
	Elapsed time.Duration
}

// New returns an initialized Reporter that's compatible with the current OS.
// An error is returned if the OS is not supported.
func New() (Reporter, error) {
	glog.V(1).Infof("os is %q", runtime.GOOS)

	var r Reporter
	switch runtime.GOOS {
	case "linux":
		r = NewLinux()
	default:
		return &NilReporter{}, fmt.Errorf("os %q not supported", runtime.GOOS)
	}

	// Call Report() twice to initialize the data, replacing all default values,
	// but do not output anything.
	r.Report()
	r.Report()
	return r, nil
}
