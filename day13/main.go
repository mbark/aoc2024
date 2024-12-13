package day13

import (
	"fmt"
	"strings"

	"github.com/mbark/aoc2024/maps"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	var arcades []arcade
	for _, group := range util.ReadInput(input, "\n\n") {
		var a arcade
		split := strings.Split(group, "\n")
		split[0] = strings.TrimPrefix(split[0], "Button A: ")
		split[1] = strings.TrimPrefix(split[1], "Button B: ")
		split[2] = strings.TrimPrefix(split[2], "Prize: ")

		btnA := strings.Split(split[0], ", ")
		btnB := strings.Split(split[1], ", ")
		prize := strings.Split(split[2], ", ")
		a.a = maps.Coordinate{
			X: util.Str2Int(strings.TrimPrefix(btnA[0], "X+")),
			Y: util.Str2Int(strings.TrimPrefix(btnA[1], "Y+")),
		}
		a.b = maps.Coordinate{
			X: util.Str2Int(strings.TrimPrefix(btnB[0], "X+")),
			Y: util.Str2Int(strings.TrimPrefix(btnB[1], "Y+")),
		}
		a.prize = maps.Coordinate{
			X: util.Str2Int(strings.TrimPrefix(prize[0], "X=")),
			Y: util.Str2Int(strings.TrimPrefix(prize[1], "Y=")),
		}

		arcades = append(arcades, a)
	}

	fmt.Printf("first: %d\n", first(arcades))
	fmt.Printf("second: %d\n", second(arcades))
}

func first(arcades []arcade) int {
	var cost int
	for _, a := range arcades {
		// A*ax + B*bx = cx
		// A*ay + B*by = cy
		ax, ay := a.a.X, a.a.Y
		bx, by := a.b.X, a.b.Y
		cx, cy := a.prize.X, a.prize.Y

		// A = (bx*cy - by*cx) / (bx*ay - by*ax)
		// B = (cx - A*ax)/(bx)
		btnA := (bx*cy - by*cx) / (bx*ay - by*ax)
		btnB := (cx - btnA*ax) / bx

		if (bx*cy-by*cx)%(bx*ay-by*ax) != 0 || (cx-btnA*ax)%bx != 0 {
			continue
		}

		cost += btnA*3 + btnB*1
	}
	return cost
}

func second(arcades []arcade) int {
	var cost int
	for _, a := range arcades {
		a.prize.X += 10000000000000
		a.prize.Y += 10000000000000
		// A*ax + B*bx = cx
		// A*ay + B*by = cy
		ax, ay := a.a.X, a.a.Y
		bx, by := a.b.X, a.b.Y
		cx, cy := a.prize.X, a.prize.Y

		// A = (bx*cy - by*cx) / (bx*ay - by*ax)
		// B = (cx - A*ax)/(bx)
		btnA := (bx*cy - by*cx) / (bx*ay - by*ax)
		btnB := (cx - btnA*ax) / bx

		if (bx*cy-by*cx)%(bx*ay-by*ax) != 0 || (cx-btnA*ax)%bx != 0 {
			continue
		}

		cost += btnA*3 + btnB*1
	}
	return cost
}

type arcade struct {
	a     maps.Coordinate
	b     maps.Coordinate
	prize maps.Coordinate
}

func (a arcade) String() string {
	return fmt.Sprintf("A: %s, B: %s, Prize: %s", a.a, a.b, a.prize)
}
