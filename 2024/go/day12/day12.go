package main

import (
	"advent/lib"
	"fmt"
)

func main() {
    lib.RunAndScore("Part 1", p1)
    lib.RunAndScore("Part 2", p2)
}

func p1() int {
    sum := 0

    grid := lib.ReadFileToGrid("day12_sample.txt")

    for _, row := range grid {
        for col := range row {

            fmt.Print(col, " ")
        }
        fmt.Println()
    }

    return sum
}
func p2() int {
    sum := 0

    return sum
}
