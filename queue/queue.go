package queue

import "github.com/boyke2448/gollections/linkedList"

// Queue[T interface{}] is a queue which will accept elements of type T
type Queue[T interface{}] struct {
	count int
	front *linkedList.LinkedList[T]
	back  *linkedList.LinkedList[T]
}

// Count gets the size of the queue
func (q Queue[T]) Count() int {
	return q.count
}

// Enqueue add item to the back of the queue
func (q *Queue[T]) Enqueue(value *T) {
	queueItem := &linkedList.LinkedList[T]{Value: value}
	if q.front == nil {
		q.front = queueItem
	} else {
		if q.front.Next == nil {
			q.back = queueItem
			q.front.Next = q.back
		} else {
			q.back.Next = queueItem
			q.back = queueItem
		}
	}
	q.count++
}

// Dequeue remove item from the front of the queue.
// if queue is empty nil is returned with a QueueEmptyError
// if an element is available this is returned and removed from the queue.
func (q *Queue[T]) Dequeue() (*T, error) {
	if q.front == nil {
		return nil, QueueEmptyError{}
	}
	item := q.front
	q.front = q.front.Next
	q.count--
	return item.Value, nil
}
