// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?

// check numbers up to floor(sqrt(N)) + 1, not N/2.

package main

import (
	"fmt"
	"math"
	"sort"
)

func isPrime(n int) bool {
	if n == 1 || n == 2 {
		return true
	}
	if math.Mod(float64(n), 2) == 0 {
		return false
	}
	for i := 3; i <= int(math.Floor(math.Sqrt(float64(n)))); i += 2 {
		if math.Mod(float64(n), (float64(i))) == 0 {
			return false
		}
	}
	return true
}

func isFactorOf(num1 int, num2 int) bool {
	if math.Mod(float64(num1), float64(num2)) == 0 {
		return true
	} else {
		return false
	}
}

func getPrimeFactors(num int) []int {
	retval := make([]int, 0)

	if isPrime(num) {
		retval = append(retval, num)
	} else {
		limit := int(math.Floor(math.Sqrt(float64(num))))
		for i := 2; i <= limit; i++ {
			if isPrime(i) && isFactorOf(num, i) {
				retval = append(retval, i)
				retval = append(retval, getPrimeFactors(num / i)...)
				break
			}
		}
	}
	return retval
}

func main() {
	primeFactors := getPrimeFactors(600851475143)
	sort.Ints(primeFactors)
	fmt.Println("prime factors:", primeFactors)
	largestPrimeFactor := primeFactors[len(primeFactors) - 1]
	fmt.Println("largest prime factor:", largestPrimeFactor)
}
