package main

import (
	"advent/lib"
	"advent/lib/maths"
)

func main() {
	lib.RunAndScore("Part 1", p1)//Score = 180233229, Laptop Time 2490 us
	lib.RunAndScore("Part 2", p2)//Score =  95411583. Laptop Time 2559 us
}

func buildSuperline(doEnabled bool) string {
	lines := lib.ReadFile("day3.txt")

	res := ""

	for _, line := range lines {
		res += line
	}
	if doEnabled {
		res = lib.RemoveStrBetweenOrAfter(res, `don't\(\)`, `do\(\)`)
	}
	return res
}

func countMulTotal(superline string) int {

	sum := 0
	muls := lib.FindAllMatches(`mul\(\d{1,3},\d{1,3}\)`, superline)

	for _, mul := range muls {
		sum += lib.EvaluateMatch(
			`\d+`,
			mul,
			func(s []string) int {
				left := maths.ToInt(s[0])
				right := maths.ToInt(s[1])

				return left * right
			},
		)
	}
	return sum;
}

func p1() int {
	superline := buildSuperline(false)
	return countMulTotal(superline)
}
func p2() int {
	superline := buildSuperline(true)
	return countMulTotal(superline)
}
