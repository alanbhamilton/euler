package util

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		skip bool
		in   string
		out  bool
	}{
		{
			name: "palindrome with odd number of chars",
			in:   "abcdcba",
			out:  true,
		},
		{
			name: "palindrome with even number of chars",
			in:   "abccba",
			out:  true,
		},
		{
			name: "not a palindrome",
			in:   "abcdcb",
			out:  false,
		},
	}

	for _, tt := range tests {
		if tt.skip {
			continue
		}
		if got := IsPalindrome(tt.in); got != tt.out {
			t.Errorf("results diff on %q\nGOT:\n%t\nWANT:\n%t\n", tt.name, got, tt.out)
		}
	}
}
