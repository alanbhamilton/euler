package main

import (
  "fmt"
  "strconv"
  "strings"
  "sort"
  "math"
)

func isPandigital(n int) bool {
  s := strconv.Itoa(n)
  sl := strings.Split(s, "")
  sort.Strings(sl)
  for i := 0; i < len(sl); i++ {
    if sl[i] != strconv.Itoa(i+1) {
      return false
    }
  }
  // fmt.Println(sl)
  return true
}

func isPrime(n int) bool {
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

func main() {
  for i := 999999999; i > 100000000; i -= 2 {
    fmt.Println(i)
    if isPrime(i) && isPandigital(i) {
      fmt.Println("result:", i)
      break
    }
  }
}
