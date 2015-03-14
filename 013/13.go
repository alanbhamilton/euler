package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "bytes"
  "strconv"
)


func check(e error) {
  if e != nil {
    panic(e)
  }
}

func splitAtNewline(s []byte) []string {
  eol := []byte{'\n'}
  lines := bytes.Split(s, eol)
  var rv []string
  for _, line := range lines {
    rv = append(rv, string(line))
  }
  if len(rv[len(rv)-1]) == 0 {
    rv = rv[:len(rv)-1]   // remove empty last string if insterted by Split
  }
  return rv
}

func parseGridData(s []byte) [][]int {
  lines := splitAtNewline(s)    // make array of strings
  rv := make([][]int, len(lines))   // allocate 1d
  for i, line := range lines {
    l := strings.Split(line, "")   // make values (still strings)
    rv[i] = make([]int, len(l))   // allocate 2d
    for j, value := range l {
      var err error
      rv[i][j], err = strconv.Atoi(value)
      check(err)
    }
  }
  return rv
}

func carryUp(sums []int, i int) {
  if i == len(sums) - 1 {
    return
  } else {
    if sums[i] > 9 {
      s := strconv.Itoa(sums[i])
      sl := strings.Split(s, "")

      lsd, err := strconv.Atoi(sl[len(sl)-1:][0])
      check(err)
      sums[i] = lsd

      carry, err := strconv.Atoi(strings.Join(sl[:len(sl)-1], ""))
      check(err)

      sums[i+1] += carry
    }
    carryUp(sums, i + 1)
    return
  }
}

func carry(sums []int) []int {
  carryUp(sums, 0)
  if sums[len(sums)-1] > 9 {
    s := strconv.Itoa(sums[len(sums)-1])
    sl := strings.Split(s, "")
    sl = reverseString(sl)
    var err error
    sums[len(sums)-1], err = strconv.Atoi(sl[0])
    check(err)
    for i := 1; i < len(sl); i++ {
      n, err := strconv.Atoi(sl[i])
      check(err)
      sums = append(sums, n)
    }
  }
  return sums
}

func reverseInt(s []int) []int {
  for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
  return s
}

func reverseString(s []string) []string {
  for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
  return s
}

func main() {
  dat, err := ioutil.ReadFile("./grid.txt")
  // dat, err := ioutil.ReadFile("./test.txt")
  check(err)

  data := parseGridData(dat)

  var sums []int
  for i := len(data[0]) - 1; i >= 0; i-- {
    sum := 0
    for j := 0; j < len(data); j++ {
      sum += data[j][i]
    }
    sums = append(sums, sum)
  }

  sums = carry(sums)
  sums = reverseInt(sums)
  fmt.Println(sums[:10])

}
