package stack

// Stack[T interface{}] is a stack which will accept elements of type T
type Stack[T interface{}] struct {
	count int
	top   *stackItem[T]
}

// Count gets the size of the stack
func (s Stack[T]) Count() int {
	return s.count
}

// Push adds item to the top of the stack
func (s *Stack[T]) Push(value *T) {
	item := &stackItem[T]{value: value}
	if s.top == nil {
		s.top = item
	} else {
		currentTop := s.top
		item.next = currentTop
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
	item := s.top.value
	s.top = s.top.next
	s.count--
	return item, nil
}
