package main

import "testing"

func TestGetNthPrime(t *testing.T) {
	const nth int = 6
	const expected int = 13

	if got := getNthPrime(nth); got != expected {
		t.Errorf("results diff on getNthPrime\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}
