package day2

import (
	"fmt"
	"strings"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/maths"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	var levels [][]int
	for _, l := range util.ReadInput(input, "\n") {
		s := strings.Split(l, " ")
		levels = append(levels, fns.Map(s, func(i string) int { return util.ParseInt[int](i) }))
	}

	fmt.Printf("first: %d\n", first(levels))
	fmt.Printf("second: %d\n", second(levels))
}

func first(levels [][]int) int {
	var safe int

	for _, level := range levels {
		if isSafe(level) {
			safe++
		}
	}

	return safe
}

func second(levels [][]int) int {
	var safe int

	for _, level := range levels {
		ok := isSafe(level)
		if !ok {
			for idx := range level {
				cp := make([]int, len(level))
				copy(cp, level)
				newLevel := append(cp[:idx], cp[idx+1:]...)
				ok = isSafe(newLevel)
				if ok {
					break
				}
			}

		}

		if ok {
			safe++
		} else {
		}
	}

	return safe
}

func isSafe(level []int) bool {
	diff1 := level[1] - level[0]
	if maths.AbsInt(diff1) < 1 || maths.AbsInt(diff1) > 3 {
		return false
	}

	for i := 2; i < len(level); i++ {
		prev2 := level[i-2]
		prev1 := level[i-1]
		curr := level[i]

		currDiff := curr - prev1
		prevDiff := prev1 - prev2

		if prevDiff > 0 && currDiff < 0 {
			return false
		}
		if prevDiff < 0 && currDiff > 0 {
			return false
		}
		if maths.AbsInt(currDiff) < 1 || maths.AbsInt(currDiff) > 3 {
			return false
		}
	}

	return true
}
