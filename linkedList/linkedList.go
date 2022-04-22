package linkedList

type LinkedList[T interface{}] struct {
	Value *T
	Next  *LinkedList[T]
}
