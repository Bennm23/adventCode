package main

import (
	"advent/lib"
	"advent/lib/structures"
	"fmt"
)

func main() {
	solve()
}

// type Position [2]int
type Coord [2]int
type Position struct {
	coords   Coord
	useCount int
}

var GRID_LEN int = 0
var GRID_WIDTH int = 0
var MOVES [][]int = [][]int{ {0, 1}, {0, -1}, {1, 0}, {-1, 0} }
var ROCKS structures.Set[Coord] = structures.Set[Coord]{}
var START Key

func solve() {

	grid := lib.ReadFileToGrid("sample")

	GRID_LEN = len(grid)
	GRID_WIDTH = len(grid[0])

	rockCount := 0

	for i, row := range grid {
		fmt.Println(string(row))
		for j, col := range row {
			if col == 'S' {
				// START = Position{Coord{i, j}, 1}
				START = Key{Coord{i, j}, 0}
			}
			if col == '#' {
				ROCKS.Add( Coord{ i, j } )
				rockCount++
			}
		}
	}


	fmt.Println("Part 1 = 3677") //3677

	fmt.Println("Part 1 = ", doLoop(100))
}

type Key struct {
	coord Coord;
	step int;
}

func doLoop(steps int) int64 {

	visited := make(map[Coord]int)
	explore := structures.Stack[Key]{}
	explore.Push( START )

	for explore.Size() > 0 {
		e := explore.Pop()

		if _, found := visited[e.coord]; found {
			continue
		}
		if e.step == steps + 1 {
			break
		}
		visited[e.coord] = e.step

		for _, mv := range MOVES {
			moveTo := Coord { e.coord[0] + mv[0], e.coord[1] + mv[1] }
			
			if moveTo[0] < 0 || moveTo[0] >= GRID_LEN || moveTo[1] < 0 || moveTo[1] >= GRID_WIDTH {
				continue
			}
			if ROCKS.Contains(moveTo) {
				continue
			}

			explore.Push( Key{ moveTo, e.step + 1 } )
		}

	}

	sum := int64(0)

	fmt.Println("Visited Len = ", len(visited))
	for _, val := range visited {
		if val % 2 == 0 {
			sum++
		}
	}

	return sum
}