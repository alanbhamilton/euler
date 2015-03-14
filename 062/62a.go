package main

import (
  "fmt"
  "sort"
  "strconv"
  "strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func createCubes() map[int]string {
  rv := make(map[int]string)
  for i := 3; i < 20000; i++ {
    cube := i * i * i
    s := strconv.Itoa(cube)
    sl := strings.Split(s, "")
    sort.Strings(sl)
    sorted := strings.Join(sl, "")
    rv[cube] = sorted
  }
  return rv
}

func main() {
  cubes := createCubes()
  lowestCube := 0
  for cube, perm := range cubes {
    var matches []int
    for c, p := range cubes {
      if perm == p {
        matches = append(matches, c)
      }
    }
    if len(matches) == 5 {
      if lowestCube == 0 || lowestCube > cube {
        lowestCube = cube
      }
    }
  }
  fmt.Println("result:", lowestCube)
}
