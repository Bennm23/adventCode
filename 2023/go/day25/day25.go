package main

import (
	"advent/lib"
	"fmt"
	"strings"
)

func main() {

	lines := lib.ReadFile("sample")

	for _, line := range lines {
		split := strings.Split(line, ": ")

		left := split[0]

		rights := strings.Split(split[1], " ")

		fmt.Println(left, " <=> ", rights)
	}
}