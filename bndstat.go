// Bndstat displays simple network device throughput data, inspired by unix
// tools like vmstat, iostat, mpstat, netstat, etc.
//
// Quick start:
//
//    $ git clone https://github.com/robkingsbury/bndstat
//    $ cd bndstat
//    $ go build
//    $ ./bndstat
//    $ ./bndstat --help
//
// See https://github.com/robkingsbury/bndstat for full documentation on
// building and using the tool.
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

// tag is used when printing version info
var tag = "v0.4.1"

var countFlag = flag.Int("count", 0, "count of updates, any zero or negative values are considered infinity")
var devicesFlag = flag.String("devices", "", "comma separated list of devices to output; all non-loopback devices included if empty")
var helpfullFlag = flag.Bool("helpfull", false, "complete list of available cmdline options")
var intervalFlag = flag.Float64("interval", 1.0, "period time between updates in `seconds`")
var showUnitFlag = flag.Bool("showunit", false, "show the units used in the output header")
var versionFlag = flag.Bool("version", false, "print version information")
var unitFlag = flag.String("unit", "kbps", "the bits per second unit to use")

func init() {
	flag.Usage = func() {
		u := "Usage: bndstat [option]... [interval [count]]\n"
		u += "Output the average throughput of network devices over a time interval.\n"
		u += "\n"
		u += "Interval and count have the same behavior as the options of the same\n"
		u += "name. However, when both an option and the non-option arg are present,\n"
		u += "the value specified by the option takes precedence.\n"
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
		u += "  --showunit            Show the --unit used in the output header [false]\n"
		u += "  --version             Print version information [false]\n"
		u += "  --helpfull            List all available options [false]\n"
		fmt.Fprintf(flag.CommandLine.Output(), u)
	}
}

func main() {
	flag.Parse()

	if *helpfullFlag {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *versionFlag {
		printVersionInfo()
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
	interval, count, err := parseUnflaggedArgs()
	if err != nil {
		return fmt.Errorf("error parsing options: %s", err)
	}
	glog.V(1).Infof("interval = %f, count = %d", interval, count)

	unit, err := throughput.ParseUnit(*unitFlag)
	if err != nil {
		return err
	}

	// Calculate the intervalDuration by taking the input interval, converting it
	// to milliseconds and parsing the result. Trying to cast a float as a
	// time.Duration doesn't work well because anything less than 1 is rounded
	// down to zero, etc.
	intervalDuration, err := time.ParseDuration(fmt.Sprintf("%dms", int(interval*1000)))
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
	t.Header(devices(d), unit, *showUnitFlag)

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

		t.Write(throughput.TableWriteOpts{
			Stats:    stats,
			Devices:  d,
			Unit:     unit,
			ShowUnit: *showUnitFlag,
		})
		updateCount++

		if count > 0 && updateCount >= count {
			return nil
		}

		time.Sleep(intervalDuration)
	}

	return nil
}

// Returns the interval and count as specified on the cmdline. Since both of
// these can be set as flag options or as unflagged args, this function returns
// the preferred value if there is a conflict. When both the flagged and the
// unflagged args are present, the flagged options take precedent.
func parseUnflaggedArgs() (interval float64, count int, err error) {
	intervalFlagSet := false
	countFlagSet := false
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "interval":
			intervalFlagSet = true
		case "count":
			countFlagSet = true
		}
	})

	intervalArg := flag.Arg(0)
	intervalArgSet := false
	if intervalArg != "" {
		intervalArgSet = true
	}

	countArg := flag.Arg(1)
	countArgSet := false
	if countArg != "" {
		countArgSet = true
	}

	interval = *intervalFlag
	if !intervalFlagSet && intervalArgSet {
		i, err := strconv.ParseFloat(intervalArg, 64)
		if err != nil {
			return 0, 0, err
		}
		interval = i
	}

	count = *countFlag
	if !countFlagSet && countArgSet {
		c, err := strconv.Atoi(countArg)
		if err != nil {
			return 0, 0, err
		}
		count = c
	}

	return interval, count, nil
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

func printVersionInfo() {
	v := fmt.Sprintf("bndstat %s\n", tag)
	v += "Rob Kingsbury\n"
	v += "https://github.com/robkingsbury/bndstat\n"
	fmt.Printf("%s", v)
}
