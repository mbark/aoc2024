package day17

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/maths"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`

var re = regexp.MustCompile(`Register ([A-Z]): (\d+)`)
var re2 = regexp.MustCompile(`Program: (.+)`)

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	matches := util.RegexSubexps(re, input)

	registers := make(map[string]int)
	for i := 0; i < len(matches); i += 2 {
		registers[matches[i]] = util.Str2Int(matches[i+1])
	}

	matches = util.RegexSubexps(re2, input)
	ops := util.Str2IntSlice(strings.Split(matches[0], ","))

	fmt.Printf("first: %s\n", strings.Join(fns.Map(first(ops, registers), strconv.Itoa), ","))
	fmt.Printf("second: %d\n", second(ops, registers))
}

var verbose = false

func log(args ...any) {
	if verbose {
		fmt.Println(args...)
	}
}

func first(ops []int, registers map[string]int) []int {
	var output []int

	for i := 0; i < len(ops); {
		op := ops[i]

		combo := func() int {
			switch ops[i+1] {
			case 0, 1, 2, 3:
				return ops[i+1]
			case 4:
				return registers["A"]
			case 5:
				return registers["B"]
			case 6:
				return registers["C"]
			}
			panic("invalid combo")
		}
		literal := func() int {
			return ops[i+1]
		}

		switch op {
		case 0:
			log("adv A=", registers["A"], "/", maths.PowInt(2, combo()))
			registers["A"] = registers["A"] / maths.PowInt(2, combo())
		case 1:
			log("bxl B=", registers["B"], "^", literal())
			registers["B"] = registers["B"] ^ literal()
		case 2:
			log("bst B=", combo(), "%", 8)
			registers["B"] = combo() % 8
		case 3:
			if registers["A"] != 0 {
				log("jnz jmp", literal())
				i = literal()
				continue
			}
		case 4:
			log("bxc B=", registers["B"], "^", registers["C"])
			registers["B"] = registers["B"] ^ registers["C"]
		case 5:
			log("out ", combo(), "%", 8)
			output = append(output, combo()%8)
		case 6:
			log("bdv C=", registers["A"], "/", maths.PowInt(2, combo()))
			registers["B"] = registers["A"] / maths.PowInt(2, combo())
		case 7:
			log("cdv C=", registers["A"], "/", maths.PowInt(2, combo()))
			registers["C"] = registers["A"] / maths.PowInt(2, combo())
		}

		i += 2
	}

	return output
}

func second(ops []int, registers map[string]int) int {
	return solve(ops, registers, 0, len(ops)-1)
}

func solve(ops []int, registers map[string]int, start, idx int) int {
	for add := 0; add <= 8; add++ {
		registers["A"] = start*8 + add
		solved := first(ops, registers)

		var neq bool
		for i := range solved {
			if i >= len(ops[idx:]) {
				neq = true
				break
			}
			if solved[i] != ops[idx:][i] {
				neq = true
			}
		}
		if neq {
			continue
		}

		if idx == 0 {
			return start*8 + add
		}

		res := solve(ops, registers, start*8+add, idx-1)
		if res > 0 {
			return res
		}
	}

	return 0
}
