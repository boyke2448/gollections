/*
Go Stack
This package implements a stacking mechanisme for Go with O(1) time complexisty and O(n) space complexity.
It uses a linked list as internal implementation to achieve this.
IMPORTANT: This package needs go 1.18 or greater because it uses generics.

Usage:
	q := stack.Stack[SomeType]()
	q.Push(&SomeType{})
	q.Push(&SomeType{})
	q.Push(&SomeType{})
	var item, err := q.Pop()
	if err != nil {
		switch err.(type) {
		case stack.StackEmptyError:
			log.Fatal(err)
			os.Exit(1)
		default:
			log.Println(err)
	}
*/
package stack
