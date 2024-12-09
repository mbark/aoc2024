package day7

import (
	"fmt"
	"strings"

	"github.com/mbark/aoc2024/util"
)

const testInput = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	var sums []test
	for _, l := range util.ReadInput(input, "\n") {
		split := strings.Split(l, ": ")
		s := util.Str2Int(split[0])
		parts := util.Str2IntSlice(strings.Split(split[1], " "))
		sums = append(sums, test{s, parts})
	}

	fmt.Printf("first: %d\n", first(sums))
	fmt.Printf("second: %d\n", second(sums))
}

type test struct {
	sum   int
	parts []int
}

func first(sums []test) int {
	var res int
	for _, t := range sums {
		if part1(t.sum, t.parts[0], t.parts[1:]) {
			res += t.sum
		}
	}

	return res
}

func second(sums []test) int {
	var res int
	for _, t := range sums {
		if part2(t.sum, t.parts[0], t.parts[1:]) {
			res += t.sum
		}
	}

	return res
}

func part1(target, current int, parts []int) bool {
	if len(parts) == 0 {
		return target == current
	}
	if current > target {
		return false
	}

	p := parts[0]
	next := parts[1:]

	return part1(target, current+p, next) || part1(target, current*p, next)
}

func part2(target, current int, parts []int) bool {
	if len(parts) == 0 {
		return target == current
	}
	if current > target {
		return false
	}

	p := parts[0]
	next := parts[1:]

	concat := util.Str2Int(fmt.Sprintf("%d%d", current, p))

	return part2(target, current+p, next) ||
		part2(target, current*p, next) ||
		part2(target, concat, next)
}
