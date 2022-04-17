package linkedList

type DoubleLinkedList[T interface{}] struct {
	Value T
	Left  *DoubleLinkedList[T]
	Right *DoubleLinkedList[T]
}
