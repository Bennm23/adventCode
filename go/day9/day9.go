package main

import (
	"advent/lib"
	"advent/lib/avstrings"
	"fmt"
)

func main() {
	histories, err := lib.ReadFile("day9.txt")

	if err != nil {
		panic("Failed to parse Day 9")
	}

	p1, p2 := solve(histories)
	fmt.Println("Part 1 = ", p1)//2105961943
	fmt.Println("Part 2 = ", p2)//1019
	
}

func solve(histories []string) (int, int) {

	total := 0
	total2 := 0
	for _, history := range histories {
		//For each history
		//	1) Walk down from root until current arr is all zeros

		historyInts := avstrings.SplitTextToInts(history)
		var steps [][]int
		steps = append(steps, historyInts)


		done := false;

		var res []int
		res = historyInts

		for !done {
			res = stepDown(res)
			steps = append(steps, res)
		

			done = true
			for _, r := range res {
				if r != 0 {
					done = false
					break
				}
			}
		}

		total += walkUp(steps)
		total2 += walkUp2(steps);
	}

	return total, total2
}

func stepDown(ints []int) []int {
	var res []int

	for i := 0; i < len(ints) - 1; i++ {
		res = append(res, ints[i+1] - ints[i])
	}

	return res
}

func walkUp(steps [][]int) int {
	//Starting at one above the bottom
	// 1) Append to end of line line[len - 1] + line+1[len - 2]

	for i := len(steps) - 2; i > -1; i-- {
		steps[i] = append(steps[i], steps[i][len(steps[i]) - 1] + steps[i + 1][len(steps[i+1]) - 1])
	}

	return steps[0][len(steps[0]) - 1]
}

func walkUp2(steps [][]int) int {
	//Starting at one above the bottom
	// 1) Append to front of line (steps[i][0] - steps[i+1][0])

	for i := len(steps) - 2; i > -1; i-- {
		steps[i] = append([]int{steps[i][0] - steps[i+1][0]}, steps[i]...)
	}

	return steps[0][0]
}