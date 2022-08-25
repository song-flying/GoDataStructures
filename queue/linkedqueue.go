package queue

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/contract"
)

type LinkedQueue[T comparable] struct {
	front *linked.Node[T]
	back  *linked.Node[T]
}

// IsLinkedQueue data structure invariant
func (q *LinkedQueue[T]) IsLinkedQueue() bool {
	return q != nil && !linked.HasCycle(q.front) && linked.IsSegment(q.front, q.back)
}

func NewLinkedQueue[T comparable]() (result *LinkedQueue[T]) {
	defer func() {
		contract.Ensure(result.IsLinkedQueue(), "queue invariant holds")
		contract.Ensure(result.IsEmpty(), "new queue is empty")
	}()

	dummy := linked.NewDummyNode[T]()
	return &LinkedQueue[T]{
		front: &dummy,
		back:  &dummy,
	}
}

func (q *LinkedQueue[T]) IsEmpty() bool {
	contract.Require(q.IsLinkedQueue(), "queue invariant holds")

	return q.front == q.back
}

func (q *LinkedQueue[T]) Enqueue(x T) {
	contract.Require(q.IsLinkedQueue(), "queue invariant holds")
	defer func() {
		contract.Ensure(q.IsLinkedQueue(), "queue invariant holds")
	}()

	dummy := linked.NewDummyNode[T]()
	q.back.Data = x
	q.back.Next = &dummy
	q.back = q.back.Next
}

func (q *LinkedQueue[T]) Dequeue() (result T) {
	contract.Require(q.IsLinkedQueue(), "queue invariant holds")
	contract.Require(!q.IsEmpty(), "queue is not empty")
	defer func() {
		contract.Ensure(q.IsLinkedQueue(), "queue invariant holds")
	}()

	result = q.front.Data
	q.front = q.front.Next

	return
}

func (q *LinkedQueue[T]) Head() T {
	contract.Require(q.IsLinkedQueue(), "queue invariant holds")
	contract.Require(!q.IsEmpty(), "queue is not empty")

	return q.front.Data
}
