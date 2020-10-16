// Bndstat displays simple network device throughput data, inspired by unix
// tools such as vmstat, iostat, mpstat, netstat, etc.
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/robkingsbury/bndstat/throughput"
)

var countFlag = flag.Int("count", 0, "count of updates, any zero or negative values are considered infinity")
var devicesFlag = flag.String("devices", "", "comma separated list of devices to output; all non-loopback devices included if empty")
var intervalFlag = flag.Int("interval", 3, "period time between updates in `seconds`")

func init() {
	// Define a custom usage message that more pleasing to mine eye.
	flag.Usage = func() {
		u := "Usage: bndstat [option]... [interval [count]]\n"
		u += "Output the average throughput of network devices over an interval.\n"
		u += "\n"
		u += "Interval and count have the same behavior as the options of the same\n"
		u += "name. However, when both an option and the non-option arg are present,\n"
		u += "the value specified in the option takes precedence.\n"
		u += "\n"
		u += "Options:\n"
		u += "  --interval=seconds    Number of seconds between updates\n"
		u += "  --count=num           Number of updates to print, any num less than one\n"
		u += "                          will output infinite updates until\n"
		u += "  --devices=list        Comma separated list of devices to output; if empty,\n"
		u += "                          all non-loopback devices are included\n"
		fmt.Fprintf(flag.CommandLine.Output(), u)
	}
}

func main() {
	flag.Parse()

	if err := bndstat(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func bndstat() error {
	defer glog.Flush()

	// Get the interval and count, which can be specified either as standard
	// flags are unflagged args.
	interval, count, err := parseUnflaggedArgs(flag.Arg(0), flag.Arg(1), *intervalFlag, *countFlag)
	if err != nil {
		return fmt.Errorf("error parsing options: %s", err)
	}
	glog.V(1).Infof("interval = %d, count = %d", interval, count)

	t := throughput.NewTable()
	r, err := throughput.New()
	if err != nil {
		return err
	}

	unit := throughput.Kbps
	devices := strings.Split(*devicesFlag, ",")

	// Sleep for a small amount of time so that the first line output is not all
	// zeros.
	time.Sleep(10 * time.Millisecond)

	updateCount := 1
	for {
		stats, err := r.Report()
		if err != nil {
			glog.Exitf("%s", err)
		}

		if *devicesFlag == "" {
			devices = []string{}
			for _, d := range stats.Devices() {
				if d != "lo" {
					devices = append(devices, d)
				}
			}
		}

		sort.Strings(devices)
		t.Write(stats, devices, unit)

		if count > 0 && updateCount >= count {
			return nil
		}

		updateCount++
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
