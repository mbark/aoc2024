package day21

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/maps"
	"github.com/mbark/aoc2024/util"
)

type memoKey struct {
	pos maps.Coordinate
	key string
}

type Memo map[memoKey][]moves

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

	fmt.Println("first:", first(util.ReadInput(input, "\n")))
	fmt.Println("second:", second(util.ReadInput(input, "\n")))
}

var re = regexp.MustCompile(`[0-9]+`)

func first(inputs []string) int {
	var sum int
	for _, in := range inputs {
		presses := pressKey(Memo{}, numKeyMapping, numKeyPad, maps.C(numKeyPad.Columns-1, numKeyPad.Rows-1), in)
		var states []string
		var nextStates []string

		setState := func() { states, nextStates = filterShortest(nextStates, func(t string) int { return len(t) }), nil }

		for _, p := range presses {
			slices.Reverse(p)
			nextStates = append(nextStates, p.flatten())
		}
		setState()

		for _, s := range states {
			presses := pressKey(memo[0], dirKeyMapping, dirKeyPad, maps.C(dirKeyPad.Columns-1, 0), s)
			for _, p := range presses {
				slices.Reverse(p)
				nextStates = append(nextStates, p.flatten())
			}
		}
		setState()

		for _, s := range states {
			presses := pressKey(memo[1], dirKeyMapping, dirKeyPad, maps.C(dirKeyPad.Columns-1, 0), s)
			for _, p := range presses {
				slices.Reverse(p)
				nextStates = append(nextStates, p.flatten())
			}
		}
		setState()

		shortest := len(states[0])
		i := util.Str2Int(re.FindString(in))
		fmt.Println(in, "->", shortest, "*", i, "=", shortest*i)
		sum += shortest * i
	}

	return sum
}

func second(inputs []string) int {
	var sum int
	for _, in := range inputs {
		presses := pressKey(Memo{}, numKeyMapping, numKeyPad, maps.C(numKeyPad.Columns-1, numKeyPad.Rows-1), in)
		var states []string
		var nextStates []string

		setState := func() { states, nextStates = filterShortest(nextStates, func(t string) int { return len(t) }), nil }

		for _, p := range presses {
			slices.Reverse(p)
			nextStates = append(nextStates, p.flatten())
		}
		setState()

		for i := 0; i < 25; i++ {
			fmt.Println("robot", i)
			for _, s := range states {
				presses := pressKey(memo[i], dirKeyMapping, dirKeyPad, maps.C(dirKeyPad.Columns-1, 0), s)
				for _, p := range presses {
					slices.Reverse(p)
					nextStates = append(nextStates, p.flatten())
				}
			}
			setState()
		}

		shortest := len(states[0])
		i := util.Str2Int(re.FindString(in))
		fmt.Println(in, "->", shortest, "*", i, "=", shortest*i)
		sum += shortest * i
	}

	return sum
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

type moves []string

func (m moves) flatten() string {
	return strings.Join(m, "")
}

var memo = map[int]Memo{}

func pressKey(memo Memo, mapping Mapping, m maps.Map[byte], at maps.Coordinate, keys string) (ret []moves) {
	if v, ok := memo[memoKey{pos: at, key: keys}]; ok {
		return v
	}
	defer func() { memo[memoKey{pos: at, key: keys}] = ret }()

	if len(keys) == 0 {
		return []moves{{}}
	}

	var ms []moves
	k := keys[0]
	next := keys[1:]

	paths := mapping[m.At(at)][k]
	if m.At(at) == k {
		paths = [][]maps.Direction{nil}
	}

	for _, dirs := range paths {
		newAt := at.Apply(dirs...)
		for _, mv := range pressKey(memo, mapping, m, newAt, next) {
			ms = append(ms, append(util.CopyList(mv), dirString(dirs)+"A"))
		}
	}

	return filterShortest(ms, func(m moves) int { return len(m) })
}

func dirString(dirs []maps.Direction) string {
	return strings.Join(fns.Map(dirs, func(d maps.Direction) string { return d.String() }), "")
}

func fillMaps() {
	fillMap(numKeyPad, numKeyMapping)
	fillMap(dirKeyPad, dirKeyMapping)

	aPos := maps.C(dirKeyPad.Columns-1, dirKeyPad.Rows-1)
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
				paths := pressKey(Memo{}, dirKeyMapping, dirKeyPad, aPos, dirString(m)+"A")
				if len(paths[0].flatten()) < shortest {
					shortest = len(paths[0].flatten())
					shortestMappings = nil
				}

				if len(paths[0].flatten()) == shortest {
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
