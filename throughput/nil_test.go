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
