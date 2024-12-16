package main

import (
	"advent/lib"
	"advent/lib/maths"
	"advent/lib/structures"
)

func main() {
    lib.RunAndScore("Part 1", p1)//Score = 1370100. Total Time 547787 us
    lib.RunAndScore("Part 2", p2)//Score = 818286. Total Time 592298 us
}

func p1() int {
    sum := 0

    regions := findRegions()

    for _, region := range regions {
        perimeter := 0
        for _, node := range region {

            for _, move := range VALID_MOVES {
                if !region.Contains(move.Add(node)) {
                    perimeter += 1
                }
            }
        }
        sum += len(region) * perimeter
    }

    return sum
}

func p2() int {
    sum := 0

    regions := findRegions()
    for _, region := range regions {
        sum += len(region) * countCorners(&region)
    }
    return sum
}

type VisitedSet = structures.Set[maths.Position]
var VALID_MOVES = []maths.Position{
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
}

func findRegions() []VisitedSet {
    grid := lib.ReadFileToGrid("day12.txt")
    regions := make([]VisitedSet, 0)
    visited := make(VisitedSet, 0)

    for rix, row := range grid {
        for cix, col := range row {

            newPos := maths.NewPosition(rix, cix)
            if visited.Contains(newPos) {
                continue
            }
            localVisits := make(VisitedSet, 0)
            walk(grid, col, newPos, &visited, &localVisits)
            regions = append(regions, localVisits)
        }
    }
    return regions
}

func walk(grid [][]rune, letter rune, position maths.Position, visited *VisitedSet, localVisits *VisitedSet) {

    if localVisits.Contains(position) {
        return
    }
    visited.Add(position)
    localVisits.Add(position)

    for _, move := range VALID_MOVES {

        newPos := position.Add(move)
        if !newPos.InBounds(len(grid)) {
            continue
        }
        newChar := grid[newPos.X][newPos.Y]
        if newChar != letter {
            continue
        }
        walk(grid, newChar, newPos, visited, localVisits)
    }
}
func countCorners(region *VisitedSet) int {
    sum := 0

    for _, node := range *region {
        neighbors := maths.GetNeighbors(node)

        for _, nb := range [][]maths.Position{
            {neighbors[maths.N], neighbors[maths.E], neighbors[maths.NE]},
            {neighbors[maths.S], neighbors[maths.E], neighbors[maths.SE]},
            {neighbors[maths.S], neighbors[maths.W], neighbors[maths.SW]},
            {neighbors[maths.N], neighbors[maths.W], neighbors[maths.NW]},
        } {
            //Corner if the directional squares are not in region
            if !region.Contains(nb[0]) && !region.Contains(nb[1]) {
                sum++
            }
            //Or corner if directional squares are in and diagonal square is not
            if region.ContainsAll(nb[0], nb[1]) && !region.Contains(nb[2]) { 
                sum++
            }
        }
    }

    return sum
}
