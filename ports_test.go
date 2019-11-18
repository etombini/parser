package parser

import "testing"

func TestParsePorts(t *testing.T) {
	testCases := []struct {
		desc    string
		in      string
		out     []uint16
		wantErr bool
	}{
		{
			desc:    "one port",
			in:      "22",
			out:     []uint16{22},
			wantErr: false,
		},
		{
			desc:    "several ports",
			in:      "22,80,443",
			out:     []uint16{22, 80, 443},
			wantErr: false,
		},
		{
			desc:    "several unsorted ports",
			in:      "22,80,443,23",
			out:     []uint16{22, 23, 80, 443},
			wantErr: false,
		},
		{
			desc:    "port interval",
			in:      "22-25",
			out:     []uint16{22, 23, 24, 25},
			wantErr: false,
		},
		{
			desc:    "port interval and unsorted port",
			in:      "443,22-25,80",
			out:     []uint16{22, 23, 24, 25, 80, 443},
			wantErr: false,
		},
		{
			desc:    "multiple port intervals and unsorted port",
			in:      "443,22-25,80,30-31",
			out:     []uint16{22, 23, 24, 25, 30, 31, 80, 443},
			wantErr: false,
		},
		{
			desc:    "multiple port intervals and unsorted port (bis)",
			in:      "443,22-25,30-31,80",
			out:     []uint16{22, 23, 24, 25, 30, 31, 80, 443},
			wantErr: false,
		},
		{
			desc:    "multiple port intervals and unsorted port (ter)",
			in:      "22-25,30-31,443,80",
			out:     []uint16{22, 23, 24, 25, 30, 31, 80, 443},
			wantErr: false,
		},
		{
			desc:    "out of range port",
			in:      "80000",
			out:     []uint16{},
			wantErr: true,
		},
		{
			desc:    "out of range interval",
			in:      "1-80000",
			out:     []uint16{},
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res, err := ParsePorts(tC.in)
			if err != nil && !tC.wantErr {
				t.Errorf("%s: wanted no error, returned %s", tC.desc, err)
			}
			if err == nil && tC.wantErr {
				t.Errorf("%s: wanted an error, but none returned", tC.desc)
			}
			if len(res) != len(tC.out) {
				t.Errorf("%s: expected %d ports, returned %d", tC.desc, len(tC.out), len(res))
			}
			for i := range res {
				if res[i] != tC.out[i] {
					t.Errorf("%s: expected %d at position %d, returned %d", tC.desc, tC.out[i], i, res[i])
				}
			}

		})
	}
}
