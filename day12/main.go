package day12

import (
	"fmt"

	"github.com/mbark/aoc2024/maps"
)

const testInput = `
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	m := maps.NewByte(input)
	price, areas := first(m)
	fmt.Printf("first: %d\n", price)
	fmt.Printf("second: %d\n", second(m, areas))
}

func first(m maps.Map[byte]) (int, []map[maps.Coordinate]bool) {
	skip := make(map[maps.Coordinate]bool)
	var areas []map[maps.Coordinate]bool
	var price int
	for _, c := range m.Coordinates() {
		if skip[c] {
			continue
		}

		visited := bfs(c, m)
		areas = append(areas, visited)
		area := len(visited)
		var perimeter int
		for v := range visited {
			skip[v] = true
			for _, n := range v.Adjacent() {
				if !visited[n] || !m.Exists(n) {
					perimeter++
				}
			}
		}

		price += area * perimeter
	}

	return price, areas
}

func second(m maps.Map[byte], areas []map[maps.Coordinate]bool) int {
	var price int

	for _, area := range areas {
		var perimeters int

		for y := -1; y < m.Rows; y++ {
			var fenceIn, fenceOut bool
			for x := 0; x < m.Columns; x++ {
				curr := maps.Coordinate{X: x, Y: y}
				below := maps.Down.Apply(curr)

				inwards := area[curr] && !area[below]
				outwards := !area[curr] && area[below]

				if inwards && !fenceIn {
					perimeters++
					fenceIn = true
				}
				if outwards && !fenceOut {
					perimeters++
					fenceOut = true
				}
				if !inwards {
					fenceIn = false
				}
				if !outwards {
					fenceOut = false
				}
			}
		}

		for x := -1; x < m.Columns; x++ {
			var fenceIn, fenceOut bool
			for y := 0; y < m.Rows; y++ {
				curr := maps.Coordinate{X: x, Y: y}
				right := maps.Right.Apply(curr)

				inwards := area[curr] && !area[right]
				outwards := !area[curr] && area[right]

				if inwards && !fenceIn {
					perimeters++
					fenceIn = true
				}
				if outwards && !fenceOut {
					perimeters++
					fenceOut = true
				}
				if !inwards {
					fenceIn = false
				}
				if !outwards {
					fenceOut = false
				}
			}
		}

		price += len(area) * perimeters
	}

	return price
}

func bfs(start maps.Coordinate, m maps.Map[byte]) map[maps.Coordinate]bool {
	queue := []maps.Coordinate{start}
	visited := map[maps.Coordinate]bool{start: true}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		for _, n := range m.Adjacent(next) {
			if !visited[n] && m.At(start) == m.At(n) {
				queue = append(queue, n)
				visited[n] = true
			}
		}
	}

	return visited
}
