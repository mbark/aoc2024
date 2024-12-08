package day6

import (
	"fmt"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/maps"
)

const testInput = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

const (
	Empty    = '.'
	Obstacle = '#'
	Guard    = '^'
)

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	var guard maps.Coordinate
	m := maps.New(input, func(x, y int, b byte) byte {
		if b == Guard {
			guard = maps.C(x, y)
		}
		return b
	})

	fmt.Printf("first: %d\n", first(m, guard))
	fmt.Printf("second: %d\n", second(m, guard))
}

func first(m maps.Map[byte], at maps.Coordinate) int {
	return len(walk(m, at))
}

func second(m maps.Map[byte], at maps.Coordinate) int {
	var found int
	for _, c := range walk(m, at) {
		switch m.At(c) {
		case Guard:
			continue
		case Obstacle:
			continue
		default:
			m.Set(c, Obstacle)
			if isCycle(m, at) {
				found += 1
			}
			m.Set(c, Empty)
		}
	}

	return found
}

func walk(m maps.Map[byte], at maps.Coordinate) []maps.Coordinate {
	visited := map[maps.Coordinate]bool{}
	dir := maps.Up
	for {
		if !m.Exists(at) {
			break
		}

		visited[at] = true
		next := dir.Apply(at)
		if m.AtSafe(next) == Obstacle {
			dir = dir.Rotate(maps.Right)
			continue
		} else {
			at = next
		}

	}

	return fns.Keys(visited)
}

type visit struct {
	coord maps.Coordinate
	dir   maps.Direction
}

func isCycle(m maps.Map[byte], at maps.Coordinate) bool {
	visited := map[visit]bool{}
	dir := maps.Up
	for {
		if !m.Exists(at) {
			return false
		}
		if visited[visit{at, dir}] {
			return true
		}

		visited[visit{coord: at, dir: dir}] = true
		next := dir.Apply(at)
		if m.AtSafe(next) == Obstacle {
			dir = dir.Rotate(maps.Right)
			continue
		} else {
			at = next
		}
	}
}
