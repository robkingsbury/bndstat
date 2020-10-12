// Package throughput provides an interface and implementations of that
// interface to read the throughput of network devices.
package throughput

import (
	"fmt"
	"runtime"

	"github.com/golang/glog"
)

// Reporter is the interface implemented by types that report basic network
// device throughput stats.
type Reporter interface {
	// Report should return a pointer to a Stats struct containing data for all
	// discovered network devices.
	Report() (*Stats, error)
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
