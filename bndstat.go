// Bndstat displays simple network device throughput data, inspired by unix
// tools like vmstat, iostat, mpstat, netstat, etc.
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
var helpfull = flag.Bool("helpfull", false, "complete list of available cmdline options")
var intervalFlag = flag.Float64("interval", 1.0, "period time between updates in `seconds`")
var unitFlag = flag.String("unit", "kbps", "the bits per second unit to use")

func init() {
	flag.Usage = func() {
		u := "Usage: bndstat [option]... [interval [count]]\n"
		u += "Output the average throughput of network devices over an interval.\n"
		u += "\n"
		u += "Interval and count have the same behavior as the options of the same\n"
		u += "name. However, when both an option and the non-option arg are present,\n"
		u += "the value specified in the non-option arg takes precedence.\n"
		u += "\n"
		u += "Interval is specified as a float for both the option and arg.\n"
		u += "\n"
		u += "Options [default]:\n"
		u += "  --interval=seconds    Number of seconds between updates [1.0]\n"
		u += "  --count=num           Number of updates to print, any num less than one\n"
		u += "                          will output infinite updates [0]\n"
		u += "  --devices=list        Comma separated list of devices to output; if empty,\n"
		u += "                          all non-loopback devices are included [empty]\n"
		u += "  --unit=string         Specify the output unit; is one of bps, kbps, mbps\n"
		u += "                          or tbps; input is not case sensitive [kbps]\n"
		u += "  --helpfull            List all available options [false]\n"
		fmt.Fprintf(flag.CommandLine.Output(), u)
	}
}

func main() {
	flag.Parse()

	if *helpfull {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if err := bndstat(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func bndstat() error {
	defer glog.Flush()

	// Get the interval and count, which can be specified either as standard
	// flags or unflagged args.
	interval, count, err := parseUnflaggedArgs(flag.Arg(0), flag.Arg(1), *intervalFlag, *countFlag)
	if err != nil {
		return fmt.Errorf("error parsing options: %s", err)
	}
	glog.V(1).Infof("interval = %f, count = %d", interval, count)

	unit, err := throughput.ParseUnit(*unitFlag)
	if err != nil {
		return err
	}

	r, err := throughput.NewReporter()
	if err != nil {
		return err
	}

	stats, err := r.Report()
	if err != nil {
		return err
	}

	t := throughput.NewTable()
	d := devices(stats.Devices())
	sort.Strings(d)
	t.Header(devices(d))

	// Calculate the intervalDuration by taking the input interval, converting it
	// to milliseconds and parsing the result. Trying to cast a float as a
	// time.Duration doesn't work well because anything less than 1 is rounded
	// down to zero, etc.
	intervalDuration, err := time.ParseDuration(fmt.Sprintf("%dms", int(interval*1000)))
	if err != nil {
		return err
	}

	// Sleep for an initial interval to collect data for the first output.
	time.Sleep(intervalDuration)

	updateCount := 0
	for {
		stats, err := r.Report()
		if err != nil {
			return err
		}

		d := devices(stats.Devices())
		sort.Strings(d)
		t.Write(stats, d, unit)
		updateCount++

		if count > 0 && updateCount >= count {
			return nil
		}

		time.Sleep(intervalDuration)
	}

	return nil
}

func parseUnflaggedArgs(interval string, count string, intervalFlag float64, countFlag int) (float64, int, error) {
	// If the unflagged interval is empty, return the flagged values.
	if interval == "" {
		return intervalFlag, countFlag, nil
	}

	i, err := strconv.ParseFloat(interval, 64)
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

func devices(statDevices []string) []string {
	devices := strings.Split(*devicesFlag, ",")
	if *devicesFlag == "" {
		devices = []string{}
		for _, d := range statDevices {
			if d != "lo" {
				devices = append(devices, d)
			}
		}
	}
	return devices
}
