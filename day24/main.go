package day24

import (
	"fmt"
	"os"
	"os/exec"
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

const testInputSmall = `
x00: 0
x01: 1
x02: 0
x03: 1
x04: 0
x05: 1
y00: 0
y01: 0
y02: 1
y03: 1
y04: 0
y05: 1

x00 AND y00 -> z05
x01 AND y01 -> z02
x02 AND y02 -> z01
x03 AND y03 -> z03
x04 AND y04 -> z04
x05 AND y05 -> z00
`

var (
	reIn = regexp.MustCompile(`([a-z0-9]+): (\d+)`)
	reOp = regexp.MustCompile(`([a-z0-9]+) (AND|OR|XOR) ([a-z0-9]+) -> ([a-z0-9]+)`)
)

func Run(input string, isTest bool) {
	if isTest {
		input = testInputSmall
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

	slices.SortFunc(gates, func(a, b gate) int { return strings.Compare(a.out, b.out) })
	fmt.Printf("first: %d\n", first(util.CopyMap(inputs), util.CopyList(gates)))
	swapped := swapGates(util.CopyList(gates))
	// generate the plot
	//plot(util.CopyMap(inputs), util.CopyList(swapped))
	fmt.Printf("second: %s\n", secondCopy(util.CopyList(swapped)))
	swappedIDs := fns.Keys(swaps)
	slices.Sort(swappedIDs)
	fmt.Printf("second: %s\n", strings.Join(swappedIDs, ","))
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

	return value(inputs, "z")
}

func secondCopy(gates []gate) string {
	bySuffix := map[string]gate{}
	for _, g := range gates {
		if strings.HasPrefix(g.out, "z") {
			continue
		}

		suffix := getSuffix(g.out)
		bySuffix[suffix] = g
	}

	zeroInputs := map[string]int{}
	for i := 0; i < 45; i++ {
		zeroInputs[fmt.Sprintf("x%02d", i)] = 0
		zeroInputs[fmt.Sprintf("y%02d", i)] = 0
	}

	inputs := map[string][]map[string]int{}
	for i := 0; i < 45; i++ {
		suffix := fmt.Sprintf("%02d", i)
		if _, ok := inputs[suffix]; !ok {
			inputs[suffix] = []map[string]int{}
		}

		inputs[suffix] = []map[string]int{
			{fmt.Sprintf("y%02d", i): 1},
			{fmt.Sprintf("x%02d", i): 1},
			{fmt.Sprintf("x%02d", i): 1, fmt.Sprintf("y%02d", i): 1},
		}
	}

	keys := fns.Keys(inputs)
	slices.Sort(keys)

	for _, suffix := range keys {
		var isInvalid int
		for _, in := range inputs[suffix] {
			cp := util.CopyMap(zeroInputs)
			in = fns.Union(cp, in)

			x := value(in, "x")
			y := value(in, "y")
			sum := run(in, gates)
			if sum != x+y {
				isInvalid++
				fmt.Printf("mismatch for z%s %d + %d != %d\n", suffix, x, y, sum)
			}
		}

		if isInvalid > 1 {
			fmt.Printf("z%s invalid: %s\n", suffix, buildPath(fmt.Sprintf("z%s", suffix), zeroInputs, gates, 3))
		} else {
			fmt.Printf("z%s valid: %s\n", suffix, buildPath(fmt.Sprintf("z%s", suffix), zeroInputs, gates, 3))
		}
	}

	return ""
}

var swaps = map[string]string{
	"ggn": "z10",
	"grm": "z32",
	"twr": "z39",
	"jcb": "ndw",
}

func swapGates(gates []gate) []gate {
	for k, v := range swaps {
		swaps[v] = k
	}

	for i, g := range gates {
		swap, ok := swaps[g.out]
		if !ok {
			continue
		}

		gates[i].out = swap
	}

	return gates
}

func run(inputs map[string]int, gates []gate) int {
	for {
		hasChanges := false
		for _, g := range gates {
			if _, ok := inputs[g.out]; ok {
				continue
			}

			in1, ok1 := inputs[g.in1]
			in2, ok2 := inputs[g.in2]

			if !ok1 || !ok2 {
				continue
			}

			hasChanges = true
			inputs[g.out] = g.op(in1, in2)
		}

		if !hasChanges {
			break
		}
	}

	return value(inputs, "z")
}

func getSuffix(s string) string {
	num, err := strconv.Atoi(strings.TrimPrefix(s, "z"))
	if err != nil {
		return s
	}

	return fmt.Sprintf("%02d", num)
}

func buildPath(from string, inputs map[string]int, gates []gate, depth int) string {
	if _, ok := inputs[from]; ok {
		return from
	}
	if depth == 0 {
		return from
	}

	for _, g := range gates {
		if g.out != from {
			continue
		}

		return fmt.Sprintf("%s(%s %s %s)", g.out, buildPath(g.in1, inputs, gates, depth-1), g.opName, buildPath(g.in2, inputs, gates, depth-1))
	}

	return ""
}

func value(inputs map[string]int, prefix string) int {
	keys := fns.Keys(inputs)
	keys = fns.Filter(keys, func(k string) bool { return strings.HasPrefix(k, prefix) })
	slices.Sort(keys)
	slices.Reverse(keys)

	var s strings.Builder
	for _, k := range keys {
		s.WriteString(fmt.Sprintf("%d", inputs[k]))
	}
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
func plot(inputs map[string]int, gates []gate) {
	var sb strings.Builder
	sb.WriteString("digraph G {\n")

	var zKeys, xKeys, yKeys []string
	for _, g := range gates {
		if strings.HasPrefix(g.out, "z") {
			zKeys = append(zKeys, g.out)
			xKeys = append(xKeys, g.out)
			yKeys = append(yKeys, g.out)
		}
	}
	sb.WriteString(fmt.Sprintf("subgraph z {rank=same; %s}\n", strings.Join(zKeys, "; ")))
	sb.WriteString(fmt.Sprintf("subgraph x {rank=same; %s}\n", strings.Join(xKeys, "; ")))
	sb.WriteString(fmt.Sprintf("subgraph z {rank=same; %s}\n", strings.Join(yKeys, "; ")))

	for _, g := range gates {
		var color string
		switch g.opName {
		case "AND":
			color = "red"
		case "OR":
			color = "green"
		case "XOR":
			color = "blue"
		}

		sb.WriteString(fmt.Sprintf("{%s;%s} -> %s [label=\"%s\"]\n", g.in1, g.in2, g.out, g.opName))
		sb.WriteString(fmt.Sprintf("%s [shape=circle, style=filled, fillcolor=%s]", g.out, color))
	}
	for i := range inputs {
		sb.WriteString(fmt.Sprintf("%s [shape=box]\n", i))
	}
	sb.WriteString("}\n")

	_ = os.WriteFile("day24.dot", []byte(sb.String()), 0644)
	err := exec.Command("dot", "-Tsvg", "day24.dot", "-o", "day24.svg").Run()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	err = exec.Command("open", "day24.svg").Run()
	if err != nil {
		panic(err)
	}
}
