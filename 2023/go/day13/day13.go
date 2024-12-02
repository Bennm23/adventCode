package main

import (
	"advent/lib"
	"advent/lib/maths"
	"fmt"
)

func main() {

	lib.RunAndPrintDuration(func() {
		groups := lib.ReadFileToGroups("day13.txt", "")

		var p1, p2 int64 = 0, 0
		for i,group := range groups {
			fmt.Println("Group ", i)
			grid := buildGrid(group)

			ver1, ver2 := checkReflect(grid);

			transposed := maths.Transpose[rune](grid)
			hor, hor2 := checkReflect(transposed);

			p1 += ver1*100 + hor
			p2 += ver2*100 + hor2
		}

		fmt.Println("Part 1 = ", p1)//37975
		fmt.Println("Part 2 = ", p2)//32497
	})//780, 734, 782, 729
}

func buildGrid(group []string) [][]rune {
	var grid [][]rune

	for _, row := range group {
		temp := make([]rune, 0)
		for _, c := range row {
			temp = append(temp, c)
		}
		grid = append(grid, temp)
	}

	return grid
}

func checkReflect(group [][]rune) (int64, int64) {
	var p1, p2 int64


	//For each row 
	for row := 1; row < len(group); row++ {
		reflectRange := min(row, len(group) - row)//The range to search for reflection
		upRow, downRow := row - 1, row

		matchCount := 0

		//For each possible row in reflection
		//Compare chars in each row, increment count if found
		for dRow := 0; dRow < reflectRange; dRow++ {
			upLine, downLine := group[upRow - dRow], group[downRow + dRow]

			for index := 0; index < len(upLine); index++ {
				if downLine[index] == upLine[index] {
					matchCount++
				}
			}
		}

		//Maximum possible char matches is number of rows compared times length of rows
		perfectCount := reflectRange * len(group[row])

		if matchCount == perfectCount {
			p1 = int64(row)
		} else if matchCount == perfectCount - 1 {
			p2 = int64(row)
		}

	}


	return p1, p2
}