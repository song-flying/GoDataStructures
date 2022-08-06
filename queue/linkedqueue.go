package queue

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/assertion"
)

type LinkedQueue[T any] struct {
	front *linked.Node[T]
	back  *linked.Node[T]
}

// IsLinkedQueue data structure invariant
func (q *LinkedQueue[T]) IsLinkedQueue() bool {
	return q != nil && !linked.HasCycle(q.front) && linked.IsSegment(q.front, q.back)
}

func NewLinkedQueue[T any]() (result LinkedQueue[T]) {
	defer func() {
		assertion.Ensure(result.IsLinkedQueue(), "queue invariant holds")
		assertion.Ensure(result.IsEmpty(), "new queue is empty")
	}()

	dummy := linked.NewDummyNode[T]()
	return LinkedQueue[T]{
		front: dummy,
		back:  dummy,
	}
}

func (q *LinkedQueue[T]) IsEmpty() bool {
	assertion.Require(q.IsLinkedQueue(), "queue invariant holds")

	return q.front == q.back
}

func (q *LinkedQueue[T]) Enqueue(x T) {
	assertion.Require(q.IsLinkedQueue(), "queue invariant holds")
	defer func() {
		assertion.Ensure(q.IsLinkedQueue(), "queue invariant holds")
	}()

	q.back.Data = x
	q.back.Next = linked.NewDummyNode[T]()
	q.back = q.back.Next
}

func (q *LinkedQueue[T]) Dequeue() (result T) {
	assertion.Require(q.IsLinkedQueue(), "queue invariant holds")
	assertion.Require(!q.IsEmpty(), "queue is not empty")
	defer func() {
		assertion.Ensure(q.IsLinkedQueue(), "queue invariant holds")
	}()

	result = q.front.Data
	q.front = q.front.Next

	return
}

func (q *LinkedQueue[T]) Head() T {
	assertion.Require(q.IsLinkedQueue(), "queue invariant holds")
	assertion.Require(!q.IsEmpty(), "queue is not empty")

	return q.front.Data
}
