package main

import (
	"testing"
)

func TestSumOfPrimes(t *testing.T) {
	const in int = 10
	const expected int = 17

	if got := sumOfPrimes(in); got != expected {
		t.Errorf("results diff on sumOfPrimes\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}
