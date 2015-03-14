package main

import (
  "fmt"
)


func makeSpiralGrid(size int) [][]int {
  var dirs = [4][2]int{{0,-1}, {1,0}, {0,1}, {-1,0}} // left, down, right, up [row, col]
  // for _, d := range dirs {
  //   fmt.Println(d)
  // }
  result := make([][]int, size)
  for i := range result {
    result[i] = make([]int, size)
  }

  dir, row, col := 0, 0, size - 1
  for num := size * size; num > 0; num-- {
    result[row][col] = num
    if row + dirs[dir][0] < 0 ||
       row + dirs[dir][0] >= size ||
       col + dirs[dir][1] < 0 ||
       col + dirs[dir][1] >= size ||
       result[row + dirs[dir][0]][col + dirs[dir][1]] != 0 {
      if dir++; dir >= len(dirs) {
        dir = 0
      }
    }
    row += dirs[dir][0]
    col += dirs[dir][1]
  }
  return result
}

func main() {
  // const gridSize = 1001
  const gridSize = 11
  grid := makeSpiralGrid(gridSize)
  for _, line := range grid {
    fmt.Println(line)
  }

  var sum int
  for i, j := 0, 0; i < gridSize; i, j = i+1, j+1 {
    sum += grid[i][j]
    if j != gridSize - j - 1 {
      sum += grid[i][gridSize - j - 1]
    }
  }
  fmt.Println(sum)
}
