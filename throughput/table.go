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

// Write outputs the average throughput of each device in an aligned,
// absolutely gorgeous rendition of data, designed to instill feelings of joy
// at the beauty in the world. Each time Table is called, the average
// throughput since the last call to Table is printed to stdout.
//
// devices specifies a list of devices to output. unit specifies the unit to
// use.
func (t *Table) Write(stats *Stats, devices []string, unit Unit) error {
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
	if len(devices) != t.prevDeviceCount {
		glog.V(1).Infof("Device count has changed from %d to %d",
			t.prevDeviceCount, len(devices))
		deviceCountChanged = true
	}

	// Decide whether to print a header or not. Subtract 3 from row count so that
	// labels aren't completed rolled off the screen.
	glog.V(2).Infof("tableLineCount = %d, rows-3 = %d", t.tableLineCount, rows-3)
	if t.tableLineCount%(rows-3) == 0 || deviceCountChanged {
		t.Header(devices)
	}

	for _, device := range devices {
		in, out := stats.Avg(device, unit)
		fmt.Printf("%11s %11s",
			fmt.Sprintf("%.2f", in),
			fmt.Sprintf("%.2f", out))
	}
	fmt.Printf("\n")
	t.tableLineCount++

	t.prevDeviceCount = len(devices)
	return nil
}

// Header prints the table header lines only.
func (t *Table) Header(devices []string) {
	for _, device := range devices {
		fmt.Printf("%18s", device)
		fmt.Printf("%5s", " ")
	}
	fmt.Printf("\n")

	for range devices {
		fmt.Printf("%11s %11s", "In", "Out")
	}
	fmt.Printf("\n")

	t.tableLineCount += 2
	t.prevDeviceCount = len(devices)
}
