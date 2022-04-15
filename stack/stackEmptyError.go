package stack

type StackEmptyError struct {
}

func (StackEmptyError) Error() string {
	return "Stack is empty"
}
