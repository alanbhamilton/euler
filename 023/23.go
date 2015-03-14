package main

import (
  "fmt"
  "math"
  "sort"
)

const LIMIT = 28123

func getDivisors(n int) []int {
  var divisors = []int{1}
  var limit = int(math.Floor(math.Sqrt(float64(n))))
  for i := 2; i <= limit; i++ {
    if math.Mod(float64(n), float64(i)) == 0 {
      if i == n / i {
        divisors = append(divisors, i)
      } else {
        divisors = append(divisors, i, n / i)
      }
    }
  }
  sort.Ints(divisors)
  return divisors
}

func isAbundant(n int) bool {
  divisors := getDivisors(n)
  sum := 0
  for _, divisor := range divisors {
    sum += divisor
  }
  if sum > n {
    return true
  }
  return false
}

func containsInt(ints []int, n int) bool {
  for i := 0; i < len(ints); i++ {
    if n == ints[i] {
      return true
    } else if n < ints[i] {
      return false
    }
  }
  return false
}

func containsSum(n int, nums []int) bool {
  for i := 0; i < len(nums); i++ {
    diff := n - nums[i]
    if diff < 0 {
      return false
    } else if containsInt(nums, diff) {
      return true
    }
  }
  fmt.Println("containsSum fallthrough")
  return false
}

func main() {
  var abundantNums []int
  for i := 1; i <= LIMIT; i++ {
    if isAbundant(i) {
      abundantNums = append(abundantNums, i)
    }
  }
  var result int
  for i := 1; i <= LIMIT; i++ {
    if !containsSum(i, abundantNums) {
      result += i
    }
  }
  fmt.Println("result:", result)
}
