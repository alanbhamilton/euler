package main

import (
  "fmt"
  "strconv"
  "strings"
  "math"
)

const Limit = 1000000

func isPalindrome(s string) bool {
  sl := strings.Split(s, "")
  for i := 0; i < int(math.Ceil(float64(len(sl)/2))); i++ {
    if sl[i] != sl[len(sl) - 1 - i] {
      return false
    }
  }
  return true
}

func main() {
  var result int
  for i := 1; i < Limit; i++ {
    base10 := strconv.FormatInt(int64(i), 10)
    base2 := strconv.FormatInt(int64(i), 2)
    if isPalindrome(base10) && isPalindrome(base2) {
      fmt.Println(base10, base2)
      result += i
    }
  }
  fmt.Println("Result:", result)
}
