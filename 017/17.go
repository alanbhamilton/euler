package main

import(
  "fmt"
  "strconv"
)

var under20 = [...]string { "",   // zero case, does not get printed
                            "one", "two", "three", "four", "five",
                            "six", "seven", "eight", "nine", "ten",
                            "eleven", "twelve", "thirteen", "fourteen", "fifteen",
                            "sixteen", "seventeen", "eighteen", "nineteen" }

var decades = [...]string { "twenty", "thirty", "forty", "fifty",
                            "sixty", "seventy", "eighty", "ninety"}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func toNaturalLang(n int) string {
  var retval string
  var remainder int
  s := strconv.Itoa(n)

  if n <= 19 {
    return under20[n]    // end case
  } else if n <= 99 {
    var err error
    decade, err := strconv.Atoi(string(s[:1]))
    check(err)
    retval = decades[decade-2]
    remainder, err = strconv.Atoi(string(s[1:]))
    check(err)
  } else if n <= 999 {
    var err error
    hundreds, err := strconv.Atoi(string(s[:1]))
    check(err)
    if hundreds > 0 {
      retval = under20[hundreds] + "hundred"
    }
    remainder, err = strconv.Atoi(string(s[1:]))
    check(err)
    if remainder > 0 {
      retval = retval + "and"
    }
  } else if n == 1000 {
    retval = "one" + "thousand"
  }
  return retval + toNaturalLang(remainder)
}

func main() {
  var result int
  for i := 1; i <= 1000; i++ {
    result += len(toNaturalLang(i))
  }
  fmt.Println("length:", result)
}
