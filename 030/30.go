package main

import (
  "fmt"
  "time"
  "math"
  "strconv"
  "strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func sumOfPowers(n int, exp int) int {
  s := strconv.Itoa(n)
  sl := strings.Split(s, "")
  var sum int
  for _, num := range sl {
    base, err := strconv.Atoi(num)
    check(err)
    sum += int(math.Pow(float64(base), float64(exp)))
  }
  return sum
}

func main() {
  const Exponent = 5
  startTime := time.Now()
  var sum int
  for x := 2; x <= 1000000; x++ {
    if sumOfPowers(x, Exponent) == x {
      sum += x
      fmt.Println(x)
    }
  }
  fmt.Println("Sum:", sum)
  fmt.Println("Time Taken:", time.Since(startTime))
}
