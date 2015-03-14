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
  count := 0
  var result int
  for i := 1; count <= 10001 ; i++ {
    if isPrime(i) {
      count += 1
      result = i
    }
  }
  fmt.Println(result)
}
