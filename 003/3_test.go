package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetPrimeFactors(t *testing.T) {
	const number int = 13195
	expected := []int{5, 7, 13, 29}

	got := getPrimeFactors(number)
	sort.Ints(got)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("results diff on getPrimeFactors\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}

func TestGetLargestPrimeFactor(t *testing.T) {
	const number int = 13195
	const expected int = 29

	if got := getLargestPrimeFactor(number); got != expected {
		t.Errorf("results diff on getLargestPrimeFactor\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}
