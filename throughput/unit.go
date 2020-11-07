package throughput

import (
	"fmt"
	"strings"
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

var allUnits = []Unit{Bps, Kbps, Mbps, Gbps, Tbps}

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

// ParseUnit returns the Unit associated with the input. Use the string form
// of the const name for the unit that you want. Parsing is not case sensitive.
// For example, an input of "Bps" or "bps", will return the unit Bps, "Mbps",
// "mbps", or "mBps" will return the Mbps unit, etc.
//
// An error is returned if the input does not match any Unit const.
func ParseUnit(u string) (Unit, error) {
	switch strings.ToLower(u) {
	case "bps":
		return Bps, nil
	case "kbps":
		return Kbps, nil
	case "mbps":
		return Mbps, nil
	case "gbps":
		return Gbps, nil
	case "tbps":
		return Tbps, nil
	}

	return Bps, fmt.Errorf("could not parse %q into a Unit", u)
}
