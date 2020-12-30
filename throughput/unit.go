package throughput

import (
	"fmt"
	"math"
	"strings"
)

// A unit is used to format the value returned by Stat.Avg().
type Unit int

const (
	Bps Unit = iota
	Kbps
	Mbps
	Gbps
	Tbps
)

var allUnits = []Unit{Bps, Kbps, Mbps, Gbps, Tbps}

// String implements Stringer for a Unit.
func (u Unit) String() string {
	switch u {
	case Bps:
		return "bps"
	case Kbps:
		return "kbps"
	case Mbps:
		return "mbps"
	case Gbps:
		return "gbps"
	case Tbps:
		return "tbps"
	default:
		return "unk"
	}
}

// ParseUnit returns the Unit associated with a unit name. Use the string form
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

// Base2 returns the base 2 value of the unit.
func (u Unit) Base2() float64 {
	return math.Pow(2, float64(u*10))
}

// Base10 returns the base 10 value of the unit.
func (u Unit) Base10() float64 {
	return math.Pow(10, float64(u*3))
}
