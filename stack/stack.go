package stack

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/assertion"
)

type Header[T any] struct {
	top    linked.List[T]
	bottom linked.List[T]
}

type LinkedStack[T any] *Header[T]

// data structure invariant
func (s LinkedStack[T]) isLinkedStack() bool {
	return s != nil && !linked.HasCycle(s.top) && linked.IsSegment(s.top, s.bottom)
}

func NewLinkedStack[T any]() (result LinkedStack[T]) {
	defer func() {
		assertion.Ensure(result.isLinkedStack(), "stack invariant holds")
		assertion.Ensure(result.IsEmpty(), "new stack is empty")
	}()

	dummy := linked.NewDummyNode[T]()
	return &Header[T]{
		top:    dummy,
		bottom: dummy,
	}
}

func (s LinkedStack[T]) IsEmpty() bool {
	return s.top == s.bottom
}

func (s LinkedStack[T]) Push(x T) {
	assertion.Require(s.isLinkedStack(), "stack invariant holds")
	defer func() {
		assertion.Ensure(s.isLinkedStack(), "stack invariant holds")
	}()

	l := linked.NewNode(x)
	l.Next = s.top
	s.top = l
}

func (s LinkedStack[T]) Pop() (result T) {
	assertion.Require(s.isLinkedStack(), "stack invariant holds")
	assertion.Require(!s.IsEmpty(), "stack is not empty")
	defer func() {
		assertion.Ensure(s.isLinkedStack(), "stack invariant holds")
	}()

	result = s.top.Data
	s.top = s.top.Next

	return
}
