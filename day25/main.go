package day25

import (
	"fmt"
	"slices"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/maps"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	var ms []maps.Map[byte]
	for _, l := range util.ReadInput(input, "\n\n") {
		ms = append(ms, maps.NewByte(l))
	}

	var byLocks, byKeys []maps.Map[byte]
	for _, m := range ms {
		if m.At(maps.C(0, 0)) == '#' {
			byLocks = append(byLocks, m)
		} else {
			byKeys = append(byKeys, m)
		}
	}

	var locks, keys [][]int
	for _, m := range byLocks {
		locks = append(locks, columns(m))
	}
	for _, m := range byKeys {
		keys = append(keys, columns(m))
	}

	fmt.Printf("first: %d\n", first(locks, keys))
}

func first(locks, keys [][]int) int {
	var fits int
	for _, lock := range locks {
		for _, key := range keys {
			if !overlaps(lock, key) {
				fits++
			}
		}
	}

	return fits
}

func overlaps(lock, key []int) bool {
	for i := range lock {
		if lock[i]+key[i] > 5 {
			return true
		}
	}

	return false
}

func columns(m maps.Map[byte]) []int {
	count := map[int]int{}
	for c := range m.IterVertical() {
		if m.At(c) == '#' {
			count[c.X] += 1
		}
	}

	keys := fns.Keys(count)
	slices.Sort(keys)

	var result []int
	for _, k := range keys {
		result = append(result, count[k]-1)
	}

	return result
}
