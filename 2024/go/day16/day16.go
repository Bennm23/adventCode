package main

import (
	"advent/lib"
	"advent/lib/maths"
	"fmt"
	"math"
)

func main() {
    lib.RunAndScore("Part 1", p1)
    lib.RunAndScore("Part 2", p2)
}

type Location struct {
    position maths.Position
    direction maths.Position
}

func (loc Location) turnClockwise() Location {
    location := Location{}
    location.position = loc.position
    if loc.direction.X == -1 {
        location.direction = maths.NewPosition(0, 1)
    } else if loc.direction.Y == 1 {
        location.direction = maths.NewPosition(1, 0)
    } else if loc.direction.X == 1 {
        location.direction = maths.NewPosition(0, -1)
    } else if loc.direction.Y == -1 {
        location.direction = maths.NewPosition(-1, 0)
    }
    return location
}

func p1() int {
    sum := 0

    grid := lib.ReadFileToGrid("day16_sample.txt")

    var startPos, goal Location

    for rix, row := range grid {
        for cix, col := range row {
            fmt.Print(string(col))
        
            if col == 'S' {
                startPos = Location{maths.Position{rix, cix}, maths.Position{0, 1}}
            } else if col == 'E' {
                goal = Location{maths.Position{rix, cix}, maths.Position{0, 0}}
            }
        }
        fmt.Println()
    }

    fmt.Println("Start Pos = ", startPos)
    fmt.Println("Goal = ", goal)

    // explorationSet := structures.NewStack[Location]()
    // explorationSet.Push(startPos)


    currPos := startPos

    for currPos != goal {
    // for !explorationSet.IsEmpty() {

        // currPos := explorationSet.Pop()

        fmt.Println("At Pos = ", currPos)
        options := make([]Location, 0)

        for _, move := range maths.HORIZONTAL_MOVES {

            //Turn
            if move != currPos.direction {
                newLoc := Location{currPos.position, move}
                options = append(options, newLoc)
            }

            newPos := move.Add(currPos.position)
            if !newPos.InBounds(len(grid)) {
                continue
            }
            newLoc := Location{currPos.position, currPos.direction}
            options = append(options, newLoc)
        }

        // options = append(options, currPos.turnClockwise())

        minCost := math.MaxFloat64
        var bestLoc Location
        for _, option := range options {
            cost := 1 + option.position.Distance(goal.position)
            if option.position == currPos.position {
                cost = 1000 + option.position.Distance(goal.position)
                sum += 1000
            } else {
                sum += 1
            }
            if cost < minCost {
                bestLoc = option
            }
        }
        currPos = bestLoc
        // explorationSet.Push(bestLoc)
    }

    return sum
}
func p2() int {
    sum := 0

    return sum
}
