package main

import (
	"advent/lib"
	"advent/lib/maths"
	"strconv"
)

func main() {
    lib.RunAndScore("Part 1", p1)//Score = 222461.          Total Time 582 us
    lib.RunAndScore("Part 2", p2)//Score = 264350935776416. Total Time 24794 us
}

func getStones() []string {
    return lib.ReadOneLineToChunks("day11.txt", " ")
}

type DepthKey struct {
    stone int
    blink int
}
type Scores = map[DepthKey]int

func evaluate(stone int, blink int, results *Scores, maxBlinks int) int {

    if blink == maxBlinks {
        return 1
    }
    key := DepthKey{stone, blink}
    score, found := (*results)[key]
    if found {
        return score
    }

    stoneString := strconv.Itoa(stone)
    if stone == 0 {
        score = evaluate(1, blink+1, results, maxBlinks)
    } else if len(stoneString) % 2 == 0 {
        left := maths.ToInt(stoneString[:(len(stoneString) / 2)])
        right := maths.ToInt(stoneString[(len(stoneString) / 2):])
        score = evaluate(left, blink+1, results, maxBlinks) + evaluate(right, blink+1, results, maxBlinks)
    } else {
        newVal := stone * 2024
        score = evaluate(newVal, blink+1, results, maxBlinks)
    }
    (*results)[key] = score
    return score
}

func p1() int {
    sum := 0
    stones := getStones()

    scores := make(Scores, 0)

    for _, stone := range stones {
        sum += evaluate(maths.ToInt(stone), 0, &scores, 25)
    }

    return sum
}
func p2() int {
    sum := 0
    stones := getStones()

    scores := make(Scores, 0)

    for _, stone := range stones {
        sum += evaluate(maths.ToInt(stone), 0, &scores, 75)
    }

    return sum
}
