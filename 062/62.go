package main

import (
  "fmt"
  "math"
  "strconv"
  // "sort"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func permuteGen(level int, permuted string, used []bool, original string, rv map[string]bool) {
  if level == len(original) { // permutation complete, add to return values array
    rv[permuted] = true
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
  fmt.Printf("permute... ")
  var used = make([]bool, len(s))
  resultsMap := make(map[string]bool)
  permuteGen(0, "", used, s, resultsMap)
  var results = make([]string, len(resultsMap))
  i := 0
  for key, _ := range resultsMap {
    results[i] = key
    i++
  }
  fmt.Printf(" %d perms\n", len(resultsMap))
  return results
}

func numCubicPerms(n int) int {
  cube := int(math.Pow(float64(n), 3))
  fmt.Printf("%d : %d : ", n, cube)
  s := strconv.Itoa(cube)
  perms := permute(s)
  results := make(map[int]bool)
  for _, perm := range perms {
    num, err := strconv.Atoi(perm)
    check(err)
    root := int(math.Cbrt(float64(num)))
    if int(math.Pow(float64(root), 3)) == num {
      results[root] = true
      // fmt.Println(perm, root)
    }
  }
  if len(results) >= 3 {
    fmt.Println(n, results)
  }
  return len(results)
}

func main() {
  for i := 990; ; i++ {
    // fmt.Printf("%d ", i)
    if numCubicPerms(i) == 5 {
      fmt.Println("result:", i)
      break
    }
  }

}
