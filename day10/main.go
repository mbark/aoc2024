package day10

import (
	"fmt"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/maps"
)

const testInput = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	zeroes := make(map[maps.Coordinate]bool)
	nines := make(map[maps.Coordinate]bool)
	m := maps.NewIntMap(input)

	for _, c := range m.Coordinates() {
		if m.At(c) == 0 {
			zeroes[c] = true
		}
		if m.At(c) == 9 {
			nines[c] = true
		}
	}

	fmt.Println("first: ", first(m, zeroes, nines))
	fmt.Println("second: ", second(m, zeroes, nines))
}

func first(m maps.Map[int], zeroes, nines map[maps.Coordinate]bool) int {
	var sum int
	for z := range zeroes {
		s1, _ := bfs(z, m, nines)
		sum += s1
	}
	return sum
}

func second(m maps.Map[int], zeroes, nines map[maps.Coordinate]bool) int {
	var sum int
	for z := range zeroes {
		_, s2 := bfs(z, m, nines)
		sum += s2
	}
	return sum
}

func bfs(start maps.Coordinate, m maps.Map[int], ends map[maps.Coordinate]bool) (int, int) {
	queue := []maps.Coordinate{start}
	visited := map[maps.Coordinate]bool{start: true}

	from := make(map[maps.Coordinate][]maps.Coordinate)

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		visited[next] = true

		for _, n := range m.Adjacent(next) {
			if m.At(n) == m.At(next)+1 {
				from[n] = append(from[n], next)
				queue = append(queue, n)
			}
		}
	}

	nines := fns.Filter(fns.Keys(visited), func(c maps.Coordinate) bool {
		return ends[c]
	})

	var sum int
	for _, n := range nines {
		sum += len(from[n])
	}

	return len(nines), sum
}
