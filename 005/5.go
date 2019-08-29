// 2520 is the smallest number that can be divided by each of the
// numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible
// by all of the numbers from 1 to 20?

// TODO: refactor based on approach from
// http://mathforum.org/library/drmath/view/62527.html

package main

import (
	"fmt"
	"math"
	"time"

	"github.com/alanbhamilton/euler/util"
)

const seriesStart int = 20
const seriesEnd int = 1

func isDivisible(num, divisor, end int) bool {
	if divisor == 1 {
		return true
	} else if divisor == end {
		return (math.Mod(float64(num), float64(divisor)) == 0)
	} else {
		return (math.Mod(float64(num), float64(divisor)) == 0) && isDivisible(num, divisor-1, end)
	}
}

func findLeastCommonMultiple(seriesStart, seriesEnd int) int {
	for i := seriesStart; ; i += seriesStart {
		if isDivisible(i, seriesStart, seriesEnd) {
			return i
		}
	}
}

func main() {
	defer util.Timetrack(time.Now(), "findLeastCommonMultiple")
	fmt.Println(findLeastCommonMultiple(seriesStart, seriesEnd))
}
