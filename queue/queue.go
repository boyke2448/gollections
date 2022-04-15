package queue

// Queue[T interface{}] is a queue which will accept elements of type
type Queue[T interface{}] struct {
	count int
	front *queueItem[T]
	back  *queueItem[T]
}

// GetLength gets the size of the queue
func (q Queue[T]) Count() int {
	return q.count
}

// Enqueue add item to the back of the queue
func (q *Queue[T]) Enqueue(value *T) {
	queueItem := &queueItem[T]{value: value}
	if q.front == nil {
		q.front = queueItem
	} else {
		if q.front.next == nil {
			q.back = queueItem
			q.front.next = q.back
		} else {
			q.back.next = queueItem
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
	q.front = q.front.next
	q.count--
	return item.value, nil
}
