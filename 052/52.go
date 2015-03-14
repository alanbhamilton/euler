package main

import (
  "fmt"
  "strconv"
  "strings"
  "time"
)

func isPermutation(n int, lastPerm int) bool {
  ref := strconv.Itoa(lastPerm)
  s := strconv.Itoa(n)
  // if len(s) == len(ref) - 1 {
  //   s = "0" + s
  // }
  for i := 0; i < len(ref); i++ {
    if !strings.Contains(s, string(ref[i])) {
      return false
    }
  }
  return true
}

func main() {
  startTime := time.Now()
  for n := 1; ; n++ {
    if isPermutation(n, n * 2) && isPermutation(n, n * 3) &&
       isPermutation(n, n * 4) && isPermutation(n, n * 5) &&
       isPermutation(n, n * 5) {
      fmt.Println(n)
      break
    }
  }
  fmt.Println("Time Taken:",time.Since(startTime))
}
