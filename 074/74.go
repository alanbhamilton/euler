package main

import (
  "fmt"
  "math"
)

const LIMIT = 1000000

func factorial(x int) int {
  if x == 0 {
    return 1
  }
  return x * factorial(x-1)
}

func sumFactDigits(n int) int {
  sum := 0
  for i := n; i > 0; {
    digit := int(math.Mod(float64(i), 10))
    sum += factorial(digit)
    i = i / 10
  }
  return sum
}

func lenDigitFactChain(n int) int {
  var terms = make(map[int]bool)
  terms[n] = true
  next := n
  for ; ; {
    // fmt.Printf("%d ", next)
    next = sumFactDigits(next)
    if terms[next] == true {
      // fmt.Printf("\n")
      return len(terms)
    } else {
      terms[next] = true
    }
  }
}

func main() {
  count := 0
  for i := 1; i < LIMIT; i++ {
    if lenDigitFactChain(i) == 60 {
      count++
    }
  }
  fmt.Println("result:", count)
}
