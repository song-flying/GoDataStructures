package set

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/assertion"
)

type HashFn[E comparable] func(E) int

type HashSet[E comparable] struct {
	size     int
	capacity int
	table    []linked.List[E]
	hashFn   HashFn[E]
}

// data structure invariant
func (h *HashSet[E]) isHashSet() bool {
	return h != nil && 0 <= h.size && h.size < h.capacity && 0 <= h.capacity && len(h.table) == h.capacity
}

func NewHashSet[E comparable](capacity int, hashFn HashFn[E]) (result HashSet[E]) {
	assertion.Require(0 < capacity, "capacity is positive")
	assertion.Require(hashFn != nil, "hash function is not nil")
	defer func() {
		assertion.Ensure(result.isHashSet(), "hash set invariant holds")
	}()

	table := make([]linked.List[E], capacity)
	return HashSet[E]{
		size:     0,
		capacity: capacity,
		table:    table,
		hashFn:   hashFn,
	}
}

func abs(x int) (result int) {
	defer func() {
		assertion.Ensure(result >= 0, "result is non-negative")
	}()

	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func (h *HashSet[E]) indexOfKey(key E) (result int) {
	assertion.Require(h.isHashSet(), "hash set invariant holds")
	defer func() {
		assertion.Ensure(0 <= result && result < h.capacity, "result is within bound")
	}()

	return abs(h.hashFn(key)) % h.capacity
}

func (h *HashSet[E]) Contains(x E) bool {
	assertion.Require(h.isHashSet(), "hash set invariant holds")

	index := h.indexOfKey(x)
	l := &h.table[index]
	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Data == x {
			return true
		}
	}

	return false
}

func (h *HashSet[E]) Add(x E) {
	assertion.Require(h.isHashSet(), "hash set invariant holds")
	defer func() {
		assertion.Ensure(h.isHashSet(), "hash set invariant holds")
		assertion.Ensure(h.Contains(x), "hash set contains x after")
	}()

	index := h.indexOfKey(x)
	l := &h.table[index]
	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Data == x {
			return
		}
	}

	newHead := linked.NewNode(x)
	newHead.Next = l.Head
	l.Head = newHead
	h.size++
}
