package linkedList

type MultiLinkedList[T comparable] struct {
	Value    T
	Children []*MultiLinkedList[T]
}
