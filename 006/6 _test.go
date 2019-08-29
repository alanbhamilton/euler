package main

import "testing"

func TestSumOfSquares(t *testing.T) {
	const limit int = 10
	const expected int = 385

	if got := sumOfSquares(limit); got != expected {
		t.Errorf("results diff on sumOfSquares\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}

func TestSquareOfSum(t *testing.T) {
	const limit int = 10
	const expected int = 3025

	if got := squareOfSum(limit); got != expected {
		t.Errorf("results diff on squareOfSum\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}

func TestDifference(t *testing.T) {
	const limit int = 10
	const expected int = 2640

	if got := difference(limit); got != expected {
		t.Errorf("results diff on difference\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}
