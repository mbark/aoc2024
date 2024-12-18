package day18

import (
	"fmt"
	"strings"

	"github.com/mbark/aoc2024/maps"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

func Run(input string, isTest bool) {
	var m maps.Map[byte]
	var blocks []maps.Coordinate
	var loops int
	if isTest {
		input = testInput
		m = maps.NewEmpty[byte](7, 7)
		loops = 12
	} else {
		m = maps.NewEmpty[byte](71, 71)
		loops = 1024
	}

	for _, l := range util.ReadInput(input, "\n") {
		split := strings.Split(l, ",")
		c := maps.C(util.Str2Int(split[0]), util.Str2Int(split[1]))
		blocks = append(blocks, c)
	}

	fmt.Printf("first: %d\n", first(m, blocks, loops))
	fmt.Printf("second: %s\n", second(m, blocks, loops))
}

func first(m maps.Map[byte], blocks []maps.Coordinate, loops int) int {
	for i := 0; i < loops; i++ {
		b := blocks[i]
		m.Set(b, '#')
	}

	return bfs(m, maps.C(0, 0), maps.C(m.Rows-1, m.Columns-1))
}

func second(m maps.Map[byte], blocks []maps.Coordinate, minCheck int) string {
	for i := 0; i < len(blocks); i++ {
		b := blocks[i]
		m.Set(b, '#')

		if i >= minCheck {
			path := bfs(m, maps.C(0, 0), maps.C(m.Rows-1, m.Columns-1))
			if path == -1 {
				return fmt.Sprintf("%d,%d", b.X, b.Y)
			}
		}
	}

	return ""
}

func bfs(m maps.Map[byte], start, end maps.Coordinate) int {
	queue := []maps.Coordinate{start}
	visited := map[maps.Coordinate]bool{start: true}
	from := make(map[maps.Coordinate]maps.Coordinate)

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if next == end {
			var count int
			for at := next; at != start; at = from[at] {
				count += 1
			}
			return count
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

	return -1
}
