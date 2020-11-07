package throughput

import (
	"fmt"
	"testing"
)

func TestUnits(t *testing.T) {
	for _, unit := range allUnits {
		s := fmt.Sprintf("%s", unit)

		u, err := ParseUnit(s)
		if err != nil {
			t.Errorf("unit string, %q, returning an error for ParseUnit()", s)
		}

		if u != unit {
			t.Errorf("ParseUnit(%q): got %s, want %s", s, u, unit)
		}
	}
}

func TestBadUnit(t *testing.T) {
	if _, err := ParseUnit("a bad unit string"); err == nil {
		t.Errorf("ParseUnit() should return an error when passed a bad unit string")
	}
}
