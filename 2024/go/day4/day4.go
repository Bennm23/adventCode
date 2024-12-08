package main

import (
	"advent/lib"
)

func main() {
    lib.RunAndScore("Part 1", p1)//2547 Time = 1119 us
    lib.RunAndScore("Part 2", p2)//1939 Time = 654 us
}

type Location struct {
	row int
	col int
}
func (location Location) inBounds(size int) bool {
	return location.row >= 0 && location.row < size && location.col >= 0 && location.col < size
}

func p1() int {
    sum := 0

	grid := lib.ReadFileToGrid("day4.txt")

	xs := findRunes(&grid, 'X', false)

	for _, x := range xs {
		//Up
		if evalDirection(&grid, x, func(loc Location) Location { return Location{loc.row - 1, loc.col}}) {
			sum += 1
		}
		//Down
		if evalDirection(&grid, x, func(loc Location) Location { return Location{loc.row + 1, loc.col}}) {
			sum += 1
		}
		//Left
		if evalDirection(&grid, x, func(loc Location) Location { return Location{loc.row, loc.col - 1}}) {
			sum += 1
		}
		//Right
		if evalDirection(&grid, x, func(loc Location) Location { return Location{loc.row, loc.col + 1}}) {
			sum += 1
		}
		//Up Right
		if evalDirection(&grid, x, func(loc Location) Location { return Location{loc.row - 1, loc.col + 1}}) {
			sum += 1
		}
		//Up Left
		if evalDirection(&grid, x, func(loc Location) Location { return Location{loc.row - 1, loc.col - 1}}) {
			sum += 1
		}
		//Down Right
		if evalDirection(&grid, x, func(loc Location) Location { return Location{loc.row + 1, loc.col + 1}}) {
			sum += 1
		}
		//Down Left
		if evalDirection(&grid, x, func(loc Location) Location { return Location{loc.row + 1, loc.col - 1}}) {
			sum += 1
		}
	}
    return sum
}

func p2() int {
    sum := 0

	grid := lib.ReadFileToGrid("day4.txt")

	as := findRunes(&grid, 'A', true)

	valids := [4][4]rune {
		//TL, TR, BL, BR
		{'M', 'M', 'S','S'},
		{'M', 'S', 'M','S'},
		{'S', 'M', 'S','M'},
		{'S', 'S', 'M','M'},
	}


	for _, a := range as {

		curr := [4]rune {
			//TL TR BL BR
			grid[a.row - 1][a.col - 1],
			grid[a.row - 1][a.col + 1],
			grid[a.row + 1][a.col - 1],
			grid[a.row + 1][a.col + 1],
		}

		for _, v := range valids {
			if v == curr {
				sum += 1
				break
			}
		}
	}

    return sum
}

func evalDirection(grid *[][]rune, start Location, move func(Location) Location) bool {
	expected := 'M'

	curr := start

	for {
		curr = move(curr)

		if !curr.inBounds(len(*grid)) {
			return false;
		}

		if expected == 'S' && (*grid)[curr.row][curr.col] == 'S' {
			return true;
		} else if expected == 'A' && (*grid)[curr.row][curr.col] == 'A' {
			expected = 'S'
		} else if expected == 'M' && (*grid)[curr.row][curr.col] == 'M' {
			expected = 'A'
		} else {
			return false
		}
	}
}
func findRunes(grid *[][]rune, r rune, ignoreEdge bool) []Location {
	var matches []Location

	for rix, row := range *grid {

		if ignoreEdge && (rix == 0 || rix == len(*grid) - 1) {
			continue
		}

		for cix, col := range row {
			if ignoreEdge && (cix == 0 || cix == len(*grid) - 1) {
				continue
			}
			if col == r {
				matches = append(matches, Location{rix, cix})
			}
		}
	}
	return matches
}