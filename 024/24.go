package main

import (
  "fmt"
  "strings"
  "strconv"
  // "math"
)

// func check(e error) {
//   if e != nil {
//     panic(e)
//   }
// }

func isPermutation(n int, lastPerm int) bool {
  ref := strconv.Itoa(lastPerm)
  s := strconv.Itoa(n)
  if len(s) == len(ref) - 1 {
    s = "0" + s
  }
  for i := 0; i < len(ref); i++ {
    if !strings.Contains(s, string(ref[i])) {
      return false
    }
  }
  return true
}

func main() {
  firstPerm := 123456789
  lastPerm := 9876543210
  count := 0
  for n := firstPerm; n <= lastPerm; n++ {
    if isPermutation(n, lastPerm) {
      fmt.Println(n, count)
      count++
      if count == 1000000 {
        fmt.Println("result:", n)
        break
      }
      // if math.Mod(float64(n), 1000000) == 0 {
      //   fmt.Println(n, "count:", count)
      // }
    }
  }
}
