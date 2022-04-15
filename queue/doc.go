/*
Go Queue
This package implements a queueing mechanisme for Go with O(1) time complexisty and O(n) space complexity.
It uses a linked list as internal implementation to achieve this.
IMPORTANT: This package needs go 1.18 or greater because it uses generics.

Usage:
	q := goqueue.New[SomeType]()
	q.Enqueue(&SomeType{})
	q.Enqueue(&SomeType{})
	q.Enqueue(&SomeType{})
	var item, err := q.Dequeue()
	if err != nil {
		switch err.(type) {
		case QueueEmptyError:
			log.Fatal(err)
			os.Exit(1)
		default:
			log.Println(err)
	}
*/
package queue
