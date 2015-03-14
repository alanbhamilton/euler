// TODO: refactor this to use a precompiled list of primes for the isPrime
//       check to see how much faster it runs

package main

import (
  "fmt"
  "strconv"
  "math"
)

const Limit = 1000000

func check(e error) {
  if e != nil {
    panic(e)
  }
}

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

func isCircularPrime(n int) bool {
  s := strconv.Itoa(n)
  for i := 0; i < len(s); i++ {
    // fmt.Println(s)
    num, err := strconv.Atoi(s)
    check(err)
    if isPrime(num) == false {
      // fmt.Println("false")
      return false
    }
    first := s[len(s)-1:]
    s = first + s[:len(s)-1]
  }
  fmt.Println(s, "true")
  return true
}

func main() {
  var count int
  for i:= 2; i < Limit; i++ {
    if isCircularPrime(i) {
      count += 1
    }
  }
  fmt.Println("Count:", count)
}
