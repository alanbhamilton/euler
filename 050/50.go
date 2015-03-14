package main

import (
  "fmt"
  "math"
  "sort"
)

const Limit = 1000000

var primes []int

func makePrimes(limit int) {
  for i := 2; i <= limit; i = genNextPrime(i) {
    primes = append(primes, i)
  }
}

func isPrime(n int) bool {
  if n == 1 {
    return false
  }
  if n == 2 {
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

func genNextPrime(n int) int {
  var i int
  for i = n + 1; isPrime(i) == false; i++ {}
  return i
}

func nextPrime(n int) int {
  i := sort.SearchInts(primes, n)
  if (i + 1) < len(primes) {
    return primes[i + 1]
  }
  return genNextPrime(n)
}

func numConsecPrimes(start int, n int) int {
  sum := 0
  count := 0
  for i := start; ; i = nextPrime(i) {
    count++
    sum += i
    if sum > n {
      count = 0
      break
    } else if sum == n {
      break
    }
  }
  return count
}

func findConsecutivePrimes(n int) int {
  longestLength := 0
  var longestPrime int
  for prime := 2; prime < n; prime = nextPrime(prime) {
    fmt.Println(prime)
    for i := 2; i < (prime / 2); i = nextPrime(i) {
      length := numConsecPrimes(i, prime)
      if length > longestLength {
        longestLength = length
        longestPrime = prime
      }
    }
  }
  fmt.Println("prime:", longestPrime)
  fmt.Println("consecutive num:", longestLength)
  return longestLength
}

func main() {
  fmt.Println("building primes table up to", Limit, "...")
  makePrimes(Limit)
  fmt.Println("done building primes")
  findConsecutivePrimes(Limit)
}
