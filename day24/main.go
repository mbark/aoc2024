package day24

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj
`

var (
	reIn = regexp.MustCompile(`([a-z0-9]+): (\d+)`)
	reOp = regexp.MustCompile(`([a-z0-9]+) (AND|OR|XOR) ([a-z0-9]+) -> ([a-z0-9]+)`)
)

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	inputs := map[string]int{}
	var gates []gate

	split := util.ReadInput(input, "\n\n")
	for _, s := range util.ReadInput(split[0], "\n") {
		m := reIn.FindStringSubmatch(s)
		if len(m) != 3 {
			continue
		}
		inputs[m[1]] = util.Str2Int(m[2])
	}

	for _, s := range util.ReadInput(split[1], "\n") {
		m := reOp.FindStringSubmatch(s)
		if len(m) != 5 {
			continue
		}

		gates = append(gates, gate{
			in1:    m[1],
			in2:    m[3],
			out:    m[4],
			op:     getOp(m[2]),
			opName: m[2],
		})
	}

	fmt.Printf("first: %d\n", first(inputs, gates))

	var outputs []string
	for _, g := range gates {
		outputs = append(outputs, g.out)
	}
	slices.Sort(outputs)
	fmt.Println(outputs)
}

func first(inputs map[string]int, gates []gate) int {
	allGates := map[string]bool{}
	for k := range inputs {
		allGates[k] = true
	}
	for _, g := range gates {
		allGates[g.out] = true
		allGates[g.in1] = true
		allGates[g.in2] = true
	}

	for len(inputs) < len(allGates) {
		for _, g := range gates {
			if _, ok := inputs[g.out]; ok {
				continue
			}

			in1, ok1 := inputs[g.in1]
			in2, ok2 := inputs[g.in2]

			if !ok1 || !ok2 {
				continue
			}

			inputs[g.out] = g.op(in1, in2)
		}
	}

	keys := fns.Keys(inputs)
	keys = fns.Filter(keys, func(k string) bool { return strings.HasPrefix(k, "z") })
	slices.Sort(keys)
	slices.Reverse(keys)
	var s strings.Builder
	for _, k := range keys {
		s.WriteString(fmt.Sprintf("%d", inputs[k]))
	}
	fmt.Println(s.String())
	i, _ := strconv.ParseInt(s.String(), 2, 64)
	return int(i)
}

type gate struct {
	in1, in2 string
	out      string
	op       func(int, int) int
	opName   string
}

func (g gate) String() string {
	//x02 OR y02 -> z02
	return fmt.Sprintf("%s %s %s -> %s", g.in1, g.opName, g.in2, g.out)
}

func getOp(s string) func(int, int) int {
	switch s {
	case "AND":
		return opAnd
	case "OR":
		return opOr
	case "XOR":
		return opXor
	default:
		panic("unknown op: " + s)
	}
}

func opAnd(a, b int) int {
	return a & b
}

func opOr(a, b int) int {
	return a | b
}

func opXor(a, b int) int {
	return a ^ b
}
