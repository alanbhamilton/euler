// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.

// What is the 10001st prime number?

package main

import (
	"fmt"
	"time"

	"github.com/alanbhamilton/euler/util"
)

const nth int = 10001

func getNthPrime(n int) int {
	count := 0
	var result int
	for i := 1; count <= n; i++ {
		if util.IsPrime(i) {
			count += 1
			result = i
		}
	}
	return result
}

func main() {
	defer util.Timetrack(time.Now(), "getNthPrime")
	fmt.Println(getNthPrime(nth))
}
