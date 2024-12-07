package day3

import (
	"fmt"
	"regexp"

	"github.com/mbark/aoc2024/util"
)

const testInput = `
xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
`

var (
	re     = regexp.MustCompile(`mul\((?P<num1>\d+),(?P<num2>\d+)\)`)
	dontre = regexp.MustCompile(`don't\(\)`)
	dore   = regexp.MustCompile(`do\(\)`)
)

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	fmt.Printf("first: %d\n", first(input))
	fmt.Printf("second: %d\n", second(input))
}

func first(input string) int {
	numbers := util.RegexCaptureGroups(re, input)
	var sum int
	for _, n := range numbers {
		sum += util.ParseInt[int](n["num1"]) * util.ParseInt[int](n["num2"])
	}

	return sum
}

func second(input string) int {
	enabled := map[int]bool{}
	isEnabled := true
	var idx int
	for {
		var next []int
		switch isEnabled {
		case true:
			next = dontre.FindStringIndex(input[idx:])
		case false:
			next = dore.FindStringIndex(input[idx:])
		}
		if next == nil {
			break
		}

		end := idx + next[1]
		for ; idx < end; idx++ {
			enabled[idx] = isEnabled
		}
		isEnabled = !isEnabled
	}
	for ; idx < len(input); idx++ {
		enabled[idx] = isEnabled
	}

	matches := re.FindAllStringSubmatch(input, -1)
	matchIndexes := re.FindAllStringSubmatchIndex(input, -1)
	groupNames := re.SubexpNames()

	var results []map[string]string
	for at, match := range matches {
		result := make(map[string]string)
		for i, name := range groupNames {
			idx := matchIndexes[at][i]
			if !enabled[idx] {
				continue
			}

			if i > 0 && name != "" {
				result[name] = match[i]
			}
		}

		if len(result) == 2 {
			results = append(results, result)
		}
	}

	var sum int
	for _, n := range results {
		sum += util.ParseInt[int](n["num1"]) * util.ParseInt[int](n["num2"])
	}

	return sum
}
