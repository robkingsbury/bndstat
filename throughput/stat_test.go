package throughput

import (
	"testing"
	"time"
)

var testAvgData = &Stats{
	devices: map[string]*stat{
		"notraffic": {
			bytesIn:  0,
			bytesOut: 0,
			elapsed:  time.Second,
		},
		"onesec": {
			bytesIn:  128,
			bytesOut: 256,
			elapsed:  time.Second,
		},
		"fivesec": {
			bytesIn:  81776,
			bytesOut: 63073,
			elapsed:  5 * time.Second,
		},
	},
}

func TestAvg(t *testing.T) {
	tests := []struct {
		name        string
		device      string
		unit        Unit
		inWant      float64
		outWant     float64
		errExpected bool
	}{
		{
			name:        "notraffic (bps)",
			device:      "notraffic",
			unit:        Bps,
			inWant:      0.0,
			outWant:     0.0,
			errExpected: false,
		},
		{
			name:        "notraffic (mbps)",
			device:      "notraffic",
			unit:        Mbps,
			inWant:      0.0,
			outWant:     0.0,
			errExpected: false,
		},
		{
			name:        "one second (bps)",
			device:      "onesec",
			unit:        Bps,
			inWant:      1024.0,
			outWant:     2048.0,
			errExpected: false,
		},
		{
			name:        "one second (kbps)",
			device:      "onesec",
			unit:        Kbps,
			inWant:      1.0,
			outWant:     2.0,
			errExpected: false,
		},
		{
			name:        "five seconds (bps)",
			device:      "fivesec",
			unit:        Bps,
			inWant:      130841.6,
			outWant:     100916.8,
			errExpected: false,
		},
		{
			name:        "five seconds (kbps)",
			device:      "fivesec",
			unit:        Kbps,
			inWant:      127.77500,
			outWant:     98.5515625,
			errExpected: false,
		},
		{
			name:        "bad device name",
			device:      "not a device in the map",
			errExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inGot, outGot, err := testAvgData.Avg(test.device, test.unit)

			if test.errExpected && err == nil {
				t.Fatalf("Error expected but none was returned")
			}

			if !test.errExpected && err != nil {
				t.Fatalf("Error not expected but one was returned: %s", err)
			}

			if inGot != test.inWant {
				t.Errorf("In: got %f, want %f", inGot, test.inWant)
			}

			if outGot != test.outWant {
				t.Errorf("Out: got %f, want %f", outGot, test.outWant)
			}
		})
	}
}
