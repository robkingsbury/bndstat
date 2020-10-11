package throughput

import (
	"testing"
)

// Ensure the NilReporter struct satisfies the Reporter interface.
func TestNilReporterImplementsReporter(t *testing.T) {
	var n Reporter
	n = &NilReporter{}
	t.Logf("%v", n)
}

// Report should always return an error.
func TestNilReport(t *testing.T) {
	n := &NilReporter{}
	if _, err := n.Report(); err == nil {
		t.Errorf("Report() on a NilReporter should always return an error")
	}
}
