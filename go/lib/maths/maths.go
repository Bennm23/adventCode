package maths

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