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

func NewEmptyList[T any]() List[T] {
	return List[T]{}
}

func NewList[T any](head *Node[T]) List[T] {
	return List[T]{
		Head: head,
	}
}

func (l *List[T]) IsEmpty() bool {
	return l.Head == nil
}

func (l *List[T]) HasCycle() bool {
	return HasCycle(l.Head)
}

func (l *List[T]) Length() int {
	assertion.Require(!l.HasCycle(), "l has no cycle")

	if l.Head == nil {
		return 0
	}

	return LengthOfSegment(l.Head, nil)
}

func (l *List[T]) Ith(i int) T {
	assertion.Require(0 <= i && i < l.Length(), "i is within bound")

	return IthSegment(l.Head, i)
}
