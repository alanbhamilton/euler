package main

import (
	"reflect"
	"testing"
)

func TestConvertStringsToNumbers(t *testing.T) {
	in := []string{"5", "7", "13", "29"}
	expected := []int{5, 7, 13, 29}

	got := convertStringsToNumbers(in)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("results diff on convertStringsToNumbers\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}

func TestFindGreatestProductOfSeq(t *testing.T) {
	in := sequence
	productWidth := 4
	expected := 5832

	if got := findGreatestProductOfSeq(in, productWidth); got != expected {
		t.Errorf("results diff on findGreatestProductOfSeq\nGOT:\n%v\nWANT:\n%v\n", got, expected)
	}
}
