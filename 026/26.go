package main

import (
  "fmt"
  "strconv"
  "strings"
)

const Limit = 1000

func main() {
  for i := 2; i <= Limit; i++ {
    // s := strconv.Itoa(1 / float64(i))
    s := strconv.FormatFloat(1 / float64(i), 'f', 65, 64)
    s = s[2:]
    s = strings.Trim(s, "0")
    fmt.Printf("1/%d = %s\n", i, s)
  }
}
