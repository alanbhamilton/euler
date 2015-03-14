package main

import (
  "fmt"
  "math"
  "sort"
  "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func isPrime(n int) bool {
  if n == 1 {
    return false
  }
  if n == 2 {
    return true
  }
  if math.Mod(float64(n), 2) == 0 {
    return false
  }
  for i := 3; i <= int(math.Floor(math.Sqrt(float64(n)))); i += 2 {
    if math.Mod(float64(n), (float64(i))) == 0 {
      return false
    }
  }
  return true
}

func permuteGen(level int, permuted string, used []bool, original string, rv *[]string) {
  if level == len(original) { // permutation complete, add to return values array
    *rv = append(*rv, permuted)
  } else {
    for i := 0; i < len(original); i++ { // try to add an unused character
      if !used[i] {
        used[i] = true
        permuteGen(level+1, string(original[i]) + permuted, used, original, rv) // find the permutations starting with this string
        used[i] = false
      }
    }
  }
}

func permute(s string) []string {
  var used = make([]bool, len(s))
  var results []string
  permuteGen(0, "", used, s, &results)
  return results
}

func main() {
  charspace := "123456789"
  var perms []string
  fmt.Println("generating permutations...")
  for i := 0; i < len(charspace); i++ {
    fmt.Println("permuting:", charspace[:i+1])
    perms = append(perms, permute(charspace[:i+1])...)
  }
  fmt.Println("sorting", len(perms), "permutations...")
  sort.Strings(perms)
  fmt.Println("checking for primes...")
  for i := len(perms) - 1; i > 0; i-- {
    n, err := strconv.Atoi(perms[i])
    check(err)
    if isPrime(n) {
      fmt.Println("largest pandigital prime:", n)
      break
    }
  }
}
