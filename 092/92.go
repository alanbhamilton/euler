package main

import (
  "fmt"
  "strconv"
)

const LIMIT = 10000000

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func squareDigits(n int) int {
  s := strconv.Itoa(n)
  var next int
  for i := 0; i < len(s); i++ {
    d, err := strconv.Atoi(string(s[i]))
    check(err)
    next += d * d
  }
  return next
}

func squareDigitChain(n int) int {
  // fmt.Printf("%d ", n)
  if n == 1 || n == 89 {
    // fmt.Printf("\n")
    return n
  }
  return squareDigitChain(squareDigits(n))
}

func main() {
  result := 0
  for i := 1; i < LIMIT; i++ {
    if squareDigitChain(i) == 89 {
      result++
    }
  }
  fmt.Println("result:", result)
}
