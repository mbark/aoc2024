package day5

import (
	"fmt"
	"strings"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	ordering := map[int][]int{}
	var pages [][]int

	sections := util.ReadInput(input, "\n\n")
	for _, line := range util.ReadInput(sections[0], "\n") {
		split := util.Str2IntSlice(strings.Split(line, "|"))
		ordering[split[0]] = append(ordering[split[0]], split[1:]...)
	}

	for _, line := range util.ReadInput(sections[1], "\n") {
		split := util.Str2IntSlice(strings.Split(line, ","))
		pages = append(pages, split)
	}

	fmt.Printf("first: %d\n", first(ordering, pages))
	fmt.Printf("second: %d\n", second(ordering, pages))
}

func first(before map[int][]int, pages [][]int) int {
	after := map[int][]int{}
	for k, v := range before {
		for _, vv := range v {
			after[vv] = append(after[vv], k)
		}
	}

	var sum int
	for _, ps := range pages {
		if isValid(after, ps) {
			// the middle element is the one we want to sum
			sum += ps[len(ps)/2]
		}
	}
	return sum
}

func second(before map[int][]int, pages [][]int) int {
	after := map[int][]int{}
	for k, v := range before {
		for _, vv := range v {
			after[vv] = append(after[vv], k)
		}
	}

	var sum int
	for _, ps := range pages {
		if isValid(after, ps) {
			continue
		}

		validList := recurse(before, after, nil, ps)
		sum += validList[len(validList)/2]
	}

	return sum
}

func recurse(before, after map[int][]int, current []int, other []int) []int {
	if len(other) == 0 {
		return current
	}

	for _, o := range other {
		if isPossible(before, after, current, other, o) {
			curr := util.CopyList(current)
			curr = append(curr, o)
			l := recurse(before, after, curr, fns.Filter(other, func(t int) bool { return t != o }))
			if l != nil {
				return l
			}
		}
	}

	return nil
}

func isPossible(before, after map[int][]int, current, other []int, page int) bool {
	for _, prev := range current {
		if fns.Some(after[prev], func(t int) bool { return page == t }) {
			return false
		}
	}
	for _, next := range other {
		if fns.Some(before[next], func(t int) bool { return page == t }) {
			return false
		}
	}

	return true
}

func isValid(after map[int][]int, pages []int) bool {
	for i, page := range pages {
		for _, prev := range pages[0:i] {
			if fns.Some(after[prev], func(t int) bool { return page == t }) {
				return false
			}
		}
	}

	return true
}
