package main

import (
  "fmt"
  "math"
  "strconv"
)

const Limit = 1000000

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func isPrime(n int) bool {
  // if n == 1 || n == 2 {
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

func isTruncPrime(num int) bool {
  if !isPrime(num) {
    return false
  }
  s := strconv.Itoa(num)
  for i := 0; i < len(s); i++ {
    left, err := strconv.Atoi(s[:i+1])
    check(err)
    right, err := strconv.Atoi(s[len(s)-1-i:])
    check(err)
    if !isPrime(left) || !isPrime(right) {
      return false
    }
  }
  return true
}

func main() {
  var count int
  var sum int
  for i := 9; i < Limit; i += 2 {
    if isTruncPrime(i) {
      fmt.Println(i)
      count++
      sum += i
    }
  }
  fmt.Println("count:", count)
  fmt.Println("sum:", sum)
}
