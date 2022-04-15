package stack_test

import (
	"testing"

	"github.com/boyke2448/gollections/stack"
)

func Test_PopEmptyStack_ShouldReturnNilAndStackEmptyError(t *testing.T) {
	type s struct{}

	stck := stack.Stack[s]{}
	item, err := stck.Pop()
	if err == nil {
		t.Error("Expected error but did not receive one.")
	}

	_, ok := err.(stack.StackEmptyError)
	if !ok {
		t.Errorf("Expected error fo tyep stack.StackEmptyError but got %T", err)
	}

	if item != nil {
		t.Errorf("Expected no item to be returned. But got %v", item)
	}
}

func Test_Pop1ElementOnStack_ShouldReturnItem(t *testing.T) {
	type s struct{}

	stck := stack.Stack[s]{}
	stck.Push(&s{})
	length := stck.Count()

	if length != 1 {
		t.Errorf("Should contain 1 element but found %d elements.", length)
	}

	item, err := stck.Pop()

	if err != nil {
		t.Error(err)
	}

	if item == nil {
		t.Error("Expected an popped item.")
	}

	length = stck.Count()

	if length != 0 {
		t.Errorf("Should contain no elements but found %d elements", length)
	}

}
