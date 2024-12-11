package day8

import (
	"fmt"

	"github.com/mbark/aoc2024/maps"
)

const testInput = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	antennas := make(map[maps.Coordinate]byte)
	m := maps.New(input, func(x, y int, b byte) byte {
		switch b {
		case '.':
		default:
			antennas[maps.C(x, y)] = b
		}

		return b
	})
	fmt.Printf("first: %d\n", first(m, antennas))
	fmt.Printf("second: %d\n", second(m, antennas))
}

func first(m maps.Map[byte], antennas map[maps.Coordinate]byte) int {
	antinodes := make(map[maps.Coordinate]bool)
	byType := make(map[byte][]maps.Coordinate)
	for at, a := range antennas {
		byType[a] = append(byType[a], at)
	}

	for _, as := range byType {
		for i, a1 := range as {
			for _, a2 := range as[i+1:] {
				diff := a1.Sub(a2)
				antinodes[a1.Add(diff)] = true
				antinodes[a2.Add(diff.Neg())] = true
			}
		}
	}

	for c := range antinodes {
		if !m.Exists(c) {
			delete(antinodes, c)
		}
	}

	fmt.Println(m.Stringf(func(c maps.Coordinate, val byte) string {
		if antinodes[c] {
			return "#"
		}

		return string(val)
	}))
	return len(antinodes)
}

func second(m maps.Map[byte], antennas map[maps.Coordinate]byte) int {
	antinodes := make(map[maps.Coordinate]bool)
	byType := make(map[byte][]maps.Coordinate)
	for at, a := range antennas {
		byType[a] = append(byType[a], at)
	}

	for _, as := range byType {
		for i, a1 := range as {
			for _, a2 := range as[i+1:] {
				diff := a1.Sub(a2)

				antinodes[a1] = true
				antinodes[a2] = true
				for at := a1.Add(diff); m.Exists(at); at = at.Add(diff) {
					antinodes[at] = true
				}
				for at := a2.Add(diff.Neg()); m.Exists(at); at = at.Add(diff.Neg()) {
					antinodes[at] = true
				}
			}
		}
	}
	for c := range antinodes {
		if !m.Exists(c) {
			delete(antinodes, c)
		}
	}

	fmt.Println(m.Stringf(func(c maps.Coordinate, val byte) string {
		if antinodes[c] {
			return "#"
		}

		return string(val)
	}))
	return len(antinodes)
}
