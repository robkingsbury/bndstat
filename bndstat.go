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
	"sort"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/robkingsbury/bndstat/throughput"
)

var countFlag = flag.Int("count", 0, "count of updates, any zero or negative values are considered infinity")
var intervalFlag = flag.Int("interval", 3, "period time between updates in `seconds`")

func init() {
	// Define a custom usage message that more pleasing to mine eye.
	flag.Usage = func() {
		u := "Usage: bndstat [option]... [interval [count]]\n"
		u += "Output the average throughput of network devices over an interval\n"
		u += "\n"
		u += "Options:\n"
		u += "  --interval=seconds    Number of seconds between updates\n"
		u += "  --count=num           Number of updates to print, any num less than one\n"
		u += "                          will output infinite updates until \n"
		fmt.Fprintf(flag.CommandLine.Output(), u)
	}
}

func main() {
	flag.Parse()

	// Get the interval and count, which can be specified either as standard
	// flags are unflagged args.
	interval, count, err := parseUnflaggedArgs(flag.Arg(0), flag.Arg(1), *intervalFlag, *countFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing options: %s\n", err)
		glog.Flush()
		os.Exit(1)
	}
	glog.V(1).Infof("interval = %d, count = %d", interval, count)

	devices := []string{}

	if err := bndstat(devices, interval, count); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		glog.Flush()
		os.Exit(1)
	}

	glog.Flush()
	os.Exit(0)
}

func bndstat(devices []string, interval int, count int) error {
	t := throughput.NewTable()
	r, err := throughput.New()
	if err != nil {
		return err
	}

	unit := throughput.Kbps

	time.Sleep(time.Second)
	i := 1
	for {
		stats, err := r.Report()
		if err != nil {
			glog.Exitf("%s", err)
		}

		devices := stats.Devices()
		sort.Strings(devices)
		t.Write(stats, devices, unit)

		if count > 0 {
			if i >= count {
				return nil
			}
			i++
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}

	return nil
}

func parseUnflaggedArgs(interval string, count string, intervalFlag int, countFlag int) (int, int, error) {
	// If the unflagged interval is empty, return the flagged values.
	if interval == "" {
		return intervalFlag, countFlag, nil
	}

	i, err := strconv.Atoi(interval)
	if err != nil {
		return 0, 0, err
	}

	// If the unflagged count is empty, return the unflagged interval but the
	// flagged count.
	if count == "" {
		return i, countFlag, nil
	}

	c, err := strconv.Atoi(count)
	if err != nil {
		return 0, 0, err
	}

	return i, c, nil
}
