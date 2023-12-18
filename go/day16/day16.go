package main

import (
	"advent/lib"
	"fmt"
	"math"
)

type Direction uint8

const (
	NORTH Direction = 0
	SOUTH Direction = 1
	EAST  Direction = 2
	WEST  Direction = 3
)

type Light struct {
	direction Direction
	position  []int
}

func NewLight(d Direction, pos []int) Light {
	return Light{
		d, pos,
	}
}

var rowLen int
var colLen int

type Visited map[string]struct{}

func addVisited(row, col int, visited *Visited) {
	str := fmt.Sprintf("%d,%d", row, col)
	if _, found := (*visited)[str]; !found {
		(*visited)[str] = struct{}{}
	}
}

func movePos(position *[]int, dir Direction) bool {
	row := (*position)[0]
	col := (*position)[1]

	switch dir {

	case NORTH:
		if row-1 < 0 {
			return false
		}
		(*position)[0] = row - 1

	case SOUTH:
		if row+1 >= rowLen {
			return false
		}
		(*position)[0] = row + 1

	case EAST:
		if col+1 >= colLen {
			return false
		}
		(*position)[1] = col + 1

	case WEST:
		if col-1 < 0 {
			return false
		}
		(*position)[1] = col - 1
	default:
		panic("Unrecognized Direction in move")
	}

	return true
}

func alreadyVisited(started *[][]int, light *Light) bool {
	if len(*started) == 0 {
		return false
	}
	for _, start := range *started {
		if start[0] == int(light.direction) && start[1] == light.position[0] && start[2] == light.position[1] {
			return true
		}
	}
	return false
}

func solveFromStart(grid [][]rune, firstLight Light) int {
	visited := make(Visited)
	var lights []Light
	lights = append(lights, firstLight)

	started := make([][]int, 0)

	for len(lights) > 0 {

		//Pop off the current light
		light := lights[0]
		lights = lights[1:]

		if alreadyVisited(&started, &light) {
			continue
		}

		started = append(started, []int{int(light.direction), light.position[0], light.position[1]})

		for {

			moveGood := movePos(&light.position, light.direction)
			//If the move is bad, then we have just moved ob, this lights run is over
			if !moveGood {
				break
			}
			row, col := light.position[0], light.position[1]
			char := grid[row][col]
			addVisited(row, col, &visited)

			if char == '.' {
				//If this is an empty space, keep moving
				continue
			} else if char == '\\' {
				//If this is a right slash, we can go either down or up
				if light.direction == EAST {
					light.direction = SOUTH
				} else if light.direction == SOUTH {
					light.direction = EAST
				} else if light.direction == WEST {
					light.direction = NORTH
				} else if light.direction == NORTH {
					light.direction = WEST
				}
				continue
			} else if char == '/' {
				if light.direction == EAST {
					light.direction = NORTH
				} else if light.direction == SOUTH {
					light.direction = WEST
				} else if light.direction == WEST {
					light.direction = SOUTH
				} else if light.direction == NORTH {
					light.direction = EAST
				}
				continue
			} else if char == '|' {
				//Encountered a vertical splitter
				//If moving W/E we must split
				if light.direction == EAST || light.direction == WEST {

					lights = append(lights, NewLight(NORTH, []int{light.position[0], light.position[1]}))
					lights = append(lights, NewLight(SOUTH, []int{light.position[0], light.position[1]}))
					//This light is no longer active now that it has split
					break
				} else {
					//Else move through in same direction
					continue
				}
			} else if char == '-' {
				//Encountered a horizontal splitter
				//If moving N/S we must split
				if light.direction == NORTH || light.direction == SOUTH {

					lights = append(lights, NewLight(WEST, []int{light.position[0], light.position[1]}))
					lights = append(lights, NewLight(EAST, []int{light.position[0], light.position[1]}))
					//This light is no longer active now that it has split
					break
				} else {
					//Else move through in same direction
					continue
				}

			} else {
				panic(fmt.Sprintf("Unrecognized Input %s", string(char)))
			}
		}
	}

	return len(visited)
}
func solve() {
	grid := lib.ReadFileToGrid("day16.txt")

	rowLen = len(grid)
	colLen = len(grid[0])

	firstLight := NewLight(EAST, []int{0, -1})

	p1 := solveFromStart(grid, firstLight)

	p2 := math.MinInt

	//[0][0:colLen] top row
	for col := 0; col < colLen; col++ {
		solves := []Light{}
		//If in top left
		if col == 0 {
			solves = append(solves, NewLight(EAST, []int{0, -1}))
			solves = append(solves, NewLight(SOUTH, []int{-1, 0}))
		} else if col == colLen - 1{
			//If in top right
			solves = append(solves, NewLight(WEST, []int{0, colLen}))
			solves = append(solves, NewLight(SOUTH, []int{-1, colLen - 1}))
		} else {
			//Else any other column on top row
			solves = append(solves, NewLight(SOUTH, []int{-1, col}))
		}

		for _, solve := range solves {
			val := solveFromStart(grid, solve)
			if val > p2 {
				p2 = val
			}
		}
	}
	//[rowlen-1][0:colLen] bottom row
	for col := 0; col < colLen; col++ {
		solves := []Light{}
		//If in bottom left
		if col == 0 {
			solves = append(solves, NewLight(EAST, []int{rowLen - 1, -1}))
			solves = append(solves, NewLight(NORTH, []int{rowLen, 0}))
		} else if col == colLen - 1{
			//If in bottom right
			solves = append(solves, NewLight(WEST, []int{rowLen - 1, colLen}))
			solves = append(solves, NewLight(NORTH, []int{rowLen, colLen - 1}))
		} else {
			//Else any other column on bottom row
			solves = append(solves, NewLight(NORTH, []int{rowLen, col}))
		}

		for _, solve := range solves {
			val := solveFromStart(grid, solve)
			if val > p2 {
				p2 = val
			}
		}
	}
	//Corners have been taken care of
	//[0:rowlen][0] left col
	for row := 1; row < rowLen - 1; row++ {
		start := NewLight(EAST, []int{row, -1})
		val := solveFromStart(grid, start)
		if val > p2 {
			p2 = val
		}

		//[0:rowlen][collen] right col
		start = NewLight(WEST, []int{row, colLen})
		val = solveFromStart(grid, start)
		if val > p2 {
			p2 = val
		}
	}

	fmt.Println("Part 1 = ", p1)//7185
	fmt.Println("Part 2 = ", p2)//7616

}
func main() {

	lib.RunAndPrintDuration( func() {
		solve()
	})//750_000 - 900_000

}