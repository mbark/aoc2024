package day22

import (
	"fmt"

	"github.com/mbark/aoc2024/fns"
	"github.com/mbark/aoc2024/util"
)

const testInput = `
1
2
3
2024
`

func Run(input string, isTest bool) {
	if isTest {
		input = testInput
	}

	in := util.Str2IntSlice(util.ReadInput(input, "\n"))
	fmt.Println("first:", first(in))
	fmt.Println("second:", second(in))
}

func first(in []int) int {
	var secrets []int
	for _, secret := range in {
		for i := 0; i < 2000; i++ {
			n := secret * 64
			secret = mix(secret, n)
			secret = prune(secret)

			n = secret / 32
			secret = mix(secret, n)
			secret = prune(secret)

			n = secret * 2048
			secret = mix(secret, n)
			secret = prune(secret)
		}

		secrets = append(secrets, secret)
	}

	var sum int
	for _, s := range secrets {
		sum += s
	}
	return sum
}

func second(in []int) int {
	buyers := generateSecrets(in)
	diffToPrice := map[diff][]int{}
	for _, secrets := range buyers {
		prices := getPrices(secrets)
		diffs := priceChanges(prices)

		buyerPrices := map[diff]int{}

		for i := 0; i < len(diffs)-4; i++ {
			d := diff{diffs[i], diffs[i+1], diffs[i+2], diffs[i+3]}
			if _, ok := buyerPrices[d]; ok {
				continue
			}

			buyerPrices[d] = prices[i+4]
		}

		for d, p := range buyerPrices {
			diffToPrice[d] = append(diffToPrice[d], p)
		}
	}

	best := 0
	for _, p := range diffToPrice {
		if best < fns.Sum(p) {
			best = fns.Sum(p)
		}
	}

	return best
}

type diff [4]int

func (d diff) String() string {
	return fmt.Sprintf("%d,%d,%d,%d", d[0], d[1], d[2], d[3])
}

func getPrices(in []int) []int {
	var prices []int
	for _, s := range in {
		prices = append(prices, s%10)
	}
	return prices
}

func priceChanges(in []int) []int {
	var diff []int
	for i := 0; i < len(in)-1; i++ {
		diff = append(diff, in[i+1]-in[i])
	}
	return diff
}

func generateSecrets(in []int) map[int][]int {
	m := make(map[int][]int)
	for i, secret := range in {
		secrets := []int{secret}
		for i := 0; i < 2000; i++ {
			n := secret * 64
			secret = mix(secret, n)
			secret = prune(secret)

			n = secret / 32
			secret = mix(secret, n)
			secret = prune(secret)

			n = secret * 2048
			secret = mix(secret, n)
			secret = prune(secret)

			secrets = append(secrets, secret)
		}

		m[i] = secrets
	}
	return m
}

func mix(s, n int) int {
	return s ^ n
}

func prune(s int) int {
	return s % 16777216
}
