package stack_test

import (
	"testing"

	"github.com/boyke2448/gollections/stack"
)

func Test_IntitialGetLength_ShouldBe0(t *testing.T) {
	s := stack.New[int]()
	if s.Count() != 0 {
		t.Errorf("Exptected stack length to be 0 but found %d", s.Count())
	}
}
