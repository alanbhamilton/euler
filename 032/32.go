package main

import (
  "fmt"
  "strconv"
  "strings"
  "sort"
)

const Limit = 2000
// const Limit = 5

func isPandigital(mul1 int, mul2 int, prod int) bool {
  s1 := strconv.Itoa(mul1)
  s2 := strconv.Itoa(mul2)
  p := strconv.Itoa(prod)
  s := s1 + s2 + p
  sl := strings.Split(s, "")
  if len(sl) != 9 {
    return false
  }
  sort.Strings(sl)
  for i := 0; i < 9; i++ {
    if sl[i] != strconv.Itoa(i+1) {
      return false
    }
  }
  fmt.Println(sl)
  return true
}

func main() {
  results := make(map[int]bool)
  for i := 1; i < Limit; i++ {
    for j := 1; j < Limit; j++ {
      if isPandigital(i, j, i*j) == true {
        fmt.Println(i, j, i*j)
        results[i*j] = true
      }
    }
  }
  var sum int
  for key, _ := range results {
    sum += key
  }
  fmt.Println("Answer:", sum)
}
