package tree

import (
	"github.com/boyke2448/gollections/linkedList"
	"github.com/boyke2448/gollections/search"
)

type ITree[T comparable] interface {
	GetRoot() *linkedList.MultiLinkedList[T]
	AddNode(parentValue, value T)
}

type tree[T comparable] struct {
	root *linkedList.MultiLinkedList[T]
}

func New[T comparable](value T) ITree[T] {
	t := new(tree[T])
	t.root = &linkedList.MultiLinkedList[T]{Value: value}
	return t
}

func (t *tree[T]) GetRoot() *linkedList.MultiLinkedList[T] {
	return t.root
}

func (t *tree[T]) AddNode(parentValue, value T) {
	node := search.BreadthFirstTreeSearch(t.root, parentValue)
	if node.Children == nil {
		node.Children = make([]*linkedList.MultiLinkedList[T], 0)
	}

	node.Children = append(node.Children, &linkedList.MultiLinkedList[T]{Value: value})
}
