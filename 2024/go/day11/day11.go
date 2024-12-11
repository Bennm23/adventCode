package main

import (
	"advent/lib"
	"advent/lib/maths"
	"strconv"
)

const BLINKS = 75

func main() {
    lib.RunAndScore("Part 1", p1)//Score = 222461. Total Time 582 us
    lib.RunAndScore("Part 2", p2)//Score = 264350935776416. Total Time 38105 us
}

func getStones() []string {
    return lib.ReadOneLineToChunks("day11.txt", " ")
}

type DepthKey struct {
    stone string
    blink int
}
type Scores = map[DepthKey]int

func evaluate(stone string, blink int, results *Scores, maxBlinks int) int {

    if blink == maxBlinks {
        return 1
    }
    key := DepthKey{stone, blink}
    score, found := (*results)[key]
    if found {
        return score
    }

    if maths.ToInt(stone) == 0 {
        score = evaluate("1", blink+1, results, maxBlinks)
    } else if len(stone) % 2 == 0 {

        left := strconv.Itoa(maths.ToInt(stone[:(len(stone) / 2)]))
        right := strconv.Itoa(maths.ToInt(stone[(len(stone) / 2):]))
        score = evaluate(left, blink+1, results, maxBlinks) + evaluate(right, blink+1, results, maxBlinks)
    } else {
        newVal := maths.ToInt(stone) * 2024
        score = evaluate(strconv.Itoa(newVal), blink+1, results, maxBlinks)
    }
    (*results)[key] = score
    return score
}

func p1() int {
    sum := 0
    stones := getStones()
    //use bits
    //1. set bit to1
    //2. if end bit != 1, split how?
    //3. left shift value 11
    //FULLY EXPAND FROM BACK TO FRONT, use bytes? local array for each stone only append at end

    scores := make(Scores, 0)

    for _, stone := range stones {
        sum += evaluate(stone, 0, &scores, 25)
    }

    return sum
}
func p2() int {
    sum := 0
    stones := getStones()

    scores := make(Scores, 0)

    for _, stone := range stones {
        sum += evaluate(stone, 0, &scores, 75)
    }

    return sum
}
