// A Pythagorean triplet is a set of three natural numbers, a < b < c,
// for which, a^2 + b^2 = c^2

// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2

// There exists exactly one Pythagorean triplet for which a + b + c = 1000
// Find the product abc.

package main

import (
	"fmt"
	"time"

	"github.com/alanbhamilton/euler/util"
)

const sum int = 1000

func isPyTriplet(a, b, c int) bool {
	if a*a+b*b == c*c {
		return true
	}
	return false
}

func findPyTriplets(num int) [][]int {
	var triplets [][]int
	for c := num; c >= 3; c-- {
		for b := c - 1; b >= 2; b-- {
			for a := b - 1; a >= 1; a-- {
				if isPyTriplet(a, b, c) {
					triplets = append(triplets, []int{a, b, c})
				}
			}
		}
	}
	return triplets
}

func findPyTripletWithSum(sum int) []int {
	triplets := findPyTriplets(sum)
	for _, t := range triplets {
		if t[0]+t[1]+t[2] == sum {
			return t
		}
	}
	return []int{0, 0, 0}
}

func main() {
	defer util.Timetrack(time.Now(), "findPyTripletWithSum")
	triplet := findPyTripletWithSum(sum)
	product := triplet[0] * triplet[1] * triplet[2]
	fmt.Println("triplet:", triplet)
	fmt.Println("product:", product)
}
