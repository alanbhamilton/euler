package main

import (
  "math/big"
  "fmt"
)

func main() {
	num := big.NewInt(2)
	num = num.Exp(num, big.NewInt(1000), nil)

	sum := 0
	for _, n := range num.String() {
		// sum += (int(n)-48)
		sum += (int(n)-'0')
	}

	fmt.Println(sum)
}
