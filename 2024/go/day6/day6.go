package main

import (
	"advent/lib"
	"advent/lib/maths"
	"advent/lib/structures"
)

type Guard struct {
	position  maths.Position
	direction maths.Position
}

var VALID_MOVES = []maths.Position{
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
}

func (guard *Guard) move(grid *[][]rune) bool {

	new_pos := maths.Position{
		X: guard.position.X + guard.direction.X,
		Y: guard.position.Y + guard.direction.Y,
	}
	if !new_pos.InBounds(len(*grid)) {
		guard.position = new_pos
		return false
	}

	if (*grid)[new_pos.X][new_pos.Y] == '#' {
		index := structures.IndexOf(VALID_MOVES, guard.direction)
		index = (index + 1) % 4
		guard.direction = VALID_MOVES[index]
	} else {
		guard.position = new_pos
	}

	return true
}

func main() {
	lib.RunAndScore("Part 1", p1)
	lib.RunAndScore("Part 2", p2)
}

func getPath(grid [][]rune) structures.Set[maths.Position] {

	var guard Guard

	for rix, row := range grid {
		for cix, col := range row {
			if col == '^' {
				guard = Guard{maths.Position{X: rix, Y: cix}, maths.Position{X: -1, Y: 0}}
			}
		}
	}

	visited := make(structures.Set[maths.Position], 0)

	for guard.move(&grid) {
		visited.Add(guard.position)
	}

	return visited
}

func p1() int {

	grid := lib.ReadFileToGrid("day6.txt")
	path := getPath(grid)
	return len(path)
}
func p2() int {
	sum := 0
	grid := lib.ReadFileToGrid("day6.txt")
	path := getPath(grid)

	return sum
}
