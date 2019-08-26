package main

import "testing"

const numberDigits int = 2
const expected int = 9009

func TestFindLargestPalindrome(t *testing.T) {
	if got := findLargestPalindrome(numberDigits); got != expected {
		t.Errorf("results diff on problem 4 test case\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}
