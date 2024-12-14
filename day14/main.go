package day14

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/maps"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

var re = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	var robots1 []robot
	var robots2 []robot
	var id int
	for _, l := range util.ReadInput(input, "\n") {
		matches := util.RegexSubexps(re, l)
		r := robot{
			id: id,
			at: maps.Coordinate{
				X: util.Str2Int(matches[0]),
				Y: util.Str2Int(matches[1]),
			},
			dir: maps.Coordinate{
				X: util.Str2Int(matches[2]),
				Y: util.Str2Int(matches[3]),
			},
		}
		id++
		robots1 = append(robots1, r)
		robots2 = append(robots2, r)
	}

	fmt.Printf("first: %d\n", first(robots1, isTest))
	fmt.Printf("second: %d\n", second(robots2, isTest))
}

type robot struct {
	id  int
	at  maps.Coordinate
	dir maps.Coordinate
}

func first(robots []robot, isTest bool) int {
	var m maps.Map[[]robot]
	if isTest {
		m = maps.NewEmpty[[]robot](11, 7)
	} else {
		m = maps.NewEmpty[[]robot](101, 103)
	}

	for i := 0; i < 100; i++ {
		for idx, r := range robots {
			robots[idx].at = m.WrapCoordinate(r.at.Add(r.dir))
		}
	}

	var q1, q2, q3, q4 int
	for _, r := range robots {
		x, y := r.at.X, r.at.Y
		switch {
		case x < m.Columns/2 && y < m.Rows/2:
			q1++
		case x > m.Columns/2 && y < m.Rows/2:
			q2++
		case x < m.Columns/2 && y > m.Rows/2:
			q3++
		case x > m.Columns/2 && y > m.Rows/2:
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func second(robots []robot, isTest bool) int {
	var m maps.Map[[]robot]
	if isTest {
		m = maps.NewEmpty[[]robot](11, 7)
	} else {
		m = maps.NewEmpty[[]robot](101, 103)
	}

	var step int
	for ; step <= 10001; step++ {
		positions := map[maps.Coordinate]bool{}

		for idx, r := range robots {
			next := m.WrapCoordinate(r.at.Add(r.dir))

			robots[idx].at = next
			positions[next] = true
		}

		var surrounded int
		for _, r := range robots {
			if fns.Every(m.Surrounding(r.at), func(s maps.Coordinate) bool {
				return positions[s]
			}) {
				surrounded++
			}
		}

		// 10% being surrounded is probably a Christmas tree, right?
		if surrounded > len(robots)/10 {
			printMap(m, robots)
			return step + 1
		}
	}

	return 0
}

func printMap(m maps.Map[[]robot], robots []robot) {
	fmt.Println(m.Stringf(func(c maps.Coordinate, _ []robot) string {
		rs := fns.Filter(robots, func(r robot) bool { return r.at == c })
		if len(rs) == 0 {
			return "."
		}

		return strconv.Itoa(len(rs))
	}))
}
