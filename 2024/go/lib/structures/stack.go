package structures

import "slices"

type Stack[T comparable] struct {

	_backed []T;
}

func NewStack[T comparable]() *Stack[T] {

	return &Stack[T]{
		_backed: make([]T, 0),
	}
}

func (stack *Stack[T]) Pop() T {
	if stack.Size() == 0 {
		panic("Cant Pop Empty Stack")
	}
	val := stack._backed[0]
	stack._backed = stack._backed[1:]
	return val
}
func (stack *Stack[T]) PushFront(val T) {
	stack._backed = append([]T{ val, }, stack._backed...)
}

func (stack *Stack[T]) PushFrontAll(vals []T) {
	stack._backed = append(vals, stack._backed...)
}

func (stack *Stack[T]) Size() int {
	return len(stack._backed)
}

func (stack *Stack[T]) Push(val T) {
	stack._backed = append(stack._backed, val)
}
func (stack *Stack[T]) PushEval(val T, lt func(a, b T) bool) {
	insert := -1
	for i, v := range stack._backed {

		if lt(val, v) {
			insert = i
			break;
		}
	}
	if insert == -1 {
		stack._backed = append(stack._backed, val)
	} else {
		stack._backed = append(stack._backed[:insert], append([]T{val}, stack._backed[insert:]...)...)
	}
}

func (stack *Stack[T]) PushAll(vals []T) {
	stack._backed = append(stack._backed, vals...)
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack._backed) == 0
}

func (stack *Stack[T]) SortFunc(sorter func(a, b T) int) {
	slices.SortFunc(stack._backed, sorter)
}