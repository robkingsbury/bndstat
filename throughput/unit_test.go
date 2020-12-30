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

func TestBase2and10(t *testing.T) {
	tests := []struct {
		unit   Unit
		want2  float64
		want10 float64
	}{
		{
			unit:   Bps,
			want2:  1,
			want10: 1,
		},
		{
			unit:   Kbps,
			want2:  1024,
			want10: 1000,
		},
		{
			unit:   Mbps,
			want2:  1048576,
			want10: 1000000,
		},
	}

	for _, test := range tests {
		t.Run(test.unit.String(), func(t *testing.T) {
			got := test.unit.Base2()
			if got != test.want2 {
				t.Errorf("%s.base2(): got %.0f, want %.0f", test.unit, got, test.want2)
			}

			got = test.unit.Base10()
			if got != test.want10 {
				t.Errorf("%s.base10(): got %.0f, want %.0f", test.unit, got, test.want10)
			}
		})
	}
}
