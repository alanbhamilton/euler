package main

import (
  "fmt"
  "sort"
)

func permuteGen(level int, permuted string, used []bool, original string, length int, rvPtr *[]string) {
  if level == length { // permutation complete, add to return values array
    *rvPtr = append(*rvPtr, permuted)
  } else {
    for i := 0; i < len(original); i++ { // try to add an unused character
      if !used[i] {
        used[i] = true
        permuteGen(level+1, string(original[i]) + permuted, used, original, length, rvPtr) // find the permutations starting with this string
        used[i] = false
      }
    }
  }
}

func permute(s string, length int) []string {
  var used = make([]bool, len(s))
  var results []string
  permuteGen(0, "", used, s, length, &results)
  return results
}

func main() {
  permuted := permute("abcd", 2)
  sort.Strings(permuted)
  fmt.Println(permuted)
}
