package day23

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

type Graph map[string]map[string]bool

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	graph := Graph{}

	for _, l := range util.ReadInput(input, "\n") {
		split := strings.Split(l, "-")
		if _, ok := graph[split[0]]; !ok {
			graph[split[0]] = map[string]bool{}
		}
		graph[split[0]][split[1]] = true
	}

	for k, vs := range graph {
		for v := range vs {
			graph[v][k] = true
		}
	}

	fmt.Printf("first: %d\n", first(graph))
	fmt.Printf("second: %s\n", second(graph))
}

func first(graph Graph) int {
	var sets [][]string
	for k1, vs1 := range graph {
		for k2 := range vs1 {
			for k3 := range vs1 {
				if graph[k2][k3] && graph[k3][k1] {
					sets = append(sets, []string{k1, k2, k3})
				}
			}
		}
	}

	var filtered [][]string
	for _, s := range sets {
		if fns.Some(s, func(t string) bool { return strings.HasPrefix(t, "t") }) {
			filtered = append(filtered, s)
		}
	}
	sets = filtered
	filtered = nil
	for _, s1 := range sets {
		contained := slices.ContainsFunc(filtered, func(s []string) bool {
			return fns.Every(s, func(t string) bool { return slices.Contains(s1, t) })
		})
		if !contained {
			filtered = append(filtered, s1)
		}
	}

	return len(filtered)
}

func second(graph Graph) string {
	nodes := map[string]bool{}
	for k := range graph {
		nodes[k] = true
	}

	cliques := BronKerbosch(map[string]bool{}, nodes, map[string]bool{}, graph)

	var largest Set
	for _, c := range cliques {
		if len(c) > len(largest) {
			largest = c
		}
	}
	var keys []string
	for k := range largest {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	return strings.Join(keys, ",")
}

type Set map[string]bool

func BronKerbosch(r Set, p map[string]bool, x Set, g Graph) []Set {
	if len(p) == 0 && len(x) == 0 {
		return []Set{r}
	}

	var cliques []Set
	for v := range p {
		neighbors := g[v]
		cliques = append(cliques, BronKerbosch(
			fns.Union(r, map[string]bool{v: true}),
			fns.Intersection(p, neighbors),
			fns.Intersection(x, neighbors),
			g)...)
		delete(p, v)
		x[v] = true
	}

	return cliques
}
