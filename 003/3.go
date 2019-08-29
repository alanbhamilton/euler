// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?

// check numbers up to floor(sqrt(N)), not N/2.

package main

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/alanbhamilton/euler/util"
)

const number int = 600851475143

func getPrimeFactors(num int) []int {
	retval := make([]int, 0)

	if util.IsPrime(num) {
		retval = append(retval, num)
	} else {
		limit := int(math.Floor(math.Sqrt(float64(num))))
		for i := 2; i <= limit; i++ {
			if util.IsPrime(i) && util.IsFactorOf(num, i) {
				remainder := num / i
				retval = append(retval, i)
				retval = append(retval, getPrimeFactors(remainder)...)
				break
			}
		}
	}
	return retval
}

func getLargestPrimeFactor(n int) int {
	primeFactors := getPrimeFactors(n)
	sort.Ints(primeFactors)
	// fmt.Println("prime factors:", primeFactors)
	return primeFactors[len(primeFactors)-1]
}

func main() {
	defer util.Timetrack(time.Now(), "getPrimeFactors")
	largestPrimeFactor := getLargestPrimeFactor(number)
	fmt.Println("largest prime factor:", largestPrimeFactor)
}
