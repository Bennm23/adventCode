package main

import (
	"advent/lib"
	"fmt"
	"strconv"
	"strings"
)

type Direction [2]int
type Terrain [][]byte
type Coord [2]int


var MOVES = []Direction{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func BuildDirection(dirString string) Direction {
	if dirString == "R" || dirString == "0" {
		return MOVES[0]
	} else if dirString == "L"  || dirString == "2"{
		return MOVES[1]
	} else if dirString == "D"  || dirString == "1"{
		return MOVES[2]
	} else if dirString == "U"  || dirString == "3"{
		return MOVES[3]
	} else {
		panic(fmt.Sprintf("Invalid Direction = %s", dirString))
	}
}

func solve(lines []string, parser func(string) (Direction, int64)) int64{
	path := make([]Coord, 0)

	curr := Coord{0, 0}

	for _, line := range lines {

		dir, length := parser(line)
		// rgb := split[2][2:len(split[2]) - 1]


		var i int64 = 0
		for i = 0; i < length; i++ {
			curr = move(curr, dir)
			path = append(path, curr)
		}

	}

	//https://en.wikipedia.org/wiki/Shoelace_formula
	var doubleArea int64 = 0 //2A = SUM n | path[n].y  path[n+1].y |
					//			 | path[n].x  path[n+1].x |

	for n := 0; n < len(path) - 1; n++ {
		doubleArea += int64((path[n][0] + path[n+1][0]) * (path[n][1] - path[n + 1][1]))
	}
	if doubleArea < 0 {
		doubleArea *= -1
	}
	interiorPoints := doubleArea/2 - int64(len(path)) / 2 + 1

	return int64(len(path)) + interiorPoints
}

func main() {
	lib.RunAndPrintDurationMillis(func() {
		ans()
	})//4447, 4334, 4040
}

func ans() {
	lines := lib.ReadFile("day18.txt")
	

	p1 := solve(lines, func(line string) (Direction, int64) {
		split := strings.Split(line, " ")
		dir := BuildDirection(split[0])
		length, _ := strconv.Atoi(split[1])

		return dir, int64(length)
	})

	p2 := solve(lines, func(line string) (Direction, int64) {
		split := strings.Split(line, " ")

		hex := split[2][2:len(split[2]) - 1]
		dir := BuildDirection(hex[len(hex) - 1:])
		length, _ := strconv.ParseInt(hex[:len(hex)-1], 16, 64);


		return dir, length
	})

	fmt.Println("Part 1 = ", p1)//50746
	fmt.Println("Part 2 = ", p2)//70086216556038

}

func move(coord Coord, direction Direction) Coord {
	return Coord {
		coord[0] + direction[0],
		coord[1] + direction[1],
	}
}