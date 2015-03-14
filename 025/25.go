package main

import "fmt"
import "math/big"

func main() {
  a, b := big.NewInt(1), big.NewInt(1)
  // fmt.Println("f1 =", a)
  for i := 2; ; i++{
    // fmt.Printf("f%d = %d\n",i, b)
    if (len(b.String())) >= 1000 {
      fmt.Println(i)
      break
    }
    t := big.NewInt(0)
    t = t.Add(t, b)
    b.Add(b, a)
    a = t
    // fmt.Printf("a: %d  b: %d\n", a, b)
  }

}
