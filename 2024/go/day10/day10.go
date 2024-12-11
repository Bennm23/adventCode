package main

import (
	"advent/lib"
	"advent/lib/maths"
	"advent/lib/structures"
)

type VisitedSet = structures.Set[maths.Position]

func main() {
    lib.RunAndScore("Part 1", p1)//Score =  782. Total Time 678 us
    lib.RunAndScore("Part 2", p2)//Score = 1694. Total Time 716 us
}

func getData() ([][]int, []maths.Position) {

    grid := lib.ReadFileToTypeGrid("day10.txt", func(s string) []int {
        return lib.StringToInts(s)
    })

    trailheads := make([]maths.Position, 0)

    for rix, row := range grid {
        for cix, col := range row {
            if col == 0 {
                trailheads = append(trailheads, maths.NewPosition(rix, cix))
            }
        }
    }
    return grid, trailheads
}

func p1() int {
    sum := 0
    grid, trailheads := getData()

    for _, trailhead := range trailheads {

        visited := make(VisitedSet, 0)
        sum += dfs(grid, trailhead, &visited, false)
    }

    return sum
}

func p2() int {
    sum := 0
    grid, trailheads := getData()

    for _, trailhead := range trailheads {
        sum += bfs(grid, trailhead, true)
    }

    return sum
}

func dfs(grid [][]int, position maths.Position, visited *VisitedSet, allPaths bool) int {
    moves := []maths.Position{
        maths.NewPosition(-1, 0),
        maths.NewPosition(1, 0),
        maths.NewPosition(0, -1),
        maths.NewPosition(0, 1),
    }

    sum := 0
    currHeight := position.EvaluateFor(grid)

    if currHeight == 9 {
        return 1
    }

    for _, move := range moves {
        newPos := position.Add(move)

        if newPos.OutOfBounds(len(grid)) {
            continue
        }
        if !allPaths && visited.Contains(newPos) {
            continue
        }
        if newPos.EvaluateFor(grid) - currHeight != 1 {
            continue
        }
        visited.Add(newPos)
        sum += dfs(grid, newPos, visited, allPaths)
    }
    return sum
}


func bfs(grid [][]int, position maths.Position, allPaths bool) int {

    visited := VisitedSet{position}

    moves := []maths.Position{
        maths.NewPosition(-1, 0),
        maths.NewPosition(1, 0),
        maths.NewPosition(0, -1),
        maths.NewPosition(0, 1),
    }

    queue := structures.NewStack[maths.Position]()
    queue.Push(position)

    sum := 0

    for !queue.IsEmpty() {

        curr := queue.Pop()

        currHeight := curr.EvaluateFor(grid)
        if currHeight == 9 {
            sum += 1
            continue
        }

        for _, move := range moves {
            newPos := curr.Add(move)

            if newPos.OutOfBounds(len(grid)) {
                continue
            }
            if !allPaths && visited.Contains(newPos) {
                continue
            }
            if newPos.EvaluateFor(grid) - currHeight != 1 {
                continue
            }
            visited.Add(newPos)
            queue.Push(newPos)
        }
    }

    return sum
}