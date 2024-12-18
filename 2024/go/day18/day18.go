package main

import (
"advent/lib"
"advent/lib/maths"
"advent/lib/structures"
"fmt"
"math"
"slices"
"strconv"
)

func main() {
    lib.RunAndScore("Part 1", p1)//Result =   292 : Total Time  580441 us
    lib.RunAndScore("Part 2", p2)//Result = 58,44 : Total Time 5008379 us
}

const WIDTH int = 71
const HEIGHT int = 71
const MEMSIZE int = 1024

func buildMemorySpace(steps int) ([]maths.Position, [][]bool) {
    memory := lib.ReadFileToTypeVec("day18.txt", func(s string) maths.Position {
        ints := lib.SplitStringToInts(s, ",")
        return maths.NewPosition(ints[0], ints[1]);
    })

    memorySpace := maths.InitTypeGrid(false, HEIGHT, WIDTH)

    for i := range steps {

        pos := memory[i]
        memorySpace[pos.Y][pos.X] = true
    }
    return memory, memorySpace
}

func p1() int {
    _, memSpace := buildMemorySpace(MEMSIZE)
    return explore(memSpace)
}

type VisitedSet = structures.Set[maths.Position]

func explore(memorySpace [][]bool) int {

    start := maths.NewPosition(0, 0)

    distanceMap := make(map[maths.Position]int, 0)
    previousNode := make(map[maths.Position]*maths.Position, 0)

    explorable := make(VisitedSet, 0)

    for rix, row := range memorySpace {
        for cix, col := range row {
            if col {
                continue
            }
            pos := maths.NewPosition(cix, rix)
            distanceMap[pos] = math.MaxInt
            previousNode[pos] = nil
            explorable.Add(pos)
        }
    }
    distanceMap[start] = 0

    for len(explorable) != 0 {

        slices.SortFunc(explorable, func(a, b maths.Position) int {
            if distanceMap[a] < distanceMap[b] {
                return -1
            } else if distanceMap[a] > distanceMap[b] {
                return 1
            }
            return 0
        })

        bestOption := explorable[0]
        explorable = explorable[1:]

        for _, move := range maths.HORIZONTAL_MOVES {
            newPos := move.Add(bestOption);
            if !explorable.Contains(newPos) {//Then already explored
                continue
            }
            newDistance := distanceMap[bestOption] + 1

            if newDistance < distanceMap[newPos] {
                distanceMap[newPos] = newDistance
                previousNode[newPos] = &bestOption
            }

        }
    }

    sum := 0;

    curr := maths.NewPosition(WIDTH - 1, WIDTH - 1)
    cp := &curr

    positions := make(VisitedSet, 0)

    for *cp != start {

        sum += 1
        hld, found := previousNode[*cp]
        if !found || hld == nil {
            fmt.Println("NO SOLUTION")
            return 0
        }
        positions.Add(*cp)
        cp = hld
    }

    // for rix, row := range memorySpace {
    //     for cix, col := range row {
    //         if positions.Contains(maths.NewPosition(cix, rix)) {
    //             fmt.Print("O ")
    //         } else {
    //             if col {
    //                 fmt.Print("# ")
    //             } else {
    //                 fmt.Print(". ")
    //             }
    //         }
    //     }
    //     fmt.Println()
    // }
   
    return sum
}

func p2() string {
    memory, memSpace := buildMemorySpace(MEMSIZE)
    return progressiveSearch(memSpace, memory, MEMSIZE, len(memory) / 10)
}

func progressiveSearch(memSpace [][]bool, memory []maths.Position, index, searchWidth int) string {
    for i:=index; i < len(memory); i++ {
        newPos := memory[i]
        memSpace[newPos.Y][newPos.X] = true
        if i % searchWidth != 0 {
            continue
        }

        fmt.Println("I = ", i)
        val := explore(memSpace)
        if val == 0 {
            if searchWidth == 1 {
                return strconv.Itoa(memory[i].X) + "," + strconv.Itoa(memory[i].Y)
            } else {
                resetI := i - searchWidth
                for j := i; j >= resetI; j-- {
                    memSpace[memory[j].Y][memory[j].X] = false
                }
                searchWidth = maths.Max(1, searchWidth / 5)
                return progressiveSearch(memSpace, memory, resetI - (resetI % searchWidth), searchWidth)
            }
        }
    }

    return "NOT FOUND"
}
