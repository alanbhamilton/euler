package main

import (
  "fmt"
  "io/ioutil"
  "bytes"
  "strings"
  "strconv"
  "sort"
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
    l := strings.Split(line, " ")   // make values (still strings)
    rv[i] = make([]int, len(l))   // allocate 2d
    for j, value := range l {
      var err error
      rv[i][j], err = strconv.Atoi(value)
      check(err)
    }
  }
  return rv
}

func main() {
  dat, err := ioutil.ReadFile("./data.txt")
  check(err)
  grid := parseGridData(dat)
  for i := 1; i < len(grid); i++ {
    for j := 0; j < len(grid[i]); j++ {
      if j == 0 {
        grid[i][j] += grid[i-1][j]
      }
      if j > 0 && j < len(grid[i]) - 1 {
        if grid[i-1][j-1] > grid[i-1][j] {
          grid[i][j] += grid[i-1][j-1]
        } else {
          grid[i][j] += grid[i-1][j]
        }
      }
      if j == len(grid[i]) - 1 {
        val := grid[i-1][j-1]
        grid[i][j] += val
      }
    }
  }
  for _, line := range grid {
    fmt.Println(line)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(grid[len(grid)-1])))
  fmt.Println(grid[len(grid)-1][0])
}
