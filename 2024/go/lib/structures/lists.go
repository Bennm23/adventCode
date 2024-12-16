package structures

type List[T comparable] []T
type ListTwoD[T comparable] [][]T

func (list List[T]) Contains(val T) bool {
	for _, a := range list {
		if a == val {
			return true
		}
	}
	return false
}

func (list *List[T]) Add(val T) {

	*list = append(*list, val)
}

func (list ListTwoD[T]) ContainsRow(vals []T) bool {

	for _, row := range list {
		allMatch := true
		for i := 0; i < len(vals); i++ {

			if vals[i] != row[i] {
				allMatch = false
				break
			}
		}
		if allMatch {
			return true
		}
	}
	return false
}

type Set[T comparable] []T

func (set Set[T]) Contains(val T) bool {
	for _, a := range set {
		if a == val {
			return true
		}
	}
	return false
}

func (set Set[T]) ContainsAll(vals ...T) bool {
	for _, v := range vals {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

func (set *Set[T]) AddAll(vals ...T) {

	for _, v := range vals {
		set.Add(v)
	}

}
func (set *Set[T]) Add(val T) {

	if !set.Contains(val) {
		*set = append(*set, val)
	}

}

func (set *Set[T]) Remove(search T) {

	if !set.Contains(search) {
		return
	}

	i := -1

	for ix, val := range *set {

		if search == val {
			i = ix
			break
		}
	}

	*set = append((*set)[:i], (*set)[i+1:]...)
}


func (set *Set[T]) Intersect(other Set[T]) Set[T] {
	intersect := Set[T]{}

	for _, val := range other {
		if set.Contains(val) {
			intersect.Add(val)
		}
	}

	return intersect
}

func (set *Set[T]) Union(other Set[T]) Set[T] {
	union := Set[T]{}

	for _, val := range *set {
		union.Add(val)
	}
	for _, val := range other {
		union.Add(val)
	}

	return union
}


func Create3DArray[T any](dimensions []int) [][][]T {
	if len(dimensions) != 3 {
		panic("Can't Create 3 Dimension Array Without 3 dimensions")
	}

	arr := make([][][]T, dimensions[0])

	for i := range arr {
		arr[i] = make([][]T, dimensions[1])
		for j := range arr[i] {
			arr[i][j] = make([]T, dimensions[2])
		}
	}

	return arr
}

func CountMatches[T comparable](list []T, val T) int {
	counter := 0

	for _, v := range list {
		if v == val {
			counter += 1
		}
	}
	return counter
}
func IndexOf[T comparable](list []T, val T) int {
	for vix, v := range list {

		if v == val {
			return vix
		}
	}
	return -1
}