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

// IsPrime checks if an integer is a prime number
func IsPrime(n int) bool {
	if n == 1 || n == 2 {
		return true // 1 and 2 are primes
	}
	if math.Mod(float64(n), 2) == 0 {
		return false // any even number is not a prime
	}
	// check up to the square root for divisiblity
	limit := int(math.Floor(math.Sqrt(float64(n))))
	for i := 3; i <= limit; i += 2 {
		if math.Mod(float64(n), (float64(i))) == 0 {
			return false // divisible, not a prime
		}
	}
	return true // it's a prime
}

// IsFactorOf checks if one integer is a factor of another
func IsFactorOf(num1 int, num2 int) bool {
	if math.Mod(float64(num1), float64(num2)) == 0 {
		return true
	}
	return false
}

type number struct {
	value   int
	isPrime bool
}

func findNextPrime(numbers []number, index int) int {
	limit := int(math.Floor(math.Sqrt(float64(len(numbers)))))
	for i := index + 1; i <= limit; i++ {
		if numbers[i].isPrime {
			return i
		}
	}
	return len(numbers)
}

// GetPrimes calculates all primes below an integer value using sieve of eratosthenes
func GetPrimes(n int) []int {
	// initialize the numbers struct of size N
	numbers := make([]number, n)
	for i := 2; i < len(numbers); i++ {
		numbers[i] = number{value: i, isPrime: true}
	}
	// look for the primes up the the square root of N and mark multiples as not prime
	limit := int(math.Floor(math.Sqrt(float64(n))))
	for p := 2; p <= limit; p = findNextPrime(numbers, p) {
		for i := p + p; i < n; i += p {
			numbers[i].isPrime = false
		}
	}
	// collect the primes
	var primes []int
	for _, num := range numbers {
		if num.isPrime {
			primes = append(primes, num.value)
		}
	}
	return primes
}
