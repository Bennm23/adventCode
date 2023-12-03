package lib

import (
	"bufio"
	"os"
)
const FILE_PATH = "/home/benn/CODE/adventCode/";

//Input file name and return array of lines
func ReadFile(name string) ([]string, error) {
	file, err := os.Open(FILE_PATH + name)
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