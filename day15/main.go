package day15

import (
	"fmt"
	"strings"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/maps"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

const testSmall = `
#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^
`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	split := util.ReadInput(input, "\n\n")

	var start maps.Coordinate
	m := maps.New(split[0], func(x, y int, b byte) byte {
		if b == '@' {
			start = maps.C(x, y)
			return '.'
		}
		return b
	})
	split[1] = strings.Join(split[1:], "")
	split[1] = strings.TrimSpace(split[1])
	dirs := fns.FilterMap(strings.Split(split[1], ""), func(s string) (maps.Direction, bool) {
		if s == "\n" {
			return maps.Up, false
		}
		return maps.DirectionFromString(s), true
	})

	secondMap := split[0]
	secondMap = strings.ReplaceAll(secondMap, "#", "##")
	secondMap = strings.ReplaceAll(secondMap, ".", "..")
	secondMap = strings.ReplaceAll(secondMap, "O", "[]")
	secondMap = strings.ReplaceAll(secondMap, "@", "@.")

	fmt.Printf("first: %d\n", first(m, start, dirs))

	m = maps.New(secondMap, func(x, y int, b byte) byte {
		if b == '@' {
			start = maps.C(x, y)
			return '.'
		}
		return b
	})
	fmt.Printf("second: %d\n", second(m, start, dirs))
}

const (
	Empty = '.'
	Wall  = '#'
	Box   = 'O'
	BoxL  = '['
	BoxR  = ']'
)

func first(m maps.Map[byte], at maps.Coordinate, moves []maps.Direction) int {
	for _, d := range moves {
		next := d.Apply(at)
		switch m.At(next) {
		case Empty:
			at = next
		case Wall:
		case Box:
			end := canMove(m, next, d)
			if end == nil { // not possible
				continue
			}

			back := d.Opposite()
			for c := *end; c != at; c = back.Apply(c) {
				m.Move(back.Apply(c), c, Empty)
			}
			at = next
		}
	}

	var sum int
	for c := range m.IterVertical() {
		if m.At(c) == Box {
			sum += 100*c.Y + c.X
		}
	}
	return sum
}

func second(m maps.Map[byte], at maps.Coordinate, moves []maps.Direction) int {
	for _, d := range moves {
		next := d.Apply(at)

		switch m.At(next) {
		case Empty:
			at = next
		case Wall:
		case BoxR, BoxL:
			cp := m.CopyWith(func(c maps.Coordinate, val byte) byte { return val })
			if move(cp, at, d) {
				m = cp
				at = next
			}
		}
	}

	var sum int
	for c := range m.IterVertical() {
		if m.At(c) == BoxL {
			sum += 100*c.Y + c.X
		}
	}
	return sum
}

func move(m maps.Map[byte], at maps.Coordinate, dir maps.Direction) bool {
	next := dir.Apply(at)
	v := m.At(next)
	switch v {
	case Empty:
		m.Move(at, next, Empty)
		return true

	case Wall:
		return false
	}

	var other maps.Coordinate
	if v == BoxR {
		other = maps.Left.Apply(next)
	} else {
		other = maps.Right.Apply(next)
	}

	switch {
	case dir == maps.Up || dir == maps.Down:
		if move(m, next, dir) && move(m, other, dir) {
			m.Move(at, dir.Apply(at), Empty)
			return true
		}

	case dir == maps.Left || dir == maps.Right:
		if move(m, next, dir) {
			m.Move(at, dir.Apply(at), Empty)
			return true
		}
	}

	return false
}

func printMap(m maps.Map[byte], at maps.Coordinate) {
	fmt.Println(m.Stringf(func(c maps.Coordinate, val byte) string {
		if c == at {
			return "@"
		}
		return string(val)
	}))
}

func canMove(m maps.Map[byte], at maps.Coordinate, d maps.Direction) *maps.Coordinate {
	next := at
	for {
		next = d.Apply(next)
		switch m.At(next) {
		case Box, BoxL, BoxR:
		case Empty:
			return &next
		case Wall:
			return nil
		}
	}
}
