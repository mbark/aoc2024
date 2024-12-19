package day19

import (
	"fmt"
	"strings"

	"github.com/mbark/aoc2024/util"
)

const testInput = `
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	var towels []string
	var designs []string
	split := util.ReadInput(input, "\n\n")
	for _, v := range strings.Split(split[0], ", ") {
		towels = append(towels, v)
	}
	for _, v := range strings.Split(split[1], "\n") {
		designs = append(designs, v)
	}

	fmt.Printf("first: %d\n", first(towels, designs))
	fmt.Printf("second: %d\n", second(towels, designs))

}

func first(towels []string, designs []string) int {
	var possible int
	for _, d := range designs {
		if solve(towels, d) > 0 {
			possible++
		}
	}

	return possible
}
func second(towels []string, designs []string) int {
	var possible int
	for _, d := range designs {
		possible += solve(towels, d)
	}

	return possible
}

var memo = make(map[string]int)

func solve(towels []string, design string) (c int) {
	if len(design) == 0 {
		return 1
	}
	if v, ok := memo[design]; ok {
		return v
	}
	defer func() { memo[design] = c }()

	var possible int
	for _, t := range towels {
		if strings.HasPrefix(design, t) {
			possible += solve(towels, design[len(t):])
		}
	}

	return possible
}
