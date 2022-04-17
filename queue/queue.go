package queue

import "github.com/boyke2448/gollections/linkedList"

type IQueue[T interface{}] interface {
	Count() int
	Enqueue(value T)
	Dequeue() (*T, error)
}

// Queue[T interface{}] is a queue which will accept elements of type T
type queue[T interface{}] struct {
	count int
	front *linkedList.LinkedList[T]
	back  *linkedList.LinkedList[T]
}

func New[T interface{}]() IQueue[T] {
	return new(queue[T])
}

// Count gets the size of the queue
func (q queue[T]) Count() int {
	return q.count
}

// Enqueue add item to the back of the queue
func (q *queue[T]) Enqueue(value T) {
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
func (q *queue[T]) Dequeue() (*T, error) {
	if q.front == nil {
		return nil, QueueEmptyError{}
	}
	item := q.front.Value
	q.front = q.front.Next
	q.count--
	return &item, nil
}
