package main

import(
  "fmt"
  "strconv"
  "strings"
)

// const MaxUint = ^uint(0)
// const MinUint = 0
// const MaxInt = int(MaxUint >> 1)
// const MinInt = -MaxInt - 1

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}

func main() {
  // for i := 3; i < MaxInt; i++ {
  var result int
  for i := 3; i < 50000; i++ {
    s := strconv.Itoa(i)
    sl := strings.Split(s, "")
    var sum int
    for _, digit := range sl {
      n, err := strconv.Atoi(digit)
      check(err)
      sum += factorial(n)
    }
    if sum == i {
      fmt.Println(i)
      result += i
    }
  }
  fmt.Println("Result:", result)
}
