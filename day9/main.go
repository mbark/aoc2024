package day9

import (
	"fmt"
	"strings"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/util"
)

const testInput = `2333133121414131402`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	var blocks []*block
	var id block
	for i, b := range util.Str2IntSlice(strings.Split(input, "")) {
		if i%2 == 0 {
			idd := id
			blocks = append(blocks, fns.Repeat(&idd, b)...)
			id++
		} else {
			blocks = append(blocks, fns.Repeat[*block](nil, b)...)
		}

	}

	blocks1 := util.CopyList(blocks)
	blocks2 := util.CopyList(blocks)
	fmt.Printf("first: %d\n", first(blocks1))
	fmt.Printf("second: %d\n", second(blocks2))
}

type block int

func (b *block) String() string {
	if b == nil {
		return "."
	}

	return fmt.Sprintf("%d", *b)
}

func first(blocks []*block) int {
	for i := len(blocks) - 1; i >= 0; i-- {
		for j := 0; j < len(blocks) && j < i; j++ {
			if blocks[j] == nil {
				blocks[j] = blocks[i]
				blocks[i] = nil
			}
		}
	}

	var sum int
	for i, b := range blocks {
		if b == nil {
			continue
		}
		sum += i * int(*b)

	}
	return sum
}

func second(blocks []*block) int {
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i] == nil {
			continue
		}

		var size int
		b := blocks[i]
		for k := i; k >= 0; k-- {
			if blocks[k] == nil {
				break
			}
			if *blocks[k] != *b {
				break
			}
			size++
		}

		for j := 0; j < len(blocks) && j < i; {
			var k int
			for ; k < size; k++ {
				if blocks[j+k] != nil {
					break
				}
			}
			if k >= size {
				for k := 0; k < size; k++ {
					blocks[j+k] = blocks[i-k]
					blocks[i-k] = nil
				}
				break
			}
			j += k + 1
		}
		i -= size - 1
	}

	var sum int
	for i, b := range blocks {
		if b == nil {
			continue
		}
		sum += i * int(*b)

	}
	return sum
}

func debugList(blocks []*block) {
	fmt.Println(strings.Join(fns.Map(blocks, func(b *block) string { return b.String() }), ""))
}
