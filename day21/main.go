package day21

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/maps"
	"github.com/mbark/aoc2024/util"
)

var (
	numKeyPad = maps.NewByte(`
789
456
123
_0A`)
	dirKeyPad = maps.NewByte(`
_^A
<v>`)
	numKeyMapping = Mapping{}
	dirKeyMapping = Mapping{}
	reA           = regexp.MustCompile(`A`)
	re            = regexp.MustCompile(`[0-9]+`)
)

const testInput = `
029A
980A
179A
456A
379A
`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	fillMaps()
	for i := 0; i < 25; i++ {
		memo[i] = Memo{}
	}

	fmt.Println("first:", solve(util.ReadInput(input, "\n"), 2))
	fmt.Println("second:", solve(util.ReadInput(input, "\n"), 25))
}

func solve(inputs []string, robots int) int {
	var sum int

	for i := 0; i < robots+1; i++ {
		memo[i] = Memo{}
	}

	for _, in := range inputs {
		states := pressKey(numKeyMapping, 'A', in)
		totals := make([]int, len(states))
		for i, s := range states {
			for _, group := range findGroups(s) {
				totals[i] += recurse(dirKeyMapping, dirKeyPad, group, robots)
			}
		}
		shortest := findMin(totals)

		i := util.Str2Int(re.FindString(in))
		fmt.Println(in, "->", shortest, "*", i, "=", shortest*i)
		sum += shortest * i
	}

	return sum
}

var memo = map[int]Memo{}

func recurse(mapping Mapping, m maps.Map[byte], keys string, depth int) (ret int) {
	if depth == 0 {
		return len(keys)
	}

	if v, ok := memo[depth][keys]; ok {
		return v
	}
	defer func() { memo[depth][keys] = ret }()

	ms := filterShortest(pressKey(mapping, 'A', keys), length)
	totals := make([]int, len(ms))
	for i, s := range ms {
		totals[i] = 0
		for _, g := range findGroups(s) {
			totals[i] += recurse(mapping, m, g, depth-1)
		}
	}

	return findMin(totals)
}

func length(s string) int {
	return len(s)
}

func findMin(i []int) int {
	min := math.MaxInt
	for _, v := range i {
		if v < min {
			min = v
		}
	}

	return min
}

func pressKey(mapping Mapping, at byte, keys string) []string {
	k := keys[0]
	next := keys[1:]

	paths := mapping[at][k]
	if at == k {
		paths = [][]maps.Direction{nil}
	}
	var pMoves []string
	for _, p := range paths {
		pMoves = append(pMoves, dirString(p)+"A")
	}
	if len(next) == 0 {
		return pMoves
	}

	var nextMoves []string
	for _, mv := range filterShortest(pressKey(mapping, k, next), length) {
		for _, p := range pMoves {
			nextMoves = append(nextMoves, p+mv)
		}
	}

	return nextMoves
}

func filterShortest[T any](paths []T, fn func(t T) int) []T {
	shortest := math.MaxInt
	for _, p := range paths {
		if fn(p) < shortest {
			shortest = fn(p)
		}
	}

	return fns.Filter(paths, func(p T) bool { return fn(p) == shortest })
}

func findGroups(s string) []string {
	var groups []string
	for idx := 0; idx < len(s); {
		next := strings.Index(s[idx:], "A")
		if next == -1 {
			groups = append(groups, s[idx:])
			break
		}
		groups = append(groups, s[idx:idx+next+1])
		idx += next + 1
	}

	return groups
}

func dirString(dirs []maps.Direction) string {
	return strings.Join(fns.Map(dirs, func(d maps.Direction) string { return d.String() }), "")
}

func fillMaps() {
	fillMap(numKeyPad, numKeyMapping)
	fillMap(dirKeyPad, dirKeyMapping)

	for c1 := range dirKeyPad.IterHorizontal() {
		for c2 := range dirKeyPad.IterVertical() {
			if c1 == c2 {
				continue
			}

			mapping := dirKeyMapping[dirKeyPad.At(c1)][dirKeyPad.At(c2)]
			if len(mapping) <= 1 {
				continue
			}

			shortest := math.MaxInt
			var shortestMappings []int
			for i, m := range mapping {
				paths := pressKey(dirKeyMapping, 'A', dirString(m)+"A")
				if len(paths[0]) < shortest {
					shortest = len(paths[0])
					shortestMappings = nil
				}

				if len(paths[0]) == shortest {
					shortestMappings = append(shortestMappings, i)
				}
			}

			dirKeyMapping[dirKeyPad.At(c1)][dirKeyPad.At(c2)] = fns.Map(shortestMappings, func(i int) []maps.Direction { return mapping[i] })
		}
	}
}

func fillMap(m maps.Map[byte], mapping Mapping) {
	coords := m.Coordinates()
	for i, c1 := range coords {
		for j, c2 := range coords {
			if j == i {
				continue
			}
			if m.At(c1) == '_' || m.At(c2) == '_' {
				continue
			}
			if _, ok := mapping[m.At(c1)]; !ok {
				mapping[m.At(c1)] = map[byte][][]maps.Direction{}
			}

			paths := buildPaths(m, c1, c2)
			var dirPaths [][]maps.Direction
			for _, p := range paths {
				var dirs []maps.Direction
				for i := 0; i < len(p)-1; i++ {
					dirs = append(dirs, maps.Direction(p[i+1].Sub(p[i])))
				}

				dirPaths = append(dirPaths, dirs)
			}

			mapping[m.At(c1)][m.At(c2)] = dirPaths
		}
	}
}

func buildPaths(m maps.Map[byte], start, end maps.Coordinate) [][]maps.Coordinate {
	queue := [][]maps.Coordinate{{start}}
	var paths [][]maps.Coordinate
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if next[len(next)-1] == end {
			paths = append(paths, next)
			continue
		}

		for n := range m.IterAdjacent(next[len(next)-1]) {
			visited := fns.Some(next, func(c maps.Coordinate) bool { return c == n })
			if visited || m.At(n) == '_' || m.At(n) == m.At(start) {
				continue
			}

			queue = append(queue, append(util.CopyList(next), n))
		}
	}

	minLength := math.MaxInt
	for _, p := range paths {
		if len(p) < minLength {
			minLength = len(p)
		}
	}

	return fns.Filter(paths, func(p []maps.Coordinate) bool { return len(p) == minLength })
}

type Memo map[string]int

type Mapping map[byte]map[byte][][]maps.Direction

func (m Mapping) String() string {
	var sb strings.Builder
	for k, v := range m {
		for k2, v2 := range v {
			vs := strings.Join(fns.Map(v2, func(d []maps.Direction) string { return dirString(d) }), ",")
			sb.WriteString(fmt.Sprintf("%c->%c: %s\n", k, k2, vs))
		}
	}
	return sb.String()
}
