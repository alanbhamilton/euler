package main

import (
  "fmt"
  "math"
  "math/big"
  "strings"
  // "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func isPalindrome(s string) bool {
  sl := strings.Split(s, "")
  for i := 0; i < int(math.Ceil(float64(len(sl)/2))); i++ {
    if sl[i] != sl[len(sl) - 1 - i] {
      return false
    }
  }
  return true
}

func reverse(s string) string {
  var reversed string
  for i := len(s); i > 0; i-- {
    reversed += string(s[i - 1])
  }
  return reversed
}

func isLychrel(num int) bool {
  next := big.NewInt(int64(num))
  // for i := 1; i <= 50; i++ {
  for i := 1; i <= 1000; i++ {
    // fmt.Printf("%d + ", next)
    reversed := new(big.Int)
    _, ok := reversed.SetString(reverse(next.String()), 10)
    if !ok {
      panic("can not set number")
    }
    // fmt.Printf("%d = ", reversed)
    next = next.Add(next, reversed)
    // fmt.Printf("%d\n", next)
    if isPalindrome(next.String()) {
      // fmt.Println(num, "becomes palindrome in", i, "iterations")
      return false
    }
  }
  return true
}

func main() {
  var count int
  for i := 1; i < 10000; i++ {
    if isLychrel(i) {
      count++
    }
  }
  fmt.Println("result:", count)
}
