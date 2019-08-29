package main

import (
	"testing"
)

func TestFindLeastCommonMultiple(t *testing.T) {
	tests := []struct {
		name    string
		skip    bool
		inStart int
		inEnd   int
		out     int
	}{
		{
			name:    "LCM for series 10 to 1",
			inStart: 10,
			inEnd:   1,
			out:     2520,
		},
		{
			name:    "LCM for series 6 to 2",
			inStart: 6,
			inEnd:   2,
			out:     60,
		},
	}

	for _, tt := range tests {
		if tt.skip {
			continue
		}
		if got := findLeastCommonMultiple(tt.inStart, tt.inEnd); got != tt.out {
			t.Errorf("results diff on %q\nGOT:\n%v\nWANT:\n%v\n", tt.name, got, tt.out)
		}
	}
}
