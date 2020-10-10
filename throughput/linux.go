package throughput

import (
	"time"
)

// Linux implements the Reporter interface for linux systems.
type Linux struct {
	devices []deviceData
}

// deviceData is the persistent data held in a Linux struct. When Report() is
// called, it updates deviceData for all devices and returns []Stats from that
// data.
type deviceData struct {
	name string

	lastTime    time.Time
	lastBitsIn  int
	lastBitsOut int

	currentTime    time.Time
	currentBitsIn  int
	currentBitsOut int
}

// singleRead is the struct representing the bitsIn and bitsOut of a device at
// one point in time.
type singleRead struct {
	name    string
	bitsIn  int
	bitsOut int
}

func (l *Linux) Report() []Stat {
	return []Stat{}
}

// read opens /proc/net/dev and updates the deviceData for all devices.
func (l *Linux) read() error {
	return nil
}

func (l *Linux) parseNetDev(b []byte) []singleRead {
	return []singleRead{}
}
