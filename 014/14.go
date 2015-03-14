package main

import (
  "fmt"
  "math"
)

func collatzSeq(n int) []int {
  var rv []int
  if n == 1 {
    return []int{1}
  }
  var next int
  if math.Mod(float64(n), 2) == 0 {
    next = n / 2    // even
  } else {
    next = (3 * n) + 1  // odd
  }
  sl := collatzSeq(next)
  rv = append(rv, n)
  rv = append(rv, sl...)
  return rv
}

func main() {
  // fmt.Println(collatzSeq(13))
  longestLength := 0
  longestValue := 0
  for i := 1; i < 1000000; i++ {
    length := len(collatzSeq(i))
    if length > longestLength {
      longestLength = length
      longestValue = i
    }
  }
  fmt.Println("longest:", longestValue)
}
