package main

import (
	"fmt"
	"advent/lib"
)

type Position struct {
	x int;
	y int;
	val rune;
}
type Direction int;
const (
	UP Direction = 0;
	DOWN Direction = 1;
	LEFT Direction = 2;
	RIGHT Direction = 3;
)
func main() {

	grid := lib.ReadFileToGrid("day10.txt")
	
	var start Position

	for i, row := range grid {

		for j, col := range row {
			if col == 'S' {
				start = Position{
					i, j, col,
				}
			}
		}
	}

	//Find one valid path from start
	//At each node, find the two valid routes to go
	//If start, pick any
	//If not, pick the one we did not just come from
	current := start;

	var path []Position

	path = append(path, current)

	s, direction := findAStart(grid, start)

	current = s

	for current.val != 'S' {

		path = append(path, current)

		if direction == LEFT {
			if current.val == '-' {
				direction = LEFT
			} else if current.val == 'L' {
				direction = UP
			} else if current.val == 'F' {
				direction = DOWN
			} else {
				panic("FAILED TO MOVE LEFT")
			}
		} else if direction == RIGHT {
			if current.val == '-' {
				direction = RIGHT
			} else if current.val == 'J' {
				direction = UP
			} else if current.val == '7' {
				direction = DOWN
			} else {
				panic("FAILED TO MOVE RIGHT")
			}
		} else if direction == UP {
			if current.val == '|' {
				direction = UP
			} else if current.val == 'F' {
				direction = RIGHT
			} else if current.val == '7' {
				direction = LEFT
			} else {
				panic("FAILED TO MOVE UP")
			}
		} else if direction == DOWN {
			if current.val == '|' {
				direction = DOWN
			} else if current.val == 'L' {
				direction = RIGHT
			} else if current.val == 'J' {
				direction = LEFT
			} else {
				panic("FAILED TO MOVE DOWN")
			}
		} else {
			panic("Direction is invalid")
		}

		current = buildPosition(grid, current, direction)
	}


	for _, p := range path {
		fmt.Println(p)
	}

	fmt.Println("Part 1 = ", len(path)/2)


	//https://en.wikipedia.org/wiki/Shoelace_formula
	doubleArea := 0 //2A = SUM n | path[n].y  path[n+1].y |
					//			 | path[n].x  path[n+1].x |

	for n := 0; n < len(path) - 1; n++ {
		doubleArea += (path[n].x + path[n+1].x) * (path[n].y - path[n+1].y)
	}

	if doubleArea < 0 {
		doubleArea *= -1
	}

	//A = interior + boundary/2 - 1
	//A - boundary/2 + 1 = interior
	interiorPoints := doubleArea/2 - len(path) / 2 + 1
	fmt.Println("Part 2 = ", interiorPoints)//https://en.wikipedia.org/wiki/Pick%27s_theorem
}

func buildPosition(grid [][]rune, current Position, dir Direction) Position {
	x := current.x
	y := current.y
	if dir == LEFT {
		return Position {x, y - 1, grid[x][y - 1]}
	}
	if dir == RIGHT {
		return Position {x, y + 1, grid[x][y + 1]}
	}
	if dir == UP {
		return Position {x - 1, y, grid[x - 1][y]}
	}
	if dir == DOWN {
		return Position {x + 1, y, grid[x + 1][y]}
	}

	panic("Failed to build position")
}

func findAStart(grid [][]rune, start Position) (Position, Direction) {
	val, found := checkLeft(grid, start)
	if found {
		return Position { start.x, start.y - 1, val }, LEFT
	}
	val, found = checkRight(grid, start)
	if found {
		return Position { start.x, start.y + 1, val }, RIGHT
	}
	val, found = checkTop(grid, start)
	if found {
		return Position { start.x - 1, start.y, val }, UP
	}
	val, found = checkBottom(grid, start)
	if found {
		return Position { start.x + 1, start.y, val }, DOWN
	}

	panic("Failed to find a start")
}

func checkLeft(grid [][]rune, current Position) (rune, bool) {
	row := grid[current.x]

	if current.y == 0 {
		return '.', false
	}

	left := row[current.y - 1]

	if left == 'L' || left == 'F' || left == '-' {
		return left, true
	}

	return '.', false
}

func checkRight(grid [][]rune, current Position) (rune, bool) {
	row := grid[current.x]

	if current.y == len(row) - 1 {
		return '.', false
	}

	right := row[current.y + 1]

	if right == 'J' || right == '7' || right == '-' {
		return right, true
	}

	return '.', false
}

func checkTop(grid [][]rune, current Position) (rune, bool) {

	if current.x == 0 {
		return '.', false
	}

	top := grid[current.x - 1][current.y]

	if top == 'F' || top == '7' || top == '|' {
		return top, true
	}

	return '.', false
}

func checkBottom(grid [][]rune, current Position) (rune, bool) {

	if current.x == len(grid) - 1 {
		return '.', false
	}

	bottom := grid[current.x + 1][current.y]

	if bottom == 'J' || bottom == 'L' || bottom == '|' {
		return bottom, true
	}

	return '.', false
}