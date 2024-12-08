package main

import (
	"advent/lib"
	"advent/lib/structures"
	"sort"
	"strconv"
	"strings"
)

func main() {

	lib.RunAndScore("Part 1", p1)//Score = 1579939.  Time 284 us
	lib.RunAndScore("Part 2", p2)//Score = 20351745. Time 603 us
}

type Pair struct {
	Left int
	Right int
}

func getLeftAndRightList() ([]int, []int) {
	pairs := lib.ReadFileToTypeVec("day1.txt", func(s string) Pair {

		s = strings.ReplaceAll(s, "   ", ",")
		split := strings.Split(s, ",")

		left, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		return Pair{left, right}
	})

	var left []int
	var right []int

	for _, pair := range pairs {
		left = append(left, pair.Left)
		right = append(right, pair.Right)
	}
	return left, right
}

func p1() int {

	left, right := getLeftAndRightList()

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i, l := range left {
		sum += lib.Absi(l - right[i])
	}
	return sum
}

func p2() int {

	left, right := getLeftAndRightList()
	sum := 0
	for _, l := range left {
		sum += l * structures.CountMatches(right, l)
	}
	return sum
}