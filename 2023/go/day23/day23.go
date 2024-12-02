package main

import (
	"advent/lib"
	"advent/lib/structures"
	"fmt"
)

type PointGraph map[Position]map[Position]int
type Position [2]int
type INode struct {
	position Position;
	weight int;
}

var MOVES [][2]int = [][2]int{ {0, 1}, {0, -1}, {1, 0}, {-1, 0} }

var GRID_LEN int = 0
var GRID_WIDTH int = 0

var endPos Position
var pointGraph = PointGraph{}

func getValidMoves(standingOn rune) [][2]int {
	return MOVES//Part 2

	// if standingOn == 'v' {
	// 	return [][2]int{ {1, 0} }
	// }
	// if standingOn == '^' {
	// 	return [][2]int{ {-1, 0} }
	// }
	// if standingOn == '>' {
	// 	return [][2]int{ {0, 1} }
	// }
	// if standingOn == '<' {
	// 	return [][2]int{ {0, -1} }
	// }
	// if standingOn == '.' {
	// 	return MOVES
	// }
	// panic("Got To Standing on Forest")
}


func solve() {
	grid := lib.ReadFileToGrid("day23.txt")

	GRID_LEN = len(grid)
	GRID_WIDTH = len(grid[0])

	startPos := Position {}
	forest := structures.Set[Position]{}

	for r, row := range grid {
		for c, col := range row {
			if r == 0 && col == '.' {
				startPos = Position { r, c }
			}
			if r == GRID_LEN - 1 && col == '.' {
				endPos = Position { r, c }
			}
			if col == '#' {
				forest.Add(Position { r, c })
			}
		}
	}


	//shorten edges. For each node, if it touches more then two valid points, add it
	important := structures.Set[Position]{startPos, endPos}

	for i, row := range grid {
		for j, col := range row {
			if col == '#' {
				continue
			}

			decisions := 0
			for _, move := range MOVES {

				xx, yy := i + move[0], j + move[1]

				if xx < 0 || xx >= GRID_LEN || yy < 0 || yy >= GRID_WIDTH {
					continue
				}
				if grid[xx][yy] != '#' {
					decisions++
				}
			}
			if decisions > 2 {
				important.Add(Position{i, j})
			}
		}
	}

	for _, point := range important {
		pointGraph[point] = map[Position]int{}

		explore := structures.Stack[INode]{}
		explore.Push(INode{point, 0})
		seen := structures.Set[Position]{ point }

		for explore.Size() != 0 {
			curr := explore.Pop()

			if curr.weight != 0 && important.Contains(curr.position) {
				pointGraph[point][curr.position] = curr.weight
				continue
			}

			for _, move := range getValidMoves(grid[curr.position[0]][curr.position[1]]) {
			
				xx, yy := curr.position[0] + move[0], curr.position[1] + move[1]
				if xx < 0 || xx >= GRID_LEN || yy < 0 || yy >= GRID_WIDTH || grid[xx][yy] == '#' || seen.Contains(Position{xx, yy}) {
					continue
				}
				seen.Add( Position{xx, yy} )
				explore.Push(INode{ Position{xx, yy}, curr.weight + 1})
			}
		}
	}


	longest := search(startPos)

	fmt.Println("Longest = ", longest)//P1: 2246, P2: 6622

}


var searchSeen = structures.Set[Position]{}

func search(position Position) int {
	total := -1

	if position == endPos {
		return 0
	}

	searchSeen.Add(position)
	for child, weight := range pointGraph[position] {
		if searchSeen.Contains(child) {
			continue
		}
		s := weight

		s += search(child)

		if s > total {
			total = s
		}
	}
	searchSeen.Remove(position)

	return total
}

func main() {

	lib.RunAndPrintDurationMillis(func () {
		solve()
	})
	//No Global 7137, 6893 MS
	//Global 6959 6998, 7057, 6911
}