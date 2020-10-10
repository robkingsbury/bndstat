// Bndstat displays simple network device throughput data, inspired by unix
// tools such as vmstat, iostat, mpstat, netstat, etc.
//
// A note about logging: bndstat uses the glog facility for logging but, under
// default conditions i.e. no log flags provided, bndstat should not write log
// files to log_dir. To accomplish this, all Info-like glog calls are done
// through the V() function. This should avoid a proliferation of log files
// being written through the normal use of the program, which also avoids
// having to write code to clean up those log files.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/robkingsbury/bndstat/throughput"
)

func main() {
	flag.Parse()

	devices := []string{}

	if err := bndstat(devices); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		glog.Flush()
		os.Exit(1)
	}

	glog.Flush()
	os.Exit(0)
}

func bndstat(devices []string) error {
	r, err := throughput.New()
	if err != nil {
		return err
	}

	r.Report()
	return nil
}
