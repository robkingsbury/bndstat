package throughput

import (
	"fmt"
	"math"
	"os"
	//"text/tabwriter"
	"time"

	"github.com/golang/glog"
	"golang.org/x/crypto/ssh/terminal"
)

// A unit is used to format the value returned by Stat.Avg().
type Unit int

const (
	Bps  Unit = 1
	Kbps Unit = 10
	Mbps Unit = 20
	Gbps Unit = 30
	Tbps Unit = 40
)

// String implements Stringer for a Unit.
func (u Unit) String() string {
	switch u {
	case Kbps:
		return "kbps"
	case Mbps:
		return "mbps"
	case Gbps:
		return "gbps"
	case Tbps:
		return "tbps"
	default:
		return "bps"
	}
}

// stat reports data on how much traffic has passed through a network device.
type stat struct {
	// bytesIn is the count of inbound Bytes that passed through the interface
	// since the last innvocation of Report().
	bytesIn uint64

	// bytesOut is the count of outbound Bytes that passed through the interface
	// since the last innvocation of Report().
	bytesOut uint64

	// elapsed is the amount of time that has elapsed since the last invocation of
	// Report().
	elapsed time.Duration
}

// Stats contains a Stat for each network device.
type Stats struct {
	devices        map[string]*stat // All the stats.
	tableLineCount int              // Continuous count of lines output by Table().
}

// Devices returns a slice of device names that Stats has information on.
func (s *Stats) Devices() []string {
	devices := []string{}
	for k := range s.devices {
		devices = append(devices, k)
	}
	return devices
}

// Avg returns the average throughput for the device, in the units specified.
// If the device does not exist, zeros are returned.
func (s *Stats) Avg(device string, unit Unit) (in float64, out float64) {
	stat, ok := s.devices[device]
	if !ok {
		return 0, 0
	}

	div := math.Pow(2, float64(unit))
	in = (float64(stat.bytesIn*8) / div) / stat.elapsed.Seconds()
	out = (float64(stat.bytesOut*8) / div) / stat.elapsed.Seconds()
	return in, out
}

// Table outputs the average throughput of each device in an aligned,
// absolutely gorgeous rendition of data, designed to instill feelings of joy
// at the beauty in the world. Each time Table is called, the average
// throughput since the last call to Table is printed to stdout.
//
// devices specifies a list of devices to output. unit specifies the unit to
// use.
func (s *Stats) Table(devices []string, unit Unit) error {
	currentTerminal := int(os.Stdout.Fd())
	if !terminal.IsTerminal(currentTerminal) {
		return fmt.Errorf("stdout does not appear to be a terminal")
	}

	columns, rows, err := terminal.GetSize(currentTerminal)
	if err != nil {
		return fmt.Errorf("could not get terminal size: %s", err)
	}
	glog.V(2).Infof("columns = %d, rows = %d, tableLineCount = %d",
		columns, rows, s.tableLineCount)

	// Decide whether to print a header or not. Subtract 3 from row count so that
	// labels aren't completed rolled off the screen.
	if s.tableLineCount%(rows-3) == 0 {
		for _, device := range devices {
			fmt.Printf("%18s", device)
			fmt.Printf("%5s", " ")
		}
		fmt.Printf("\n")

		for range devices {
			fmt.Printf("%11s %11s", "In", "Out")
		}
		fmt.Printf("\n")
		s.tableLineCount += 2
	}

	for _, device := range devices {
		in, out := s.Avg(device, unit)
		fmt.Printf("%11s %11s",
			fmt.Sprintf("%.2f", in),
			fmt.Sprintf("%.2f", out))
	}
	fmt.Printf("\n")
	s.tableLineCount++

	return nil
}
