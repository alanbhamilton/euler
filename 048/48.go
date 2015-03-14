package main

import (
  "fmt"
  "math/big"
  "time"
)

func main() {
  startTime := time.Now()
  limit := 1000
  // limit := 10
  result := big.NewInt(0)

  for n := 1; n <= limit; n++ {
    num := big.NewInt(int64(n))
    num = num.Exp(num, num, nil)
    result = result.Add(result, num)
  }
  s := result.String()
  fmt.Println(s[len(s)-10:])
  fmt.Println("Time Taken : ",time.Since(startTime))
}
