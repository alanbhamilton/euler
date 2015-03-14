package main

import (
  "fmt"
  "math"
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

func main() {
  sum := 2
  for i := 3; i < 2000000; i += 2 {
  // for i := 3; i < 10; i += 2 {
    if isPrime(i) {
      sum += i
    }
  }
  fmt.Println(sum)
}
