package linked

import (
	"github.com/song-flying/GoDataStructures/pkg/contract"
)

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

func NewNode[T comparable](data T) Node[T] {
	return Node[T]{
		Data: data,
		Next: nil,
	}
}

func NewDummyNode[T comparable]() Node[T] {
	return Node[T]{}
}

func Nil[T comparable]() *Node[T] {
	return nil
}

// IsList data structure invariant
func (n *Node[T]) IsList() bool {
	return !HasCycle(n)
}

type List[T comparable] struct {
	Head *Node[T]
}

// IsList data structure invariant
func (l *List[T]) IsList() bool {
	return l.Head.IsList()
}

func NewEmptyList[T comparable]() (result List[T]) {
	contract.Ensure(result.IsList(), "list invariant holds")

	return List[T]{}
}

func NewList[T comparable](head *Node[T]) (result List[T]) {
	contract.Ensure(result.IsList(), "list invariant holds")

	return List[T]{
		Head: head,
	}
}

func (l *List[T]) IsEmpty() bool {
	contract.Require(l.IsList(), "list invariant holds")

	return l.Head == nil
}

func (l *List[T]) Length() (result int) {
	contract.Require(l.IsList(), "list invariant holds")
	contract.Ensure(0 <= result, "result is non-negative")

	if l.Head == nil {
		return 0
	}

	return LengthOfSegment(l.Head, nil)
}

func (l *List[T]) Ith(i int) T {
	contract.Require(l.IsList(), "list invariant holds")
	contract.Require(0 <= i && i < l.Length(), "i is within bound")

	return IthSegment(l.Head, i)
}

func (l *List[T]) Add(element T) {
	contract.Require(l.IsList(), "list invariant holds")
	node := NewNode(element)
	node.Next = l.Head
	l.Head = &node
}

func (l *List[T]) Contains(element T) bool {
	return l.containsFrom(l.Head, element)
}

func (l *List[T]) containsFrom(start *Node[T], element T) bool {
	contract.Require(l.IsList(), "list invariant holds")
	for curr := start; curr != nil; curr = curr.Next {
		if curr.Data == element {
			return true
		}
	}

	return false
}

func (l *List[T]) ToArray() (result []T) {
	contract.Require(l.IsList(), "list invariant holds")

	for curr := l.Head; curr != nil; curr = curr.Next {
		result = append(result, curr.Data)
	}

	return
}

func (l *List[T]) IsDistinct() bool {
	return l.isDistinctFrom(l.Head)
}

func (l *List[T]) isDistinctFrom(start *Node[T]) bool {
	contract.Require(l.IsList(), "list invariant holds")

	if start == nil {
		return true
	}

	return !l.containsFrom(start.Next, start.Data) && l.isDistinctFrom(start.Next)
}
