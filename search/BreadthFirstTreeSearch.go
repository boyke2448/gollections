package search

import (
	"log"

	"github.com/boyke2448/gollections/linkedList"
	"github.com/boyke2448/gollections/queue"
)

func BreadthFirstTreeSearch[T comparable](tree *linkedList.MultiLinkedList[T], target T) *linkedList.MultiLinkedList[T] {
	q := queue.New[linkedList.MultiLinkedList[T]]()
	q.Enqueue(tree)

	for q.Count() > 0 {
		item, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}

		if item.Value == target {
			return item
		}

		for _, child := range item.Children {
			if child != nil {
				q.Enqueue(child)
			}
		}
	}
	return nil
}
