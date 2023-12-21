package structures_test

import (
	"advent/lib/structures"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := structures.NewStack[int]()
	if stack == nil {
		t.Errorf("Stack Instantiation Failed")
	}
}
func TestSize(t *testing.T) {
	stack := structures.NewStack[int]()
	if stack.Size() != 0 {
		t.Errorf("Stack Length Failed")
	}
	stack.PushFront(12)
	if stack.Size() != 1 {
		t.Errorf("Stack Length Failed After Push")
		
	}
}
func TestPop(t *testing.T) {
	stack := structures.NewStack[int]()

	stack.PushFront(12)

	res := stack.Pop()

	if stack.Size() != 0 {
		t.Errorf("Stack Pop Did not reduce size")
	}

	if res != 12 {
		t.Errorf("Stack Pop Returned Wrong Value")
	}

}