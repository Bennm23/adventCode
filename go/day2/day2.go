package main

import (
	"advent/lib"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day2")

	lines, err := lib.ReadFile("day2.txt")

	if err != nil {
		panic("Failed to parse day2.txt")
	}
	
	p1, p2 := solve(lines)

	fmt.Println("PART 1 = ", p1)//2879
	fmt.Println("PART 2 = ", p2)


}

type ColorMap map[string]int

const GREEN string = "green"
const RED string = "red"
const BLUE string = "blue"

const B1 int = 14
const G1 int = 13
const R1 int = 12

func createColorMap() ColorMap {

	return ColorMap {
		BLUE  : 0,
		RED   : 0,
		GREEN : 0,
	}

}

func getColorKeyValue(intFind *regexp.Regexp, count string) (colorKey string, intValue int) {
	intIndice := intFind.FindStringIndex(count)
	intVal, err := strconv.Atoi(count[intIndice[0]:intIndice[1]])
	if err != nil {
		panic("Couldn't convert int")
	}
	key := count[intIndice[1] + 1:]

	return key, intVal
}

func buildScoreMap(counts []string, intFind *regexp.Regexp) ColorMap {
	colorCount := createColorMap()

	for _, count := range counts {

		key, value := getColorKeyValue(intFind, count)
		colorCount[key] = value
	}

	return colorCount
}

func gamePasses(game []string, intFind  *regexp.Regexp) bool {
	for _, counts := range game {
		scoreMap := buildScoreMap(strings.Split(counts, ","), intFind)

		if (scoreMap[GREEN] > G1 || scoreMap[BLUE] > B1 || scoreMap[RED] > R1) {
			return false
		}
	}

	return true
}

func getGamePower(game string, intFind *regexp.Regexp, scoreSplit *regexp.Regexp) int {
	//find min rgb value and r*g*b return

	maxScoreMap := ColorMap {
		BLUE  : math.MinInt,
		RED   : math.MinInt,
		GREEN : math.MinInt,
	}

	for _, score := range scoreSplit.Split(game, -1) {
		key, val := getColorKeyValue(intFind, score)

		if val > maxScoreMap[key] {
			maxScoreMap[key] = val
		}

	}
	fmt.Println(maxScoreMap)

	return maxScoreMap[GREEN] * maxScoreMap[BLUE] * maxScoreMap[RED]
}

func solve(lines []string) (int, int) {
	intFind := regexp.MustCompile(`\d+`)
	scoreSplit := regexp.MustCompile(";|,")

	var score int
	var powerScore int

	for round, line := range lines {
		line = strings.Split(line, ":")[1]
		fmt.Println(line)

		if gamePasses(strings.Split(line, ";"), intFind) {
			score += round + 1
		}
	
		val := getGamePower(line, intFind, scoreSplit)
		powerScore += val
		fmt.Println(val)
	}

	return score, powerScore
}