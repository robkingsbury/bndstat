package throughput

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kr/pretty"
)

// Ensure the Linux struct satisfies the Reporter interface. Since not
// implementing the interface is a compile time error, there's no value
// to be checked here. t.Logf() is called to avoid a compile error since
// l isn't actually used.
func TestLinuxImplementsReporter(t *testing.T) {
	var l Reporter
	l = &Linux{}
	t.Logf("%v", l)
}

var netDevInput1 = []byte(
	`Inter-|   Receive                                                |  Transmit
 face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed
  eth0:       0       0    0    0    0     0          0         0        0       0    0    0    0     0       0          0
 wlan0: 365688729 1011122    0   39    0     0          0    868185  6999705   31566    0    0    0     0       0          0
    lo:   24685     271    0    0    0     0          0         0    24685     271    0    0    0     0       0          0
`)

func TestParseNetDev(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		errExpected bool
		want        []*singleRead
	}{
		{
			name:        "trivial",
			input:       []byte{},
			errExpected: false,
			want:        []*singleRead{},
		},
		{
			name:        "simple stats, one liner",
			input:       []byte(`eth0: 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0`),
			errExpected: false,
			want: []*singleRead{
				{
					name:     "eth0",
					bytesIn:  1,
					bytesOut: 2,
				},
			},
		},
		{
			name:        "real input, netDevInput1",
			input:       netDevInput1,
			errExpected: false,
			want: []*singleRead{
				{
					name:     "eth0",
					bytesIn:  0,
					bytesOut: 0,
				},
				{
					name:     "wlan0",
					bytesIn:  365688729,
					bytesOut: 6999705,
				},
				{
					name:     "lo",
					bytesIn:  24685,
					bytesOut: 24685,
				},
			},
		},
		{
			name:        "cannot parse recv int",
			input:       []byte(`eth0: 1kamala 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0`),
			errExpected: true,
			want:        []*singleRead{},
		},
		{
			name:        "cannot parse trans int",
			input:       []byte(`eth0: 1 0 0 0 0 0 0 0 2harris 0 0 0 0 0 0 0`),
			errExpected: true,
			want:        []*singleRead{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := &Linux{}
			got, err := l.parseNetDev(bytes.NewReader(test.input))

			if test.errExpected && err == nil {
				t.Fatalf("Error expected but none was returned")
			}

			if !test.errExpected && err != nil {
				t.Fatalf("Error not expected but one was returned: %s", err)
			}

			if !cmp.Equal(got, test.want, cmp.AllowUnexported(singleRead{})) {
				t.Errorf("\ngot %s\nwant %s\n", pretty.Sprint(got), pretty.Sprint(test.want))
			}
		})
	}
}
