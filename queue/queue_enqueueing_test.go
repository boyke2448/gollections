package queue_test

import (
	"testing"

	"github.com/boyke2448/gollections/queue"
)

func Test_Add1intElementToQueue_ShouldContain1QueueItem(t *testing.T) {
	test_addElementToQueue(t, 1)
}

func Test_Add1StringElementToQueue_ShouldContain1QueueItem(t *testing.T) {
	test_addElementToQueue(t, "hello")
}

func Test_Add1StructElementToQueue_ShouldContain1QueueItem(t *testing.T) {
	type s struct{}
	test_addElementToQueue(t, s{})
}

func Test_AddStructElementsToQueue_ShouldContainMultipleQueueItems(t *testing.T) {
	type s struct{}

	item := make([]s, 10)
	q := queue.New[s]()
	for _, v := range item {
		q.Enqueue(v)
	}
	if q.Count() != 10 {
		t.Errorf("Expected queue length to be 10 but found %d", q.Count())
	}
}

func test_addElementToQueue[T interface{}](t *testing.T, value T) {
	q := queue.New[T]()
	q.Enqueue(value)
	if q.Count() != 1 {
		t.Errorf("Expected queue length to be 1 but found %d", q.Count())
	}
}
