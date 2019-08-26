package util

import (
	"log"
	"math"
	"strings"
	"time"
)

// Timetrack logs the executions time of a function.
// Use at the top of a function "defer util.timetrack(time.Now(), "findLargestPalindrome")"
func Timetrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	chars := strings.Split(s, "") // split string into slice of characters
	for i := 0; i < int(math.Ceil(float64(len(chars)/2))); i++ {
		if chars[i] != chars[len(chars)-1-i] {
			return false
		}
	}
	return true
}
