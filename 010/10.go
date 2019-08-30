// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.

// TODO: calculate all primes using a seive, then get the sum

package main

import (
	"fmt"
	"time"

	"github.com/alanbhamilton/euler/util"
)

const limit int = 2e6

func sumOfPrimes(n int) int {
	primes := util.GetPrimes(n)
	sum := 0
	for _, p := range primes {
		sum += p
	}
	return sum
}

func main() {
	defer util.Timetrack(time.Now(), fmt.Sprintf("sum of primes below %d", limit))
	fmt.Println(sumOfPrimes(limit))
}
