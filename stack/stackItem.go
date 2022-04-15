package stack

type stackItem[T interface{}] struct {
	value *T
	next  *stackItem[T]
}
