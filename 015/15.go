package main

import (
  "fmt"
)

const LATTICE = 20
const GRIDSIZE = LATTICE + 1


func main() {
  var grid [GRIDSIZE][GRIDSIZE]int

  for i := 0; i < GRIDSIZE; i++ {
    for j := 0; j < GRIDSIZE; j++ {
      if i == 0 || j == 0 {
        grid[i][j] = 1
      } else {
        grid[i][j] = grid[i-1][j] + grid[i][j-1]
      }
    }
  }

  fmt.Println(grid[GRIDSIZE-1][GRIDSIZE-1])
}
