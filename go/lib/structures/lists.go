package structures

type List[T comparable] []T
type ListTwoD[T comparable] [][]T

func NewTwoDList[T comparable]() List[T] {

	return make(List[T], 0)
}

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
		for i:=0; i<len(vals);i++ {

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

func (set *Set[T]) AddAll(vals ... T) {

	for _, v := range vals {
		set.Add(v)
	}

}
func (set *Set[T]) Add(val T) {

	if !set.Contains(val) {
		*set = append(*set, val)
	}

}