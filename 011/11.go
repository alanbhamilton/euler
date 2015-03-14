package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "bytes"
  "strconv"
)

type checker func([][]int, int, int) int
const prodSize = 4

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

func checkGrid(grid [][]int, fn checker) int {
  largest := 0
  for y := 0; y < len(grid); y++ {
    for x := 0 ; x < len(grid[y]); x++ {
      if value := fn(grid, x, y); value > largest {
        largest = value
      }
    }
  }
  return largest
}

func checkNtoS(grid [][]int, x int, y int) int {
  if y >= len(grid) - prodSize {
    return 0
  }
  rv := 1
  for i := 0; i < prodSize; i++ {
    rv *= grid[y+i][x]
  }
  return rv
}

func checkWtoE(grid [][]int, x int, y int) int {
  if x >= len(grid[y]) - prodSize {
    return 0
  }
  rv := 1
  for i := 0; i < prodSize; i++ {
    rv *= grid[y][x+i]
  }
  return rv
}

func checkNWtoSE(grid [][]int, x int, y int) int {
  if (x >= len(grid[y]) - prodSize) || (y >= len(grid) - prodSize) {
    return 0
  }
  rv := 1
  for i := 0; i < prodSize; i++ {
    rv *= grid[y+i][x+i]
  }
  return rv
}

func checkNEtoSW(grid [][]int, x int, y int) int {
  if (x < prodSize) || (y >= len(grid) - prodSize) {
    return 0
  }
  rv := 1
  for i := 0; i < prodSize; i++ {
    rv *= grid[y+i][x-i]
  }
  return rv
}

func main() {
  dat, err := ioutil.ReadFile("./grid.txt")
  check(err)
  grid := parseGridData(dat)
  largest := 0
  fnlist := []checker{checkNtoS, checkWtoE, checkNWtoSE, checkNEtoSW}
  for _, fn := range fnlist {
    if value := checkGrid(grid, fn); value > largest {
      largest = value
    }
  }
  fmt.Println(largest)
}
