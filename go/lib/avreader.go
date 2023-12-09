package lib

import (
	"bufio"
	"fmt"
	"os"
	"time"
)
const FILE_PATH = "/home/benn/CODE/adventCode/";
const LAPTOP_PATH ="/home/bennmellinger/CODE/adventCode/";


//Input file name and return array of lines
func ReadFile(name string) ([]string, error) {
	prefix := FILE_PATH
	_, err := os.Open("/home/benn")

	if err != nil {
		prefix = LAPTOP_PATH
	}
	fmt.Println("OPENING FILE AT ", (prefix + name))

	file, err := os.Open(prefix + name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
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