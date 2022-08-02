package queue

type Queue[T any] interface {
	IsEmpty() bool
	Enqueue(x T)
	Dequeue() T
	Head() T
}
