package day6

import (
	"fmt"

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

	return len(visited)
}

func second(m maps.Map[byte], at maps.Coordinate) int {
	var found int
	for _, c := range m.Coordinates() {
		switch m.At(c) {
		case Guard:
			continue
		case Obstacle:
			continue
		default:
			cp := m.CopyWith(func(c1 maps.Coordinate, val byte) byte {
				if c == c1 {
					return Obstacle
				}
				return val
			})
			if isCycle(cp, at) {
				found += 1
			}
		}
	}

	return found
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
