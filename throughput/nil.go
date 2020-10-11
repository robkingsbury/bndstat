package throughput

import (
	"errors"
)

// NilReporter implements the Reporter interface in a most trivial fashion.
type NilReporter struct{}

// Report always returns an empty Stat slice and an error.
func (n *NilReporter) Report() ([]*Stat, error) {
	return []*Stat{}, errors.New("cannot report on a nil reporter")
}
