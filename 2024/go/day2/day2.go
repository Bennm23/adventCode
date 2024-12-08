package main

import (
	"advent/lib"
	"strconv"
	"strings"
)

func main() {
	lib.RunAndScore("Part 1", p1)//Score = 624. Total Time 1085 us
	lib.RunAndScore("Part 2", p2)//Score = 658. Total Time 1197 us
}

func getReports() [][]int {
	lines := lib.ReadFileToTypeGrid("day2.txt", func(s string) []int {

		var res []int
		splits := strings.Split(s, " ")

		for _, split := range splits {
			val, err := strconv.Atoi(split)
			if err != nil {
				panic(err)
			}
			res = append(res, val)
		}
		return res
	})
	return lines
}

func evaluateReport(report []int) bool {
	decreasing := report[0] > report[1]

	for cix, col := range report {
		if cix == len(report) - 1 {
			break
		}
		delta := col - report[cix + 1]

		if delta == 0 {
			return false
		}
		if lib.Absi(delta) > 3 {
			return false
		}
		if delta < 0 && decreasing {
			return false
		}
		if delta > 0 && !decreasing {
			return false
		}
	}
	return true
}

func getUnsafeReports(reports [][]int) [][]int {
	var unsafe [][]int

	for _, report := range reports {
		valid := evaluateReport(report)
		if !valid {
			unsafe = append(unsafe, report)
		}
	}
	return unsafe
}

func p1() int {
	reports := getReports()
	unsafe := getUnsafeReports(reports)

	return len(reports) - len(unsafe)
}

func p2() int {
	reports := getReports()
	unsafe := getUnsafeReports(reports)

	safe := len(reports) - len(unsafe)

	for _, report := range unsafe {
		for ix := 0; ix < len(report); ix++ {
			//Need to make a new array every time
			//Go will reuse the underlying array if it can contain the data
			cp := make([]int, len(report))
			copy(cp, report)
			valid := evaluateReport(append(cp[0:ix], cp[ix + 1:]...))
			if valid {
				safe += 1
				break
			}
		}
	}
	return safe
}