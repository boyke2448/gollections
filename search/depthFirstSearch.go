package search

import (
	"fmt"
	"log"

	"github.com/boyke2448/gollections/linkedList"
	"github.com/boyke2448/gollections/stack"
)

func DepthFirstTreeSearch[T comparable](binaryTree *linkedList.DoubleLinkedList[T], target T, actionBefore func(T), actionInner func(T), actionAfter func(T)) {
	s := stack.New[linkedList.DoubleLinkedList[T]]()
	s.Push(binaryTree)

	for s.Count() > 0 {
		item, err := s.Pop()
		if err != nil {
			log.Fatal(err)
		}

		if item.Value == target {
			if actionBefore != nil {
				actionBefore(item.Value)
			} else {
				fmt.Printf("Found  it: %v\n", item.Value)
			}
		}

		if item.Left != nil {
			s.Push(item.Left)
		}

		if item.Value == target {
			if actionInner != nil {
				actionInner(item.Value)
			} else {
				fmt.Printf("Found  it: %v\n", item.Value)
			}
		}

		if item.Right != nil {
			s.Push(item.Right)
		}

		if item.Value == target {
			if actionAfter != nil {
				actionAfter(item.Value)
			} else {
				fmt.Printf("Found  it: %v\n", item.Value)
			}
		}
	}
}
