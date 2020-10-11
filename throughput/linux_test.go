package throughput

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
			l := newLinux()
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

func TestUpdate(t *testing.T) {
	tests := []struct {
		name  string
		now   time.Time
		input []*singleRead
		state map[string]*deviceData
		want  map[string]*deviceData
	}{
		{
			name:  "trivial",
			now:   time.Unix(0, 0),
			input: []*singleRead{},
			state: map[string]*deviceData{},
			want:  map[string]*deviceData{},
		},
		{
			name: "start empty",
			now:  time.Unix(1, 0),
			input: []*singleRead{
				{
					name:     "eth0",
					bytesIn:  1,
					bytesOut: 10,
				},
				{
					name:     "eth1",
					bytesIn:  5,
					bytesOut: 50,
				},
			},
			state: map[string]*deviceData{},
			want: map[string]*deviceData{
				"eth0": {
					lastTime:        time.Time{},
					lastBytesIn:     0,
					lastBytesOut:    0,
					currentTime:     time.Unix(1, 0),
					currentBytesIn:  1,
					currentBytesOut: 10,
				},
				"eth1": {
					lastTime:        time.Time{},
					lastBytesIn:     0,
					lastBytesOut:    0,
					currentTime:     time.Unix(1, 0),
					currentBytesIn:  5,
					currentBytesOut: 50,
				},
			},
		},
		{
			name: "start initialized",
			now:  time.Unix(3, 0),
			input: []*singleRead{
				{
					name:     "eth0",
					bytesIn:  30,
					bytesOut: 300,
				},
				{
					name:     "eth1",
					bytesIn:  3000,
					bytesOut: 30000,
				},
			},
			state: map[string]*deviceData{
				"eth0": {
					lastTime:        time.Unix(1, 0),
					lastBytesIn:     10,
					lastBytesOut:    100,
					currentTime:     time.Unix(2, 0),
					currentBytesIn:  20,
					currentBytesOut: 200,
				},
				"eth1": {
					lastTime:        time.Unix(1, 0),
					lastBytesIn:     1000,
					lastBytesOut:    10000,
					currentTime:     time.Unix(2, 0),
					currentBytesIn:  2000,
					currentBytesOut: 20000,
				},
			},
			want: map[string]*deviceData{
				"eth0": {
					lastTime:        time.Unix(2, 0),
					lastBytesIn:     20,
					lastBytesOut:    200,
					currentTime:     time.Unix(3, 0),
					currentBytesIn:  30,
					currentBytesOut: 300,
				},
				"eth1": {
					lastTime:        time.Unix(2, 0),
					lastBytesIn:     2000,
					lastBytesOut:    20000,
					currentTime:     time.Unix(3, 0),
					currentBytesIn:  3000,
					currentBytesOut: 30000,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := &Linux{
				devices: test.state,
			}
			l.update(test.input, test.now)

			o := []cmp.Option{
				cmp.AllowUnexported(deviceData{}),
				cmpopts.EquateEmpty(),
			}
			if !cmp.Equal(l.devices, test.want, o...) {
				t.Errorf("\ngot %s\nwant %s\n", pretty.Sprint(l.devices), pretty.Sprint(test.want))
			}
		})
	}
}
