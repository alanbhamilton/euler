package main

import (
  "fmt"
)

func main() {
  squareSpiral := make([][]int, 1)
  squareSpiral[0] = make([]int, 1)
  squareSpiral[0][0] = 1

  length := len(squareSpiral[0])
  newLines := make([]int, length)
  squareSpiral = append(newLines, squareSpiral...)
  fmt.Println(squareSpiral)
}
