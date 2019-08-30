package main

import (
	"reflect"
	"testing"
)

func TestIsPyTriplet(t *testing.T) {
	const a, b, c int = 3, 4, 5
	const expected bool = true

	if got := isPyTriplet(a, b, c); got != expected {
		t.Errorf("results diff on isPyTriplet\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}

func TestFindPyTriplets(t *testing.T) {
	const in int = 12
	expected := [][]int{{6, 8, 10}, {3, 4, 5}}

	if got := findPyTriplets(in); !reflect.DeepEqual(got, expected) {
		t.Errorf("results diff on findPyTriplets\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}

func TestFindPyTripletWithSum(t *testing.T) {
	const in int = 12
	expected := []int{3, 4, 5}

	if got := findPyTripletWithSum(in); !reflect.DeepEqual(got, expected) {
		t.Errorf("results diff on findPyTripletWithSum\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}
