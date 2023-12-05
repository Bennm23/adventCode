package main

import (
	"advent/lib"
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	lib.RunAndPrintDuration(func() {//649-724
		symbolFind := regexp.MustCompile(`[^(\d | \.)]`)
		grid, err := lib.ReadFileWithReplace("day3.txt", func(s string) string {

			return symbolFind.ReplaceAllString(s, "*")
		})

		if err != nil {
			panic("Failed to parse day3.txt")
		}

		fmt.Println("PART 1 = ", part1(grid))//526404
		fmt.Println("PART 2 = ", part2())//84399773
	})//2908-3001

}

func part2() int {
	starFind := regexp.MustCompile(`\*`)
	isNumber := regexp.MustCompile(`[0-9]+`)

	grid, err := lib.ReadFile("day3.txt")
	if err != nil {
		panic("Failed to parse Day3")
	}

	var starIndices [][][]int

	for _, row := range grid {

		starIndices = append(starIndices, starFind.FindAllStringIndex(row, -1))
	}

	//For all rows in start index, 
	sum := 0;

	for row, stars := range starIndices {
		if len(stars) == 0 {
			continue
		}
		
		for _, star := range stars {
			var foundNumbers []int

			foundNumbers = append(foundNumbers, checkCenterForStar(&star, grid[row], isNumber)...)

			if row > 0 {
				foundNumbers = append(foundNumbers,checkLineForStar(&star, grid[row - 1], isNumber)...)
			}
			if row < len(grid) {
				foundNumbers = append(foundNumbers, checkLineForStar(&star, grid[row + 1], isNumber)...)
			}
			if len(foundNumbers) == 2 {
				sum += foundNumbers[0] * foundNumbers[1]
			}
		}

	}
	return sum
}
func checkLineForStar(star *[]int, line string, isNumber *regexp.Regexp) []int {
	//Returns all found numbers on this centerline (max of 2)
	var numbers []int

	numsInLine := isNumber.FindAllStringIndex(line, -1)

	for _, nums := range numsInLine {

		//If intersects left boundary
		if nums[0] <= (*star)[0] && nums[1] >= (*star)[0] {
			res, err := strconv.Atoi(string(line[nums[0]:nums[1]]))
			if err != nil {
				fmt.Println(err)
				panic("Failed to read line 1")
			}
			numbers = append(numbers, res)
			continue
		}

		//If intersects right boundary
		if nums[0] <= (*star)[1] && nums[1] >= (*star)[0] {
			res, err := strconv.Atoi(string(line[nums[0]:nums[1]]))
			if err != nil {
				fmt.Println(err)
				panic("Failed to read line 1")
			}
			numbers = append(numbers, res)
			continue
		}

	}

	return numbers
}

func checkCenterForStar(star *[]int, line string, isNumber *regexp.Regexp) []int {
	//Returns all found numbers on this centerline (max of 2)
	var numbers []int
	
	//Check centerline left
	if (*star)[0] > 0 && isNumber.MatchString(string(line[(*star)[0] - 1])) {
		num, err := readLeftToInt(line, (*star)[0] - 1, isNumber)
		if err != nil {
			fmt.Println(err)
			panic("Failed reading left")
		}
		numbers = append(numbers, num)
	}

	//Check centerline right

	if (*star)[1] < len(line) && isNumber.MatchString(string(line[(*star)[1]])) {
		num, err := readRightToInt(line, (*star)[1], isNumber)
		if err != nil {
			fmt.Println(err)
			panic("Failed reading right")
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func readLeftToInt(line string, rightIndex int, isNumber *regexp.Regexp) (int, error) {

	leftIndex := rightIndex
	//walk left and append to str. If non digit found, break
	for i := rightIndex; i >= 0; i-- {

		if isNumber.MatchString(string(line[i:i+1])) {
			leftIndex = i
		} else {
			break
		}
	}

	return strconv.Atoi(line[leftIndex:rightIndex + 1])
}

func readRightToInt(line string, leftIndex int, isNumber *regexp.Regexp) (int, error) {

	rightIndex := leftIndex
	//walk right. If non digit found, break
	for i := leftIndex; i < len(line); i++ {

		if isNumber.MatchString(string(line[i:i+1])) {
			rightIndex = i
		} else {
			break
		}
	}

	return strconv.Atoi(line[leftIndex: rightIndex + 1])
}

func part1(grid []string) int {
	//Find all indices of all numbers
	//Loop through each index in each row
	//At each index, check left - 1 to right + 1

	numberFind := regexp.MustCompile(`\b[\d]+`)

	var numberIndices [][][]int

	for _, row := range grid {
		numberIndices = append(numberIndices, numberFind.FindAllStringIndex(row, -1))
	}

	var sum int
	for row, rowNumbers := range numberIndices {

		if len(rowNumbers) == 0 {
			continue;
		}

		//Loop through all the indices of the numbers in this row
		for _, number := range rowNumbers {

			//If there is a symbol around this number index, add to sum
			if symbolIsAround(number, grid, row) {
				num, err := strconv.Atoi(grid[row][number[0]:number[1]])
				if err != nil {
					panic("Failed to convert to int")
					
				}
				sum += num
			}
		}
	}

	return sum
}

const STAR_VAL = 42

func symbolIsAround(number []int, inputGrid []string, row int) bool {

	if checkCenterline(inputGrid[row], number) {

		return true
	}

	if row > 0 && checkLine(inputGrid[row - 1], number){

		return true
	}

	if row < len(inputGrid) - 1 && checkLine(inputGrid[row + 1], number){

		return true
	}

	return false
}

func checkLine(line string, number []int) bool {

	//Loop for i - 1 to i + 1
	for i := lib.Max(0, number[0] - 1); i < lib.Min(len(line), number[1] + 1); i++ {

		if (line[i] == STAR_VAL) {
			return true
		}
		
	}
	return false
}


func checkCenterline(line string, number []int) bool {
	//Check centerline left
	if number[0] > 0 && line[number[0] - 1] == STAR_VAL {
		return true
	}

	//Check centerline right
	if number[1] < len(line) && line[number[1]] == STAR_VAL {
		return true
	}

	return false
}