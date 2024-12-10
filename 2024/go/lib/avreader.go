package lib

import (
	"advent/lib/maths"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

const FILE_PATH = "/home/benn/CODE/adventCode/2024/inputs/"
const LAPTOP_PATH = "/home/benn-mellinger/CODE/adventCode/2024/inputs/"

// Input file name and return array of lines
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
	file, err := os.Open(prefix + name)
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
	groups = append(groups, temps) //Catch the last group

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
func ReadFileToTypeGrid[T any](name string, convert func(string) []T) [][]T {
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
		grid = append(grid, convert(scanner.Text()))
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return grid
}

func ReadFileToTypeVec[T any](name string, convert func(string) T) []T {
	prefix := prefix()

	fmt.Println("OPENING FILE AT ", (prefix + name))
	file, err := os.Open(prefix + name)
	if err != nil {
		panic(fmt.Sprintf("Failed to Open %s", name))
	}

	defer file.Close()

	var vals []T

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		vals = append(vals, convert(scanner.Text()))
	}

	return vals
}

func Absi(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
		if len(scanner.Text()) == 0 {
			continue
		}
		text := scanner.Text()
		lines = append(lines, replacer(text))
	}

	return lines, scanner.Err()
}

type Formatter func(string) string

type Solver func()
type Scorer func() int

func RunAndScore(title string, solver Scorer) {
	start := time.Now().UnixMicro()
	score := solver()
	fmt.Printf("%s: Score = %d. Total Time %d us\n", title, score, (time.Now().UnixMicro() - start))
}

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

func FindAllMatches(regex string, source string) []string {
	matcher := regexp.MustCompile(regex)
	results := matcher.FindAllString(source, -1)

	return results
}

func EvaluateMatch[T any](regex, source string, evaluate func([]string) T) T {

	matcher := regexp.MustCompile(regex)
	result := matcher.FindAllString(source, -1)

	return evaluate(result)
}

func RemoveStrBetweenOrAfter(text string, before string, after string) string {

	match_string := fmt.Sprintf("%s.*?%s|%s.*", before, after, before)
	matcher := regexp.MustCompile(match_string)

	return matcher.ReplaceAllString(text, "_X_")
}

func SplitStringToInts(str string, delimeter string) []int {
	split := strings.Split(str, delimeter)

	var vals []int

	for _, s := range split {
		vals = append(vals, maths.ToInt(s))
	}
	return vals
}

const DEBUG = true

func Log(a ...any) {
	if !DEBUG {
		return
	}
	fmt.Println(a...)
}

func Lognl(a ...any) {
	if !DEBUG {
		return
	}
	fmt.Print(a...)
}
