package main

import (
	"advent/lib"
	"advent/lib/avstrings"
	"fmt"
	"hash/fnv"
	"strings"
)

func hash(query string, groups []int) uint64 {
	h := fnv.New64a()

	fmt.Fprint(h, query, groups)

	return h.Sum64()
}

func count(query string, groups []int, scores *map[uint64]int64) int64 {
	if query == "" {
		if len(groups) == 0 {
			return 1
		}
		return 0
	}

	if len(groups) == 0 {
		if (strings.Contains(query, "#")) {
			return 0
		}
		return 1
	}

	h := hash(query, groups)

	if val, ok := (*scores)[h]; ok {
		return val
	}


	var sum int64 = 0

	//If the start of this string is a . then we can continue on as no block can start here
	if query[0] == '.' || query[0] == '?' {
		sum += count(query[1:], groups, scores)
	}

	//If the start is a # then this could be the start of a block
	if query[0] == '#' || query[0] == '?' {
		//If len(query) >= groups[0] then this could be a block
		//If query[:groups[0]] does not contain a .
		//If nums[0] == len(query) or query[nums[0]] != "#"
		if (groups[0] <= len(query) &&
		    !strings.Contains(query[:groups[0]], ".") &&
		    (groups[0] == len(query) || query[groups[0]] != '#')) {

				//If conditions are met, then we have found a block
				//Move to next group, and move one over in string because groups must have a gap

				//If there is not enough string left to fill the last group, just empty string
				if groups[0] + 1 > len(query) {
					sum += count("", groups[1:], scores)
				} else {
					sum += count(query[groups[0] + 1:], groups[1:], scores)
				}
		} 
	}

	(*scores)[h] = sum
	return sum
}

func main() {
	lib.RunAndPrintDuration(func() {solve()})//373027, 350609, 344343
}
func solve() {
	lines, err := lib.ReadFile("day12.txt")

	if err != nil {
		panic("Failed To Read Day12")
	}

	var p1, p2 int64 = 0, 0
	scores,scores2 := make(map[uint64]int64), make(map[uint64]int64)

	for _, line := range lines {
		split := strings.Split(line, " ")

		query, groups := split[0], avstrings.StringsToInts(strings.Split(split[1], ","))
		res := count(query, groups, &scores)
		p1 += res

		query2 := ""

		for i := 0; i < 5; i++ {

			query2 += query

			if i != 4 {
				query2 += "?"
			}

		}

		groups2 := lib.Repeat[int](groups, 5)

		p2 += count(query2, groups2, &scores2)
	}

	
	fmt.Println("Part 1 = ", p1)//7110
	fmt.Println("Part 2 = ", p2)//1566786613613

}