package main

import (
  "fmt"
  "strconv"
  "strings"
  "math/big"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func factorial(n *big.Int) *big.Int {
  x := big.NewInt(1)
  if n.Cmp(big.NewInt(0)) == 0 {
    return x
  }
  return x.Mul(n, factorial(x.Sub(n, x)))
}

func factorialDigitSum(n int) int {
  f := factorial(big.NewInt(int64(n)))
  s := f.String()
  var result int
  for _, val := range strings.Split(s, "") {
    i, err := strconv.Atoi(val)
    check(err)
    result += i
  }
  fmt.Println(s)
  return result
}

func main() {
  fmt.Println(factorialDigitSum(100))
}
