package main

import (
  "fmt"
  "math/big"
  "strings"
  "strconv"
  // "reflect"
)

const X = 2
const Y = 1000
// const Y = 15

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  result := big.NewInt(1)
  x := big.NewInt(X)

  for i := 0; i < Y; i++ {
    temp := new(big.Int)
    temp.Mul(result, x)
    result = temp
  }

  sl := strings.Split(result.String(), "")
  fmt.Println(sl)

  var sum int
  for _, n := range sl {
    temp, err := strconv.Atoi(n)
    check(err)
    sum += temp
  }
  fmt.Println(sum)
}
