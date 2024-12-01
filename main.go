package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mbark/aoc2024/day1"
	"github.com/mbark/aoc2024/day10"
	"github.com/mbark/aoc2024/day11"
	"github.com/mbark/aoc2024/day12"
	"github.com/mbark/aoc2024/day13"
	"github.com/mbark/aoc2024/day14"
	"github.com/mbark/aoc2024/day15"
	"github.com/mbark/aoc2024/day16"
	"github.com/mbark/aoc2024/day17"
	"github.com/mbark/aoc2024/day18"
	"github.com/mbark/aoc2024/day19"
	"github.com/mbark/aoc2024/day2"
	"github.com/mbark/aoc2024/day20"
	"github.com/mbark/aoc2024/day21"
	"github.com/mbark/aoc2024/day22"
	"github.com/mbark/aoc2024/day23"
	"github.com/mbark/aoc2024/day24"
	"github.com/mbark/aoc2024/day25"
	"github.com/mbark/aoc2024/day3"
	"github.com/mbark/aoc2024/day4"
	"github.com/mbark/aoc2024/day5"
	"github.com/mbark/aoc2024/day6"
	"github.com/mbark/aoc2024/day7"
	"github.com/mbark/aoc2024/day8"
	"github.com/mbark/aoc2024/day9"
	"github.com/mbark/aoc2024/util"
)

func main() {
	var (
		flagDay    = flag.Int("day", 0, "use test input")
		flagTest   = flag.Bool("test", false, "use test input")
		cpuprofile = flag.Bool("profile", false, "write cpu profile to file")
	)
	flag.Parse()

	if *cpuprofile {
		fmt.Println("using cpu profile")
		fn := util.WithProfiling()
		defer fn()
	}

	var input string
	if !*flagTest {
		input = util.GetInput(*flagDay)
	}

	switch *flagDay {
	case 1:
		day1.Run(input, *flagTest)
	case 2:
		day2.Run(input, *flagTest)
	case 3:
		day3.Run(input, *flagTest)
	case 4:
		day4.Run(input, *flagTest)
	case 5:
		day5.Run(input, *flagTest)
	case 6:
		day6.Run(input, *flagTest)
	case 7:
		day7.Run(input, *flagTest)
	case 8:
		day8.Run(input, *flagTest)
	case 9:
		day9.Run(input, *flagTest)
	case 10:
		day10.Run(input, *flagTest)
	case 11:
		day11.Run(input, *flagTest)
	case 12:
		day12.Run(input, *flagTest)
	case 13:
		day13.Run(input, *flagTest)
	case 14:
		day14.Run(input, *flagTest)
	case 15:
		day15.Run(input, *flagTest)
	case 16:
		day16.Run(input, *flagTest)
	case 17:
		day17.Run(input, *flagTest)
	case 18:
		day18.Run(input, *flagTest)
	case 19:
		day19.Run(input, *flagTest)
	case 20:
		day20.Run(input, *flagTest)
	case 21:
		day21.Run(input, *flagTest)
	case 22:
		day22.Run(input, *flagTest)
	case 23:
		day23.Run(input, *flagTest)
	case 24:
		day24.Run(input, *flagTest)
	case 25:
		day25.Run(input, *flagTest)
	default:
		fmt.Println("not implemented")
		os.Exit(1)
	}
}
