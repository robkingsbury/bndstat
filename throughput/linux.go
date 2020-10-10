package throughput

import (
	"time"
)

// Linux implements the Reporter interface for linux systems.
type Linux struct {
	lastReportTime time.Time
	lastBitsIn     int
	lastBitsOut    int
}

func (l *Linux) Report(devices []string) []Stat {
	return []Stat{}
}
