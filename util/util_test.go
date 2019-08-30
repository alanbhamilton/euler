package util

import (
	"reflect"
	"testing"
)

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

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name string
		skip bool
		in   int
		out  bool
	}{
		{
			name: "1 is prime",
			in:   1,
			out:  true,
		},
		{
			name: "2 is prime",
			in:   2,
			out:  true,
		},
		{
			name: "7 is prime",
			in:   7,
			out:  true,
		},
		{
			name: "8 is not prime",
			in:   8,
			out:  false,
		},
		{
			name: "21 is not prime",
			in:   21,
			out:  false,
		},
	}

	for _, tt := range tests {
		if tt.skip {
			continue
		}
		if got := IsPrime(tt.in); got != tt.out {
			t.Errorf("results diff on %q\nGOT:\n%t\nWANT:\n%t\n", tt.name, got, tt.out)
		}
	}
}

func TestIsFactorOf(t *testing.T) {
	tests := []struct {
		name     string
		skip     bool
		inNumber int
		inFactor int
		out      bool
	}{
		{
			name:     "2 is a factor of 8",
			inNumber: 8,
			inFactor: 2,
			out:      true,
		},
		{
			name:     "3 is not a factor of 8",
			inNumber: 8,
			inFactor: 3,
			out:      false,
		},
	}

	for _, tt := range tests {
		if tt.skip {
			continue
		}
		if got := IsFactorOf(tt.inNumber, tt.inFactor); got != tt.out {
			t.Errorf("results diff on %q\nGOT:\n%t\nWANT:\n%t\n", tt.name, got, tt.out)
		}
	}
}

func TestGetPrimes(t *testing.T) {
	tests := []struct {
		name string
		skip bool
		in   int
		out  []int
	}{
		{
			name: "primes below 10",
			in:   10,
			out:  []int{2, 3, 5, 7},
		},
		{
			name: "primes below 121",
			in:   121,
			out:  []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113},
		},
	}

	for _, tt := range tests {
		if tt.skip {
			continue
		}
		if got := GetPrimes(tt.in); !reflect.DeepEqual(got, tt.out) {
			t.Errorf("results diff on %q\nGOT:\n%v\nWANT:\n%v\n", tt.name, got, tt.out)
		}
	}
}
