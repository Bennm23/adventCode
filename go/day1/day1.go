package main

import (
	"advent/lib"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now().UnixMicro()

	lines, err  := lib.ReadFile("day1.txt")

	if err != nil {
		panic("Error Reading day1.txt")
	}

	re := regexp.MustCompile("[0-9]")
	fmt.Println("PART 1 = ", part1(lines, re))//55130
	fmt.Println("PART 2 = ", part2(lines, re))//54985

	fmt.Println("Duration = ", (time.Now().UnixMicro() - start))//1655-1773
	
}

func part1(lines []string, re *regexp.Regexp) int {
	var sum int = 0;
	for _, line := range lines {
		sum += getCombinedValue(line, re)
	}

	return sum
}

func part2(lines []string, re *regexp.Regexp) int {
	type Map map[string]string
	replacements := Map{
		"one": "o1e",
		"two": "t2o",
		"three": "t3e",
		"four": "f4r",
		"five": "f5e",
		"six": "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine": "n9e",
	}

	var sum int = 0;

	for _, line := range lines {
		for k, v := range replacements {
			line = strings.ReplaceAll(line, k, v)
		}
		sum += getCombinedValue(line, re)
	}

	return sum
}

func getCombinedValue(line string, re *regexp.Regexp) int {

	matches := re.FindAllString(line, -1)

	combined := matches[0] + matches[len(matches)-1]

	total, err := strconv.Atoi(combined)
	if err != nil {
		panic("Atoi Failed")
	}

	return total
}