package main

import (
  "fmt"
  "math/big"
  "time"
  "strconv"
  "strings"
  "sort"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

const Limit = 100

func main() {
  startTime := time.Now()
  var results []int
  for base := 1; base < Limit; base++ {
    for exp := 1; exp < Limit; exp++ {
      bigBase := big.NewInt(int64(base))
      bigExp := big.NewInt(int64(exp))
      bigBase = bigBase.Exp(bigBase, bigExp, nil)
      s := bigBase.String()
      sl := strings.Split(s, "")
      var sum int
      for _, n := range sl {
        num, err := strconv.Atoi(n)
        check(err)
        sum += num
      }
      results = append(results, sum)
      fmt.Println(s)
    }
  }
  sort.Ints(results)
  fmt.Println(results)
  fmt.Println("Time Taken:",time.Since(startTime))
}
