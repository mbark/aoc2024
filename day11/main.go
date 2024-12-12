package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mbark/aoc2024/util"
)

const testInput = `
125 17
`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	stones := util.Str2IntSlice(strings.Split(strings.TrimSpace(input), " "))
	fmt.Printf("first: %d\n", first(stones))
	fmt.Printf("second: %d\n", second(stones))
}

func first(stones []int) int {
	for i := 0; i < 25; i++ {
		var next []int
		for _, s := range stones {
			digits := strconv.Itoa(s)
			switch {
			case s == 0:
				next = append(next, 1)
			case len(digits)%2 == 0:
				next = append(next, util.Str2Int(digits[:len(digits)/2]))
				next = append(next, util.Str2Int(digits[len(digits)/2:]))
			default:
				next = append(next, s*2024)
			}
		}
		stones = next
	}

	return len(stones)
}

type memoKey struct {
	stone int
	step  int
}

var memo = make(map[memoKey]int)

func second(stones []int) int {
	var count int
	for _, s := range stones {
		count += recurse(s, 75)
	}
	return count
}

func recurse(stone int, steps int) (count int) {
	defer func() {
		memo[memoKey{stone, steps}] = count
	}()
	if steps == 0 {
		return 1
	}
	if v, ok := memo[memoKey{stone, steps}]; ok {
		return v
	}

	digits := strconv.Itoa(stone)
	switch {
	case stone == 0:
		return recurse(1, steps-1)
	case len(digits)%2 == 0:
		return recurse(util.Str2Int(digits[:len(digits)/2]), steps-1) +
			recurse(util.Str2Int(digits[len(digits)/2:]), steps-1)
	default:
		return recurse(stone*2024, steps-1)
	}
}
