package queue_test

import (
	"testing"

	"github.com/boyke2448/gollections/queue"
)

func Test_DequeueEmptyQueue_ShouldReturnNilAndQueueEmptyError(t *testing.T) {
	type s struct{}

	q := queue.Queue[s]{}
	item, err := q.Dequeue()
	if err == nil {
		t.Error("Expected error but did not receive one.")
	}
	_, ok := err.(queue.QueueEmptyError)
	if !ok {
		t.Errorf("Expected error of type goqueue.QueueEmptyError but got %T", err)
	}
	if item != nil {
		t.Errorf("Expected no item to be returned. But got %v", item)
	}
}

func Test_Dequeue1ElementInQueue_ShouldReturnItem(t *testing.T) {
	type s struct{}

	q := queue.Queue[s]{}
	q.Enqueue(&s{})
	size := q.Count()

	if size != 1 {
		t.Errorf("Should contain 1 element but found %d items.", size)
	}

	item, err := q.Dequeue()

	if err != nil {
		t.Error(err)
	}

	if item == nil {
		t.Error("Expected an dequeued item.")
	}

	size = q.Count()

	if size != 0 {
		t.Errorf("Should contain no elements but found %d elements", size)
	}
}
