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