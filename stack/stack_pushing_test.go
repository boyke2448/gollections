package stack_test

import (
	"testing"

	"github.com/boyke2448/gollections/stack"
)

func Test_Add1intElementTostack(t *testing.T) {
	test_addElementToStack(t, 1)
}

func Test_Add1StringElementToStack_ShouldContain1StackItem(t *testing.T) {
	test_addElementToStack(t, "hello")
}

func Test_Add1StructElementToStack_ShouldContain1StackItem(t *testing.T) {
	type s struct{}
	test_addElementToStack(t, s{})
}

func Test_AddStructElementsToStack_ShouldContainMultipleStackItems(t *testing.T) {
	type s struct{}

	item := make([]s, 10)
	stck := stack.New[s]()
	for _, v := range item {
		stck.Push(v)
	}
	if stck.Count() != 10 {
		t.Errorf("Expected stack length to be 10 but found %d", stck.Count())
	}
}

func test_addElementToStack[T interface{}](t *testing.T, value T) {
	s := stack.New[T]()
	s.Push(value)
	if s.Count() != 1 {
		t.Errorf("Expected stack length to be 1 but found %d", s.Count())
	}
}
