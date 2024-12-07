package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)
const FILE_PATH = "/home/benn/CODE/adventCode/2023/";
const LAPTOP_PATH ="/home/benn-mellinger/CODE/adventCode/2023/";


//Input file name and return array of lines
func ReadFile(name string) []string {
	prefix := prefix()

	fmt.Println("OPENING FILE AT ", (prefix + name))
	file, err := os.Open(prefix + name)
	if err != nil {
		panic(fmt.Sprintf("Failed to Open %s", name))
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(fmt.Sprintf("Scanner Err %s", scanner.Err().Error()))
	}

	return lines
}

func ReadOneLineToChunks(name, seperator string) []string {
	prefix := prefix()

	fmt.Println("OPENING FILE AT ", (prefix + name))
	file, err := os.Open(prefix + name)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var line string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
	}

	return strings.Split(line, seperator)

}

func ReadFileToGroups(name, delimeter string) [][]string {
	prefix := prefix()

	fmt.Println("OPENING FILE AT ", (prefix + name))
	file, err := os.Open(FILE_PATH + name)
	if err != nil {
		panic("Failed To Open File")
	}

	defer file.Close()

	var groups [][]string

	scanner := bufio.NewScanner(file)

	temps := make([]string, 0)
	for scanner.Scan() {
		if scanner.Text() == delimeter {
			
			groups = append(groups, temps)
			temps = make([]string, 0)
			continue
		}
		temps = append(temps, scanner.Text())
	}
	groups = append(groups, temps)//Catch the last group

	return groups
}

func prefix() string {
	prefix := FILE_PATH
	_, err := os.Open("/home/benn")

	if err != nil {
		prefix = LAPTOP_PATH
	}
	return prefix
}

func ReadFileToGrid(name string) [][]rune {
	prefix := prefix()

	fmt.Println("OPENING FILE AT ", (prefix + name))
	file, err := os.Open(prefix + name)
	if err != nil {
		panic(fmt.Sprintf("Failed to Open %s", name))
	}

	defer file.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var line []rune
		for _, r := range scanner.Text() {
			line = append(line, r)
		}
		grid = append(grid, line)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return grid
}
func ReadFileToTypeGrid[T any](name string, convert func(rune)T) [][]T {
	prefix := prefix()

	fmt.Println("OPENING FILE AT ", (prefix + name))
	file, err := os.Open(prefix + name)
	if err != nil {
		panic(fmt.Sprintf("Failed to Open %s", name))
	}

	defer file.Close()

	var grid [][]T

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var line []T
		for _, r := range scanner.Text() {
			line = append(line, convert(r))
		}
		grid = append(grid, line)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return grid
}

func ReadFileWithReplace(name string, replacer Formatter) ([]string, error) {
	file, err := os.Open(FILE_PATH + name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, replacer(text))
	}

	return lines, scanner.Err()
}

type Formatter func(string) string

type Solver func()

func RunAndPrintDuration(solver Solver) {
	start := time.Now().UnixMicro()
	solver()
	fmt.Println("Duration = ", (time.Now().UnixMicro() - start))
}
func RunAndPrintDurationMillis(solver Solver) {
	start := time.Now().UnixMilli()
	solver()
	fmt.Println("Duration = ", (time.Now().UnixMilli() - start))
}

func Max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func Min(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

func CopyMap[K comparable, V any](copy map[K]V) map[K]V {
	cp := make(map[K]V)

	for k, v := range copy {
		cp[k] = v
	}

	return cp
}

func Contains[K comparable](search []K, val K) bool {
	for _, v := range search {
		if v == val {
			return true
		}
	}
	return false
}

func Repeat[T any](arr []T, repeats int) []T {
	var res []T

	for i := 0; i < repeats; i++ {

		res = append(res, arr...)
	}

	return res
}

type AnyMap[T comparable, R any] map[T]R

func (mp AnyMap[T, R]) ValueSet() []R {
	values := []R{}

	for _, val := range mp {
		values = append(values, val)
	}

	return values
}
