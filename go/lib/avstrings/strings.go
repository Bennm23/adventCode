package avstrings

import (
	"regexp"
	"strconv"
)

func ParseTextInParens(str string) string {
	var s string

	marked := false
	for _, c := range str {
		if c == ')' {
			break;
		}
		if marked {
			s = s + string(c)
		}
		if c == '(' {
			marked = true
			continue
		}
	}
	
	return s
}

func SplitTextToInts(str string) []int {
	intFinder := regexp.MustCompile(`[-]?[\d]+`)
	var ints []int

	foundIndices := intFinder.FindAllStringIndex(str, -1);

	for _, found := range foundIndices {
		val, err := strconv.Atoi(str[found[0]:found[1]])
		if err != nil {
			panic("Failed To Split To ints")
		}
		ints = append(ints, val)
	}

	return ints
}

func StringsToInts(arr []string) []int {
	var ints []int

	for _, val := range arr {
		v, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		ints = append(ints, v)
	}
	return ints
}