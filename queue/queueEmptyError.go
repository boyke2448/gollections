package queue

type QueueEmptyError struct {
}

func (QueueEmptyError) Error() string {
	return "Queue is empty"
}
