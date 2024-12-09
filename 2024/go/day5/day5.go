package main

import (
	"advent/lib"
	"slices"
	"strings"
)

func main() {
    lib.RunAndScore("Part 1", p1)//Total 6041, Laptop Time 1068 us
    lib.RunAndScore("Part 2", p2)//Total 4884, Laptop Time 1333 us
}

type Ruleset map[int][]int
type Pagelist [][]int

func buildInput() (Ruleset, Pagelist) {
    lines := lib.ReadFile("day5.txt")

    ruleset := make(Ruleset, 0)
    pages := make(Pagelist, 0)

    for _, line := range lines {

        if strings.Contains(line, "|") {

            vals := lib.SplitStringToInts(line, "|")
            
            _, exists := ruleset[vals[0]]
            if exists {
                ruleset[vals[0]] = append(ruleset[vals[0]], vals[1])
            } else {
                ruleset[vals[0]] = []int{vals[1]}
            }
        } else if strings.Contains(line, ",") {
            vals := lib.SplitStringToInts(line, ",")
            pages = append(pages, vals)
        }
    }

    return ruleset, pages
}

func getBadUpdates(ruleset Ruleset, pages Pagelist) ([][]int, int) {

    bad := make([][]int, 0)
    goodSum := 0

    for _, page := range pages {

        valid := true

        outer:
        for i := len(page) - 1; i > -1; i-- {

            rules, found := ruleset[page[i]]
            if !found {
                continue
            }
            for j := 0; j < i; j++ {
                if lib.Contains(rules, page[j]) {
                    bad = append(bad, page)
                    valid = false
                    break outer
                }
            }
        }
        if valid {
            goodSum += page[len(page) / 2]
            
        }
    }
    return bad, goodSum
}

func p1() int {

    ruleset, pages := buildInput()
    _, goodSum := getBadUpdates(ruleset, pages)
    return goodSum
}
func p2() int {
    sum := 0
    ruleset, pages := buildInput()

    badUpdates, _ := getBadUpdates(ruleset, pages)

    for _, bad := range badUpdates {
        slices.SortFunc(bad, func(a int, b int) int {
            rules, found := ruleset[b]

            if !found {
                return 0
            }
            if lib.Contains(rules, a) {
                return -1
            }
            return 0
        });
    }

    for _, bad := range badUpdates {
        sum += bad[len(bad) / 2]
    }
    return sum
}
