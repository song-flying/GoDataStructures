package linked

import "github.com/song-flying/GoDataStructures/pkg/assertion"

type Node[T any] struct {
	Data T
	Next *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		Data: data,
		Next: nil,
	}
}

func NewDummyNode[T any]() *Node[T] {
	return &Node[T]{}
}

func Nil[T any]() *Node[T] {
	return nil
}

type List[T any] struct {
	Head *Node[T]
}

// IsList data structure invariant
func (l *List[T]) IsList() bool {
	return !HasCycle(l.Head)
}

func NewEmptyList[T any]() (result List[T]) {
	assertion.Ensure(result.IsList(), "list invariant holds")

	return List[T]{}
}

func NewList[T any](head *Node[T]) (result List[T]) {
	assertion.Ensure(result.IsList(), "list invariant holds")

	return List[T]{
		Head: head,
	}
}

func (l *List[T]) IsEmpty() bool {
	assertion.Require(l.IsList(), "list invariant holds")

	return l.Head == nil
}

func (l *List[T]) Length() (result int) {
	assertion.Require(l.IsList(), "list invariant holds")
	assertion.Ensure(0 <= result, "result is non-negative")

	if l.Head == nil {
		return 0
	}

	return LengthOfSegment(l.Head, nil)
}

func (l *List[T]) Ith(i int) T {
	assertion.Require(l.IsList(), "list invariant holds")
	assertion.Require(0 <= i && i < l.Length(), "i is within bound")

	return IthSegment(l.Head, i)
}
