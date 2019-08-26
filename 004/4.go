// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/alanbhamilton/euler/util"
)

func findLargestPalindrome(numberOfProductDigits int) int {
	defer util.Timetrack(time.Now(), "findLargestPalindrome")
	maxFactor := int(math.Pow(10, float64(numberOfProductDigits))) - 1
	minFactor := int(math.Pow(10, float64(numberOfProductDigits-1)))
	for i := maxFactor; i >= minFactor; i-- {
		for j := maxFactor; j >= minFactor; j-- {
			product := i * j
			if util.IsPalindrome(strconv.Itoa(product)) {
				return product
			}
		}
	}
	return 0
}

func main() {
	result := findLargestPalindrome(3)
	fmt.Println(result)
}
