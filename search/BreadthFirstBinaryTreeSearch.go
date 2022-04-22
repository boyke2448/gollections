package search

import (
	"log"

	"github.com/boyke2448/gollections/linkedList"
	"github.com/boyke2448/gollections/queue"
)

func BreadthFirstBinaryTreeSearch[T comparable](binaryTree *linkedList.DoubleLinkedList[T], target T) *linkedList.DoubleLinkedList[T] {
	q := queue.New[linkedList.DoubleLinkedList[T]]()
	q.Enqueue(binaryTree)

	for q.Count() > 0 {
		item, err := q.Dequeue()

		if err != nil {
			log.Fatal(err)
		}

		if item.Value == target {
			return item
		}

		if item.Left != nil {
			q.Enqueue(item.Left)
		}

		if item.Right != nil {
			q.Enqueue(item.Right)
		}
	}

	return nil
}
