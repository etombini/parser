package parser

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {

	var flagtests = []struct {
		name        string
		expression  string
		duration    time.Duration
		expectedErr bool
	}{
		{
			name:        "normal: 1w2d3h4m5s6ms",
			expression:  "1w2d3h4m5s6ms",
			duration:    7*24*time.Hour + 2*24*time.Hour + 3*time.Hour + 4*time.Minute + 5*time.Second + 6*time.Millisecond,
			expectedErr: false,
		},
		{
			name:        "negative: -1w2d3h4m5s6ms",
			expression:  "-1w2d3h4m5s6ms",
			duration:    -(7*24*time.Hour + 2*24*time.Hour + 3*time.Hour + 4*time.Minute + 5*time.Second + 6*time.Millisecond),
			expectedErr: false,
		},
		{
			name:        "positive: +1w2d3h4m5s6ms",
			expression:  "+1w2d3h4m5s6ms",
			duration:    7*24*time.Hour + 2*24*time.Hour + 3*time.Hour + 4*time.Minute + 5*time.Second + 6*time.Millisecond,
			expectedErr: false,
		},
		{
			name:        "random: 2d1w4m3h6ms5s",
			expression:  "2d1w4m3h6ms5s",
			duration:    7*24*time.Hour + 2*24*time.Hour + 3*time.Hour + 4*time.Minute + 5*time.Second + 6*time.Millisecond,
			expectedErr: false,
		},
		{
			name:        "overlapping duration: 1d25h",
			expression:  "1d25h",
			duration:    24*time.Hour + 25*time.Hour,
			expectedErr: false,
		},
		{
			name:        "0 minute: 1w2d3h0m5s6ms",
			expression:  "1w2d3h0m5s6ms",
			duration:    7*24*time.Hour + 2*24*time.Hour + 3*time.Hour + 0*time.Minute + 5*time.Second + 6*time.Millisecond,
			expectedErr: false,
		},
		{
			name:        "missing unit: 1w2d3h6ms",
			expression:  "1w2d3h0m5s6ms",
			duration:    7*24*time.Hour + 2*24*time.Hour + 3*time.Hour + 5*time.Second + 6*time.Millisecond,
			expectedErr: false,
		},
		{
			name:        "empty",
			expression:  "",
			duration:    0,
			expectedErr: true,
		},
		{
			name:        "unknown unit (1y1d)",
			expression:  "1y1d",
			duration:    0,
			expectedErr: true,
		},
		{
			name:        "random string",
			expression:  "yadiyadiwow",
			duration:    0,
			expectedErr: true,
		},
	}

	for _, ft := range flagtests {
		d, err := ParseDuration(ft.expression)
		if ft.expectedErr && err == nil {
			t.Errorf("%s: expecting an error\n", ft.name)
		}
		if !ft.expectedErr && err != nil {
			t.Errorf("%s: not expecting an error (%s)\n", ft.name, err)
		}
		if d != ft.duration {
			t.Errorf("%s: returned %s, expecting %s\n", ft.name, d, ft.duration)
		}
	}

}
