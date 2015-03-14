// A Pythagorean triplet is a set of three natural numbers, a < b < c,
// for which, a^2 + b^2 = c^2

// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2

// There exists exactly one Pythagorean triplet for which a + b + c = 1000
// Find the product abc.

package main

import "fmt"

func pyTest(a int, b int, c int) bool {
  if a*a + b*b == c*c {
    return true
  } else {
    return false
  }
}

func findPyTriplet(num int) []int {
  for c := num - 3; c >= (num / 3); c-- {
    a := 1
    b := num - c - a
    for ; a < b; {
      if pyTest(a, b, c) {
        return []int{a, b, c}
      }
      a++
      b--
    }
  }
  return []int{0, 0, 0}
}

func main() {
  triplet := findPyTriplet(1000)
  prod := triplet[0]
  for i := 1; i < len(triplet); i++ {
      prod *= triplet[i]
  }
  fmt.Println("triplet:", triplet)
  fmt.Println("product:", prod)
}
