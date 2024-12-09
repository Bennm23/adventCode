package structures

type AnyMap[T comparable, R any] map[T]R

func (mp AnyMap[T, R]) ValueSet() []R {
	values := []R{}

	for _, val := range mp {
		values = append(values, val)
	}

	return values
}