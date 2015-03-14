package main

import (
  "fmt"
  "strings"
  "io/ioutil"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func stripChars(str, chr string) string {
  return strings.Map(func(r rune) rune {
    if strings.IndexRune(chr, r) < 0 {
      return r
    }
    return -1
  }, str)
}

func wordValue(word string) int {
  sum := 0
  for i := 0; i < len(word); i++ {
    sum += int(word[i]) - 64
  }
  return sum
}

func triangle(n int) int {
  return n * (n + 1) / 2
}

func main() {
  var triangles = make(map[int]bool)
  for i := 1; i <= 20; i++ {
    triangles[triangle(i)] = true
  }

  dat, err := ioutil.ReadFile("./words.txt")
  check(err)
  s := stripChars(string(dat), "\"")
  words := strings.Split(s, ",")

  count := 0
  for _, word := range words {
    val := wordValue(word)
    if triangles[val] == true {
      count++
    }
  }

  fmt.Println("result:", count)
}
