package main

import (
  "fmt"
  "math"
  "strings"
  "strconv"
)

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
  result := 0;
  for i := 999; i >= 100; i-- {
    for j := 999; j >= 100; j-- {
      if isPalindrome(strconv.Itoa(i * j)) {
        if i * j > result {
          result = i * j
        }
      }
    }
  }
  fmt.Println(result)
}
