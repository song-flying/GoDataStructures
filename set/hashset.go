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
	maxLoad  int
}

// data structure invariant
func (h *HashSet[E]) isHashSet() bool {
	return h != nil && 0 <= h.size && 0 < h.capacity &&
		len(h.table) == h.capacity && 0 < h.maxLoad
}

func NewHashSet[E comparable](capacity int, hashFn HashFn[E], maxLoad int) (result HashSet[E]) {
	assertion.Require(0 < capacity, "capacity is positive")
	assertion.Require(0 < maxLoad, "maxLoad is positive")
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
		maxLoad:  maxLoad,
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

func (h *HashSet[E]) indexOfElement(key E) (result int) {
	assertion.Require(h.isHashSet(), "hash set invariant holds")
	defer func() {
		assertion.Ensure(0 <= result && result < h.capacity, "result is within bound")
	}()

	return abs(h.hashFn(key) % h.capacity)
}

func (h *HashSet[E]) Contains(x E) bool {
	assertion.Require(h.isHashSet(), "hash set invariant holds")

	index := h.indexOfElement(x)
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
		assertion.Ensure(h.Contains(x), "hash set contains element x")
	}()

	index := h.indexOfElement(x)
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

	if h.size >= h.capacity*h.maxLoad {
		h.resize(h.capacity * 2)
	}
}

func (h *HashSet[E]) Remove(x E) {
	assertion.Require(h.isHashSet(), "hash set invariant holds")
	defer func() {
		assertion.Ensure(h.isHashSet(), "hash set invariant holds")
		assertion.Ensure(!h.Contains(x), "hash set does not contain element x")
	}()

	index := h.indexOfElement(x)
	l := &h.table[index]
	if l.Head == nil {
		return
	}

	isDeleted := false
	for curr := &l.Head; *curr != nil; curr = &(*curr).Next {
		if (*curr).Data == x {
			target := *curr
			*curr = target.Next
			target.Next = nil
			isDeleted = true
			break
		}
		if (*curr).Next == nil {
			break
		}
	}

	if isDeleted {
		h.size--
		if 4*h.size <= h.capacity*h.maxLoad {
			h.resize((h.capacity + 1) / 2)
		}
	}
}

func (h *HashSet[E]) Size() (result int) {
	assertion.Require(h.isHashSet(), "hash set invariant holds")
	defer func() {
		assertion.Ensure(result >= 0, "result is non-negative")
	}()

	return h.size
}

func (h *HashSet[E]) resize(newCapacity int) {
	assertion.Require(h.isHashSet(), "hash set invariant holds")
	defer func(oldSize int) {
		assertion.Ensure(h.isHashSet(), "hash set invariant holds")
		assertion.Ensure(oldSize == h.size, "resize does not change count of entries")
		assertion.Ensure(newCapacity == h.capacity, "resize changes capacity")
	}(h.size)

	if newCapacity == h.capacity {
		return
	}

	oldTable := h.table
	h.table = make([]linked.List[E], newCapacity)
	h.capacity = newCapacity
	h.size = 0

	for _, l := range oldTable {
		for curr := l.Head; curr != nil; curr = curr.Next {
			h.Add(curr.Data)
		}
	}
}
