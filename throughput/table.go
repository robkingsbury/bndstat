package throughput

import (
	"fmt"
	"os"

	"github.com/golang/glog"
	"golang.org/x/crypto/ssh/terminal"
)

// Table is used to print device Stats in a nicely formatted output.
type Table struct {
	tableLineCount  int
	prevDeviceCount int
}

// NewTable() returns a pointer to a Table.
func NewTable() *Table {
	return &Table{}
}

// TableWriteOpts is used to specify options by Table.Write().
type TableWriteOpts struct {
	// Stats should be a pointer to a Stats struct with network
	// device throughput data.
	Stats *Stats

	// Devices is a slice of devices from Stats to display.
	Devices []string

	// Unit specifies the unit to use when display the throughput.
	Unit Unit

	// ShowUnit will print the unit type in the table header when set
	// to true.
	ShowUnit bool
}

// Write outputs the average throughput of each device in an aligned,
// absolutely gorgeous rendition of data, designed to instill feelings of joy
// at the beauty in the world. Each time Table is called, the average
// throughput since the last call to Table is printed to stdout.
func (t *Table) Write(opt TableWriteOpts) error {
	currentTerminal := int(os.Stdout.Fd())
	if !terminal.IsTerminal(currentTerminal) {
		return fmt.Errorf("stdout does not appear to be a terminal")
	}

	columns, rows, err := terminal.GetSize(currentTerminal)
	if err != nil {
		return fmt.Errorf("could not get terminal size: %s", err)
	}
	glog.V(2).Infof("columns = %d, rows = %d, tableLineCount = %d",
		columns, rows, t.tableLineCount)

	deviceCountChanged := false
	if len(opt.Devices) != t.prevDeviceCount {
		glog.V(1).Infof("Device count has changed from %d to %d",
			t.prevDeviceCount, len(opt.Devices))
		deviceCountChanged = true
	}

	// Decide whether to print a header or not. Subtract 3 from row count so that
	// labels aren't completed rolled off the screen.
	glog.V(2).Infof("tableLineCount = %d, rows-3 = %d", t.tableLineCount, rows-3)
	if t.tableLineCount%(rows-3) == 0 || deviceCountChanged {
		t.Header(opt.Devices, opt.Unit, opt.ShowUnit)
	}

	for _, device := range opt.Devices {
		in, out := opt.Stats.Avg(device, opt.Unit)
		fmt.Printf("%11s %11s",
			fmt.Sprintf("%.2f", in),
			fmt.Sprintf("%.2f", out),
		)
	}
	fmt.Printf("\n")
	t.tableLineCount++

	t.prevDeviceCount = len(opt.Devices)
	return nil
}

// Header prints the table header lines only.
func (t *Table) Header(devices []string, unit Unit, showunit bool) {
	for _, device := range devices {
		fmt.Printf("%18s", device)
		fmt.Printf("%5s", " ")
	}
	fmt.Printf("\n")

	if showunit {
		for range devices {
			fmt.Printf("%18s", unit)
			fmt.Printf("%5s", " ")
		}
		fmt.Printf("\n")
	}

	for range devices {
		fmt.Printf("%11s %11s", "In", "Out")
	}
	fmt.Printf("\n")

	t.tableLineCount += 2
	t.prevDeviceCount = len(devices)
}
