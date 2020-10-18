// Package throughput provides an interface and implementations of that
// interface to read the throughput of network devices.
//
// Synopsis
//
// A brief example of using the package to print throughput stats:
//   package main
//
//   import (
//     "github.com/robkingsbury/bndstat/throughput"
//     "time"
//   )
//
//   func main() {
//     reporter, _ := throughput.NewReporter()
//     table := throughput.NewTable()
//
//     for {
//       stats, _ := reporter.Report()
//       table.Write(stats, stats.Devices(), throughput.Kbps)
//       time.Sleep(time.Second)
//     }
//   }
// Note: throwing away errors is bad practice but is done here for brevity.
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

// NewReporter returns an initialized Reporter that's compatible with the
// current OS. An error is returned if the OS is not supported.
func NewReporter() (Reporter, error) {
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
