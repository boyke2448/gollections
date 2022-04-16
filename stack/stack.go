package stack

import "github.com/boyke2448/gollections/linkedList"

// Stack[T interface{}] is a stack which will accept elements of type T
type Stack[T interface{}] struct {
	count int
	top   *linkedList.LinkedList[T]
}

// Count gets the size of the stack
func (s Stack[T]) Count() int {
	return s.count
}

// Push adds item to the top of the stack
func (s *Stack[T]) Push(value *T) {
	item := &linkedList.LinkedList[T]{Value: value}
	if s.top == nil {
		s.top = item
	} else {
		currentTop := s.top
		item.Next = currentTop
		s.top = item
	}
	s.count++
}

// Pop remove item from the top of the stack.
// if stack is empty nil is returned with a StackEmptyError
// if an element is available this is returned and removed from the stack.
func (s *Stack[T]) Pop() (*T, error) {
	if s.top == nil {
		return nil, StackEmptyError{}
	}
	item := s.top.Value
	s.top = s.top.Next
	s.count--
	return item, nil
}
