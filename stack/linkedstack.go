package stack

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/contract"
)

type LinkedStack[T any] struct {
	top    *linked.Node[T]
	bottom *linked.Node[T]
}

// IsLinkedStack data structure invariant
func (s *LinkedStack[T]) IsLinkedStack() bool {
	return s != nil && !linked.HasCycle(s.top) && linked.IsSegment(s.top, s.bottom)
}

func NewLinkedStack[T any]() (result LinkedStack[T]) {
	defer func() {
		contract.Ensure(result.IsLinkedStack(), "stack invariant holds")
		contract.Ensure(result.IsEmpty(), "new stack is empty")
	}()

	dummy := linked.NewDummyNode[T]()
	return LinkedStack[T]{
		top:    &dummy,
		bottom: &dummy,
	}
}

func (s *LinkedStack[T]) IsEmpty() bool {
	return s.top == s.bottom
}

func (s *LinkedStack[T]) Push(x T) {
	contract.Require(s.IsLinkedStack(), "stack invariant holds")
	defer func() {
		contract.Ensure(s.IsLinkedStack(), "stack invariant holds")
	}()

	l := linked.NewNode(x)
	l.Next = s.top
	s.top = &l
}

func (s *LinkedStack[T]) Pop() (result T) {
	contract.Require(s.IsLinkedStack(), "stack invariant holds")
	contract.Require(!s.IsEmpty(), "stack is not empty")
	defer func() {
		contract.Ensure(s.IsLinkedStack(), "stack invariant holds")
	}()

	result = s.top.Data
	s.top = s.top.Next

	return
}

func (s *LinkedStack[T]) Peek() T {
	contract.Require(s.IsLinkedStack(), "stack invariant holds")
	contract.Require(!s.IsEmpty(), "stack is not empty")

	return s.top.Data
}
