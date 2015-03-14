package main

import (
  "fmt"
  "math"
)

func triangle(n int) int {
  return n * (n + 1) / 2
}

func triRoot(x int) int {
  result := math.Sqrt(float64(8 * x + 1)) - 1
  if math.Mod(result, 2) == 0 {
    return int(result / 2)
  }
  return 0
}

func pentagonal(n int) int {
  return n * (3 * n - 1) / 2
}

func pentRoot(x int) int {
  result := math.Sqrt(float64(24 * x + 1)) + 1
  if math.Mod(result, 6) == 0 {
    return int(result / 6)
  }
  return 0
}

func hexagonal(n int) int {
  return n * (2 * n - 1)
}

func hexRoot(x int) int {
  result := math.Sqrt(float64(8 * x + 1)) + 1
  if math.Mod(result, 4) == 0 {
    return int(result / 4)
  }
  return 0
}

func main() {
  for i := 286; ; i++ {
  // for i := 3; ; i++ {
    x := triangle(i)
    pn, hn := pentRoot(x), hexRoot(x)
    if pn != 0 && hn != 0 {
      fmt.Println("Tn", i)
      fmt.Println("tri", i, triangle(i))
      fmt.Println("pent", pn, pentagonal(pn))
      fmt.Println("hex", hn, hexagonal(hn))
      break
    }
  }
}
