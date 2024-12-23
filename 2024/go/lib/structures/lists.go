package structures

import (
	"advent/lib/maths"
)

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

type Set[T comparable] map[T]struct{}

type ExplorationSet = Set[maths.Position]

func (set Set[T]) IsEmpty() bool {
	return len(set) == 0
}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (set Set[T]) Contains(key T) bool {
	_, found := set[key]
	return found
}

func (set Set[T]) ContainsAll(vals ...T) bool {
	for _, v := range vals {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}
func (set Set[T]) Items() []T {
	items := make([]T, 0)
	for k := range set {
		items = append(items, k)
	}
	return items
}
func (set *Set[T]) Insert(val T) {
	(*set)[val] = struct{}{}
}

func (set *Set[T]) Remove(search T) {

	if !set.Contains(search) {
		return
	}
	delete(*set, search)
}


func (set *Set[T]) Intersect(other Set[T]) Set[T] {
	intersect := Set[T]{}

	for key, _ := range other {
		if set.Contains(key) {
			intersect.Insert(key)
		}
	}

	return intersect
}

func (set *Set[T]) Union(other Set[T]) Set[T] {
	union := Set[T]{}

	for key, _ := range *set {
		union.Insert(key)
	}
	for key, _ := range other {
		union.Insert(key)
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