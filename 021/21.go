package main

import (
  "fmt"
  "math"
)

func sumOfDivisors(n int) int {
  var divisors = []int{1}
  for i := 2; i < int(math.Ceil(math.Sqrt(float64(n)))); i++ {
    if math.Mod(float64(n), float64(i)) == 0 {
      divisors = append(divisors, i, n / i)
    }
  }
  // fmt.Println("n:", n, divisors)
  var result int
  for _, num := range divisors {
    result += num
  }
  return result
}

func main() {
  var result int
  for i := 2; i < 10000; i++ {
    n := sumOfDivisors(i)
    if sumOfDivisors(n) == i && n != i {
      result += n + i
      fmt.Println("match:", i, n)
      i = n
    }
  }
  fmt.Println(result)
}
