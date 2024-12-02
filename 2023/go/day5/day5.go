package main

import (
	"advent/lib"
	"math"
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
	// "regexp"
	// "strconv"
)
type Ranges struct {
	valMin, valMax, keyMin, keyMax int64
}

func parseRange(line string) Ranges {
	nums := strings.Split(line, " ")
	toStart, err := strconv.ParseInt(nums[0], 10, 64)
	if err != nil {
		panic("Failed TO Parse Start")
	}
	fromStart, err := strconv.ParseInt(nums[1], 10, 64)
	if err != nil {
		panic("Failed TO Parse Start")
	}
	length, err := strconv.ParseInt(nums[2], 10, 64)
	if err != nil {
		panic("Failed TO Parse Start")
	}

	return Ranges{
		valMin: toStart,
		valMax: toStart + length,
		keyMin: fromStart,
		keyMax: fromStart + length,
	}
}

func parseMapping(scanner *bufio.Scanner) []Ranges {
	var ranges []Ranges
	scanner.Scan()//Read map line
	for scanner.Scan() {
		if scanner.Text() == "" {
			break;
		}
		ranges = append(ranges, parseRange(scanner.Text()))
	}
	
	return ranges
}

func main() {

	// lib.RunAndPrintDuration(func() {
		file, err := os.Open(lib.FILE_PATH + "day5.txt")
		if err != nil {
			panic("Couldnt open day 5")
		}

		defer file.Close()

		var maps [][]Ranges
		var seeds []int64

		scanner := bufio.NewScanner(file)

		scanner.Scan()
		seedSplit := strings.Split(scanner.Text(), " ")
		for i, seed := range seedSplit {
			if i == 0 {
				continue
			}
			hold, err := strconv.ParseInt(seed, 10, 64)
			if err != nil {
				panic("Failed to parse seed to long")
			}
			seeds = append(seeds, hold)
		}

		scanner.Scan()//Skip whitespace

		maps = append(maps, parseMapping(scanner))
		maps = append(maps, parseMapping(scanner))
		maps = append(maps, parseMapping(scanner))
		maps = append(maps, parseMapping(scanner))
		maps = append(maps, parseMapping(scanner))
		maps = append(maps, parseMapping(scanner))
		maps = append(maps, parseMapping(scanner))

		file.Close()

		fmt.Println("Part 1 = ", part1(seeds, maps))//173706076

		

	// })
}

func (r Ranges) inKeyRange(query int64) bool {
	return r.keyMin <= query && query <= r.keyMax
}


func walkDown(key int64, maps [][]Ranges) int64 {
	nextKey := key
	for _, mapping := range maps {

		for _, rangeVal := range mapping {

			if rangeVal.inKeyRange(nextKey) {

				nextKey = rangeVal.valMin + (nextKey - rangeVal.keyMin)
				break
			}
		}
	}
	return nextKey
}

func part1(seeds []int64, maps [][]Ranges) int64 {

	//Loop through seed
	//For each seed, walk through the maps to the location
	var bestLocation int64;
	bestLocation = math.MaxInt64

	for _, seed := range seeds {
		
		res := walkDown(seed, maps)
		if res < bestLocation {
			bestLocation = res;
		}
	}

	return bestLocation
}
