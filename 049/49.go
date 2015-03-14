package main

import (
  "fmt"
  "math"
  "sort"
  "strconv"
  "strings"
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

func permuteGen(level int, permuted string, used []bool, original string, rv map[string]bool, limit int) {
  if level == limit { // permutation complete, add to return values array
    rv[permuted] = true
  } else {
    for i := 0; i < len(original); i++ { // try to add an unused character
      // if !used[i] {
        // used[i] = true
        permuteGen(level+1, string(original[i]) + permuted, used, original, rv, limit) // find the permutations starting with this string
        // used[i] = false
      // }
    }
  }
}

func permute(s string, limit int) []string {
  var used = make([]bool, len(s))
  resultsMap := make(map[string]bool)
  permuteGen(0, "", used, s, resultsMap, limit)
  var results = make([]string, len(resultsMap))
  i := 0
  for key, _ := range resultsMap {
    results[i] = key
    i++
  }
  return results
}

func main() {
  candidates := permute("0123456789", 4)
  sort.Strings(candidates)
  primes := make(map[string][]int)
  for _, c := range candidates {
    n, err := strconv.Atoi(c)
    check(err)
    if isPrime(n) {
      sl := strings.Split(c, "")
      sort.Strings(sl)
      sorted := strings.Join(sl, "")
      primes[sorted] = append(primes[sorted], n)
    }
  }
  for _, val := range primes {
    if len(val) >= 2 {
      for i := 0; i < len(val) - 2; i++ {
        for j := i + 1; j < len(val) - 1; j++ {
          for k := j + 1; k < len(val); k++ {
            if val[k] - val[j] == val[j] - val[i] {
              fmt.Println(val[i], val[j], val[k])
            }
          }
        }
      }
    }
  }
}
