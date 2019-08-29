// The sum of the squares of the first ten natural numbers is,
//      1^2 + 2^2 + ... + 10^2 = 385

// The square of the sum of the first ten natural numbers is,
//       (1 + 2 + ... + 10)^2 = 55^2 = 3025

// Hence the difference between the sum of the squares of the
// first ten natural numbers and the square of the sum is
//        3025 − 385 = 2640.

// Find the difference between the sum of the squares of the first
// one hundred natural numbers and the square of the sum.

package main

import "fmt"

const seriesLimit int = 100

func sumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func difference(seriesLimit int) int {
	return squareOfSum(seriesLimit) - sumOfSquares(seriesLimit)
}

func main() {
	fmt.Println(difference(seriesLimit))
}
