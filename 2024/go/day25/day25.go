package main

import (
	"advent/lib"
	"advent/lib/structures"
)

func main() {
    lib.RunAndScore("Part 1", p1)//Part 1: Result = 3356 : Total Time 7199 us
}

type Keys = []structures.Vector[int]
type Locks = []structures.Vector[int]

func buildInput() (Keys, Locks, int) {
    var grids [][][]rune

    groups := lib.ReadFileToGroups("day25.txt", "");

    for _, g := range groups {
        group := make([][]rune, 0)
        for _, l := range g {

            row := make([]rune, 0);
            for _, c := range l {
                row = append(row, c)
            }
            group = append(group, row)
        }
        grids = append(grids, group)
    }
    locks := make(Locks, 0)
    keys := make(Keys, 0)
    gridHeight := len(grids[0])

    for _, grid := range grids {
    
        isLock, heights := getGridInfo(grid)
        if isLock {
            locks = append(locks, heights)
        } else {
            keys = append(keys, heights)
        }
    }
    return keys, locks, gridHeight
}

func p1() int {
    keys, locks, gridHeight := buildInput()

    goodKeys := 0
    for _, key := range keys {
        lockLoop: for _, lock := range locks {

            added := key.Plus(lock)
            for _, res := range added {
                if res + 2 > gridHeight {
                    continue lockLoop
                }
            }
            goodKeys += 1
        }
    }

    return goodKeys
}

func getGridInfo(grid [][]rune) (bool, []int) {
    heightMap := make([]int, len(grid[0]))

    isLock := true
    for _, c := range grid[0] {
        if c != '#' {
            isLock = false
            break
        }
    }

    for c := range len(grid[0]) {

        heightMap[c] = 0

        if isLock {
            for rix, row := range grid {
                if row[c] != '#' {
                    break
                }
                heightMap[c] = rix
            }
        } else {
            for rix := len(grid) - 1; rix >= 0; rix-- {
                if grid[rix][c] != '#' {
                    break
                }
                heightMap[c] = len(grid) - 1 - rix
            }
        }
    }

    return isLock, heightMap
}