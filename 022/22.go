package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "sort"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func alphaValue(s string) int {
  var result int
  for _, char := range s {
    if char >= 65 && char <= 90 {
      result += int(char - 64)
    } else {
      fmt.Printf("unexpected character: %c\n", char)
    }
  }
  return result
}

func main() {
  dat, err := ioutil.ReadFile("./p022_names.txt")
  // dat, err := ioutil.ReadFile("./names_test.txt")
  check(err)
  text := string(dat[:len(dat)-1])
  text = strings.Replace(text, `"`, "", -1)
  names := strings.Split(text, ",")
  sort.Strings(names)
  var sum int
  for i, name := range names {
    // fmt.Println(i + 1, name, alphaValue(name))
    sum += (i + 1) * alphaValue(name)
  }
  fmt.Println(sum)
}
