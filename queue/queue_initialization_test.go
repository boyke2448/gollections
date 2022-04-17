package queue_test

import (
	"testing"

	"github.com/boyke2448/gollections/queue"
)

func Test_InitialGetLength_ShouldBe0(t *testing.T) {
	q := queue.New[int]()
	if q.Count() != 0 {
		t.Errorf("Expected queue length to be 0 but found %d", q.Count())
	}
}
