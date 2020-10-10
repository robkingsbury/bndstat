package throughput

import (
	"testing"
)

// Ensure the Linux struct satisfies the Reporter interface. Since not
// implementing the interface is a compile time error, there's no value
// to be checked here.
func TestLinuxImplementsReporter(t *testing.T) {
  var l Reporter
  l = &Linux{}
  _ = l.Report([]string{})
}
