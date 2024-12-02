package main

import (
	"advent/lib"
	"fmt"
	"math"
)

type Direction [2]int
type HeatMap [][]int

type State struct {
	x, y int
	direction Direction;
	moveCount, score int;
}

func NewState(x, y, moves int, dir Direction) State{
	return State {
		x, y,
		dir,
		moves,
		0,
	}
}

var MOVES = []Direction{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var GRID_LEN = 0
var GRID_WIDTH = 0

func main() {
	fmt.Println("Day 17")

	var grid HeatMap = lib.ReadFileToTypeGrid("sample", func(r rune) int {
		return int(r - '0')
	})
	GRID_LEN = len(grid)
	GRID_WIDTH = len(grid[0])


	fmt.Println("Part 1 = ", solve(grid, 3, 1))
}

func getNeighbors(curr State, maxMoves, minMoves int) []State{
	neighbors := make([]State,0)

	for _, move := range MOVES {
		movesSoFar := 0
		x, y := curr.x + move[0], curr.y + move[1]

		if move == curr.direction {
			movesSoFar = curr.moveCount
		} else {
			movesSoFar = 1
		}
	
		if (movesSoFar > maxMoves) {
			continue
		}
		if (move != curr.direction && curr.moveCount < minMoves) {
			continue
		}
		if x < 0 || y < 0 || x >= GRID_LEN || y >= GRID_WIDTH {
			continue
		}
		if move[0] * -1 == curr.direction[0] && move[1] * -1 == curr.direction[1] {
			continue
		}

		neighbors = append(neighbors, NewState(x, y, movesSoFar, move))
	}
	return neighbors
}

func solve(grid HeatMap, maxMoves, minMoves int) int {
	openSet := make([]State, 0)
	best := make(map[State]int)

	for _, move := range MOVES {
		openSet = append(openSet, NewState(0, 0, 0, move))
	}

	for len(openSet) > 0 {
		curr := openSet[0]
		fmt.Println("Looping on ", curr)
		openSet = openSet[1:]

		if _,found := best[curr]; found {
			continue
		}

		best[curr] = curr.score

		for _, neighbor := range getNeighbors(curr, maxMoves, minMoves) {

			if _, found := best[neighbor]; !found {
				weight := best[curr] + grid[neighbor.x][neighbor.y]
				neighbor.score = weight
				openSet = append(openSet, neighbor)
			}

		}

	}

	min := math.MaxInt

	for _, val := range best {
		if val < min {
			min = val
		}
	}

	return min

}
