package main

import (
  "fmt"
  "math"
  "sort"
)


func getFactors(n int) []int {
  var rv []int
  for i := 1; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
    if math.Mod(float64(n), float64(i)) == 0 {
      if i == n / i {
        rv = append(rv, i)
      } else {
        rv = append(rv, i, n / i)
      }
    }
  }
  sort.Ints(rv)
  return rv
}

func main() {
  triangle := 0
  for i := 0; ; i++ {
    triangle += i
    // fmt.Println("triangle of", i, "=", triangle, "- num of factors:", len(getFactors(triangle)))
    if len(getFactors(triangle)) > 500 {
      fmt.Println("triangle value:", triangle, "- num of factors:", len(getFactors(triangle)))
      break
    }
  }
}
