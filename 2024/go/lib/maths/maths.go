package maths

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

type Number interface {
	int | int8 | int16 | int32 | int64 
}

func Gcd[T Number](a, b T) T {
	if b == 0 {
		return a
	}

	return Gcd(b, a % b)
}

func Lcm[T Number](a, b T) T {
	return (a * b) / Gcd(a, b)
}

func LcmRange[T Number]( vals ... T) T {
	if len(vals) < 2 {
		panic("LCM RANGE TO SMALL")
	}
	var res T = Lcm(vals[0], vals[1])
	
	if len(vals) == 2 {
		return res
	}

	for _, val := range vals[2:] {
		res = Lcm(res, val)
	}

	return res
}

func Transpose[T any](matrix [][]T) [][]T {
	transposed := make([][]T, len(matrix[0]))

	for i := range matrix[0] {
		transposed[i] = make([]T, len(matrix))
	}

	for i, row := range matrix {

		for j := range row {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func GenerateHash(values ...interface{}) uint64 {
	hash := fnv.New64a()

	for _, val := range values {
		hash.Write([]byte(fmt.Sprintf("%v", val)))
	}

	return hash.Sum64()
}

func ToInt(s string) int {
	res, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}
	return res
}