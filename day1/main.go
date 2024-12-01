package day1

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mbark/aoc2024/maths"
	"github.com/mbark/aoc2024/util"
)

var testInput = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	var l1, l2 []int
	for _, line := range util.ReadInput(input, "\n") {
		split := strings.Split(line, "   ")
		l1 = append(l1, util.Str2Int(split[0]))
		l2 = append(l2, util.Str2Int(split[1]))
	}

	fmt.Printf("first: %d\n", first(l1, l2))
	fmt.Printf("second: %d\n", second(l1, l2))
}

func first(l1, l2 []int) int {
	slices.Sort(l1)
	slices.Sort(l2)

	var diff int
	for i := range l1 {
		diff += maths.AbsInt(l1[i] - l2[i])
	}

	return diff
}

func second(l1, l2 []int) int {
	m := make(map[int]int)
	for _, i := range l2 {
		m[i] += 1
	}

	var diff int
	for _, i := range l1 {
		diff += i * m[i]
	}

	return diff
}
