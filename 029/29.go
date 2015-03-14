package main

import (
  "fmt"
  "math/big"
)

func main() {
  limit := 100
  mapVals := make(map[string]bool)

  for n := 2; n <= limit; n++ {
    for e := 2; e <= limit; e++ {
      num := big.NewInt(int64(n))
      exp := big.NewInt(int64(e))
      num = num.Exp(num, exp, nil)
      mapVals[num.String()] = true
    }
  }
  fmt.Println(len(mapVals))
}
