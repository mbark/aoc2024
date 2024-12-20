package day20

import (
	"fmt"
	"slices"

	"github.com/mbark/aoc2024/maps"
)

const testInput = `
###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	var start, end maps.Coordinate
	m := maps.New(input, func(x, y int, b byte) byte {
		switch b {
		case 'S':
			start = maps.C(x, y)
			return '.'
		case 'E':
			end = maps.C(x, y)
			return '.'
		default:
			return b
		}
	})

	fmt.Printf("first: %d\n", first(m, start, end, 2))
	fmt.Printf("first: %d\n", first(m, start, end, 20))
}

func first(m maps.Map[byte], start, end maps.Coordinate, cheatSize int) int {
	path := bfs(m, start, end)

	paths := cheats(path, cheatSize)
	chs := make(map[int]int)
	for _, p := range paths {
		if p >= len(path) {
			continue
		}

		saved := len(path) - p
		chs[saved]++
	}

	var c int
	for saved, count := range chs {
		if saved >= 100 {
			c += count
		}
	}

	return c
}

func cheats(path []maps.Coordinate, cheatSize int) []int {
	var paths []int

	for i, c1 := range path {
		for j, c2 := range path {
			if j <= i {
				continue
			}

			distance := c1.ManhattanDistance(c2)
			if distance > cheatSize {
				continue
			}

			paths = append(paths, i+distance+len(path)-j)
		}
	}

	return paths
}

func bfs(m maps.Map[byte], start, end maps.Coordinate) []maps.Coordinate {
	queue := []maps.Coordinate{start}
	visited := map[maps.Coordinate]bool{start: true}
	from := map[maps.Coordinate]maps.Coordinate{}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if next == end {
			var path []maps.Coordinate
			for at := end; at != start; at = from[at] {
				path = append(path, at)
			}
			path = append(path, start)
			slices.Reverse(path)
			return path
		}

		for _, n := range m.Adjacent(next) {
			if visited[n] || m.At(n) == '#' {
				continue
			}

			visited[n] = true
			from[n] = next
			queue = append(queue, n)
		}
	}

	return nil
}
