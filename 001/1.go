// Problem:
// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.

// Approach:
// Compute the sum of the series for each multiplier (3 and 5) and add them.
// We must subtact the sum of the series each series has in common (in this
// case 15). In the interest of learning I'm choosing to implement the
// solution so that the parameters can be easily changed.

package main

import (
	"fmt"
	"time"

	"github.com/alanbhamilton/euler/util"
)

var multipliers = []int{3, 5}

const limit int = 1000

func sumOfMultiple(multiplier int, limit int) int {
	sum := 0
	for multiple := multiplier; multiple < limit; multiple += multiplier {
		sum += multiple
	}
	return sum
}

func sumOfMultiples(m []int, limit int) int {
	defer util.Timetrack(time.Now(), "sumOfMultiples")
	sum := 0
	for _, multiplier := range m {
		sum += sumOfMultiple(multiplier, limit) // collect the sums
	}
	// we need to subtract the sum series of the multiples in common
	// find all permutations
	for i := 0; i < len(m)-1; i++ {
		for j := i + 1; j < len(m); j++ {
			sum -= sumOfMultiple(m[i]*m[j], limit)
		}
	}
	return sum
}

func main() {
	fmt.Printf("sumOfMultiples: %v\n", sumOfMultiples(multipliers, limit))
}
