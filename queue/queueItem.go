package queue

type queueItem[T interface{}] struct {
	value *T
	next  *queueItem[T]
}
