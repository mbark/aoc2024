package day16

import (
	"container/heap"
	"fmt"

	"github.com/mbark/aoc2024/maps"
	"github.com/mbark/aoc2024/queue"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################
`

func Run(input string, isTest bool) {
	var start, end maps.Coordinate
	if isTest {
		input = testInput
	}
	m := maps.New(input, func(x, y int, b byte) byte {
		if b == 'S' {
			start = maps.C(x, y)
			return '.'
		}
		if b == 'E' {
			end = maps.C(x, y)
			return '.'
		}
		return b
	})

	best := djikstra(m, start, end)
	fmt.Printf("first: %d\n", best)
	fmt.Printf("second: %d\n", djistra2(m, start, end, best))
}

type pos struct {
	at  maps.Coordinate
	dir maps.Direction
}

type posPath struct {
	pos  pos
	path []maps.Coordinate
}

func (p pos) String() string {
	return fmt.Sprintf("%v %v", p.at, p.dir)
}

var memo = map[pos]int{}

func djikstra(m maps.Map[byte], start, end maps.Coordinate) int {
	pq := queue.PriorityQueue[pos]{}
	heap.Init(&pq)
	heap.Push(&pq, &queue.Item[pos]{Value: pos{at: start, dir: maps.East}, Priority: 0})

	visited := map[pos]bool{pos{at: start, dir: maps.East}: true}
	for len(pq) > 0 {
		q := heap.Pop(&pq).(*queue.Item[pos])
		p := q.Value

		if p.at == end {
			return q.Priority
		}

		if prio, ok := memo[p]; ok && q.Priority > prio {
			continue
		}
		memo[p] = q.Priority

		a := p.dir.Apply(p.at)
		next := pos{at: a, dir: p.dir}
		if !visited[next] && m.Exists(next.at) && m.At(next.at) != '#' {
			heap.Push(&pq, &queue.Item[pos]{Value: next, Priority: q.Priority + 1})
			visited[next] = true
		}

		next = pos{at: p.at, dir: p.dir.Rotate(maps.Left)}
		if !visited[next] {
			heap.Push(&pq, &queue.Item[pos]{Value: next, Priority: q.Priority + 1000})
			visited[next] = true
		}

		next = pos{at: p.at, dir: p.dir.Rotate(maps.Right)}
		if !visited[next] {
			heap.Push(&pq, &queue.Item[pos]{Value: next, Priority: q.Priority + 1000})
			visited[next] = true
		}
	}

	return -1
}

func djistra2(m maps.Map[byte], start, end maps.Coordinate, min int) int {
	pq := queue.PriorityQueue[posPath]{}
	heap.Init(&pq)
	heap.Push(&pq, &queue.Item[posPath]{Value: posPath{pos: pos{at: start, dir: maps.East}, path: []maps.Coordinate{start}}, Priority: 0})

	var paths []maps.Coordinate
	for len(pq) > 0 {
		q := heap.Pop(&pq).(*queue.Item[posPath])
		pp := q.Value
		p := pp.pos

		if p.at == end && q.Priority == min {
			paths = append(paths, pp.path...)
			continue
		}

		// allow for some extra turns
		if prio, ok := memo[pp.pos]; ok && q.Priority > prio+1000 {
			continue
		}
		memo[p] = q.Priority

		nexts := []struct {
			next pos
			prio int
		}{
			{next: pos{at: p.dir.Apply(p.at), dir: p.dir}, prio: q.Priority + 1},
			{next: pos{at: p.at, dir: p.dir.Rotate(maps.Left)}, prio: q.Priority + 1000},
			{next: pos{at: p.at, dir: p.dir.Rotate(maps.Right)}, prio: q.Priority + 1000},
		}

		for _, n := range nexts {
			if !m.Exists(n.next.at) || m.At(n.next.at) == '#' {
				continue
			}

			path := pp.path
			if n.next.at != p.at {
				path = append(util.CopyList(path), n.next.at)
			}

			heap.Push(&pq, &queue.Item[posPath]{Value: posPath{pos: n.next, path: path}, Priority: n.prio})
		}
	}

	visits := make(map[maps.Coordinate]bool)
	for _, e := range paths {
		visits[e] = true
	}
	return len(visits)
}
