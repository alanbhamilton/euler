package main

import "testing"

func TestSumOfMultiple(t *testing.T) {
	const multiplier int = 3
	const limit int = 10
	const expected int = 18 // 3 + 6 + 9

	if got := sumOfMultiple(multiplier, limit); got != expected {
		t.Errorf("results diff on sumOfMultiple\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}

func TestSumOfMultiples(t *testing.T) {
	multipliers := []int{3, 5}
	const limit int = 10
	const expected int = 23

	if got := sumOfMultiples(multipliers, limit); got != expected {
		t.Errorf("results diff on sumOfMultiples\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}
