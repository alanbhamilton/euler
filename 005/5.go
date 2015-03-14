// 2520 is the smallest number that can be divided by each of the
// numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible
// by all of the numbers from 1 to 20?

package main

import "fmt"
import "math"

func isDivisible(num int, d int) bool {
  if d == 1 {
    return true
  } else {
    return (math.Mod(float64(num), float64(d)) == 0) && isDivisible(num, d - 1)
  }
}

func main() {
  for i := 20; ; i += 2 {
    if isDivisible(i, 20) {
      fmt.Println(i)
      break
    }
  }
}
