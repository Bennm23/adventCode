package main

import (
	"advent/lib"
	"advent/lib/maths"
	"advent/lib/structures"
	"fmt"
)

type Robot struct {
    position maths.Position
    velocity maths.Position
}

const FLOOR_WIDTH int = 101
const FLOOR_HEIGHT int = 103

func NewRobot(vals []int) Robot {
    return Robot {
        maths.NewPosition(vals[0], vals[1]),
        maths.NewPosition(vals[2], vals[3]),
    }
}
func (r Robot) Print() {
    fmt.Println("Position = ", r.position, ", Velocity = ", r.velocity)
}

func main() {
    lib.RunAndScore("Part 1", p1)//Result = 217328832 : Total Time    3509 us
    lib.RunAndScore("Part 2", p2)//Result = 7412      : Total Time 2436805 us
}

func p1() int {
    lines := lib.ReadFile("day14.txt")

    robots := make([]Robot, 0)

    for _, line := range lines {
        robots = append(robots, NewRobot(lib.ParseIntsFromString(line)))
    }
    
    for range 100 {

        for i := range robots {

            robots[i].position.PAdd(robots[i].velocity)

            robots[i].position.X %= FLOOR_WIDTH
            if robots[i].position.X < 0 {
                robots[i].position.X += FLOOR_WIDTH
            }
            robots[i].position.Y %= FLOOR_HEIGHT
            if robots[i].position.Y < 0 {
                robots[i].position.Y += FLOOR_HEIGHT
            }
        }
    }

    q1, q2, q3, q4 := 0, 0, 0, 0


    middleCol := FLOOR_WIDTH / 2
    middleRow := FLOOR_HEIGHT / 2

    robotGrid := getRobotGrid(robots)

    for rix, row := range robotGrid {
        if rix == middleRow {
            continue
        }
        for cix, col := range row {
            if col == 0 {
                continue
            }
            if cix == middleCol {
                continue
            }
            left := cix < middleCol
            up := rix < middleRow


            if left && up {
                q1 += col
            } else if !left && up {
                q2 += col
            } else if !left && !up {
                q3 += col
            } else {
                q4 += col
            }
        }
    }

    return q1 * q2 * q3 * q4
}

func getRobotGrid(robots []Robot) [][]int {

    grid := make([][]int, 0)

    for i := range FLOOR_HEIGHT {
        row := make([]int, 0)
        for j := range FLOOR_WIDTH {

            val := 0

            for _, robot := range robots {

                if robot.position.X == j && robot.position.Y == i {
                    val++
                }
            }
            row = append(row, val)
        }
        grid = append(grid, row)
    }
    return grid
}
func p2() int {
    lines := lib.ReadFile("day14.txt")

    robots := make([]Robot, 0)

    for _, line := range lines {
        robots = append(robots, NewRobot(lib.ParseIntsFromString(line)))
    }
    
    var bestGrid [][]int
    best := -1
    bestSecond := 0

    for second := range 15000 {

        positions := make(structures.Set[maths.Position], 0)
        for i := range robots {

            robots[i].position.PAdd(robots[i].velocity)

            robots[i].position.X %= FLOOR_WIDTH
            if robots[i].position.X < 0 {
                robots[i].position.X += FLOOR_WIDTH
            }
            robots[i].position.Y %= FLOOR_HEIGHT
            if robots[i].position.Y < 0 {
                robots[i].position.Y += FLOOR_HEIGHT
            }
        
            positions.Add(robots[i].position)
        }


        rowCount := make(map[int]int, 0)

        for _, pos := range positions {
            val, found := rowCount[pos.Y]
            if found {
                rowCount[pos.Y] = val + 1
            } else {
                rowCount[pos.Y] = 1
            }
        }

        for _, v := range rowCount {

            if v > best {
                best = v
                bestGrid = getRobotGrid(robots)
                bestSecond = second
            }
        }
    }

    fmt.Println("-------------\nBest Grid")
    SpecialPrint(bestGrid)

    return bestSecond + 1
}

func SpecialPrint(grid [][]int) {
	for _, row := range grid {
		for _, col := range row {
            if col == 0 {
                fmt.Print("  ")
            } else {
                fmt.Print(col, " ")
            }
		}
		fmt.Println()
	}
}