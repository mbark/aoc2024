package day4

import (
	"fmt"

	"github.com/mbark/aoc2024/maps"
)

const testInput = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	m := maps.New(input, func(x, y int, b byte) byte {
		if b == 'X' || b == 'M' || b == 'A' || b == 'S' {
			return b
		}

		return '.'
	})

	fmt.Println(m.String())
	fmt.Println("first:", first(m))
	fmt.Println("second:", second(m))
}

func first(m maps.Map[byte]) int {
	var count int
	for _, c := range m.Coordinates() {
		if m.At(c) != 'X' {
			continue
		}

		directions := []maps.Direction{
			maps.North,
			maps.NorthEast,
			maps.East,
			maps.SouthEast,
			maps.South,
			maps.SouthWest,
			maps.West,
			maps.NorthWest,
		}
		for _, d := range directions {
			if m.AtSafe(d.Apply(c)) == 'M' && m.AtSafe(d.ApplyN(c, 2)) == 'A' && m.AtSafe(d.ApplyN(c, 3)) == 'S' {
				count += 1
			}
		}
	}

	return count
}

func second(m maps.Map[byte]) int {
	var mases []maps.Coordinate
	for _, c := range m.Coordinates() {
		if m.At(c) != 'M' {
			continue
		}

		directions := []maps.Direction{
			maps.NorthEast,
			maps.SouthEast,
			maps.SouthWest,
			maps.NorthWest,
		}
		for _, d := range directions {
			if m.AtSafe(d.ApplyN(c, 1)) == 'A' && m.AtSafe(d.ApplyN(c, 2)) == 'S' {
				mases = append(mases, d.ApplyN(c, 1))
			}
		}
	}

	var count int
	for i, c1 := range mases {
		for j, c2 := range mases {
			if i <= j {
				continue
			}

			if c1 == c2 {
				count += 1
			}
		}
	}

	return count
}
