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
	Report() []Stat
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

// New returns an initialized Reporter that's compatible with the current OS.
// An error is returned if the OS is not supported.
func New() (Reporter, error) {
	glog.V(1).Infof("os is %q", runtime.GOOS)

	var r Reporter

	switch runtime.GOOS {
	case "linux":
		r = &Linux{}
	default:
		return &Linux{}, fmt.Errorf("os %q not supported", runtime.GOOS)
	}

	// Call Report() to initialize the data but do not output anything.
	r.Report()
	return r, nil
}
