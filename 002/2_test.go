package main

import "testing"

func TestSumEvenFib(t *testing.T) {
	const limit int = 100
	const expected int = 44 // 2 + 8 + 34

	if got := sumEvenFib(limit); got != expected {
		t.Errorf("results diff on sumEvenFib\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}
