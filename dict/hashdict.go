package dict

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/assertion"
)

type entry[K any, V any] struct {
	key   K
	value V
}

func (e *entry[K, V]) Key() K {
	return e.key
}

func (e *entry[K, V]) Value() V {
	return e.value
}

type HashFn[K comparable] func(K) int

type HashDict[K comparable, V comparable] struct {
	size     int
	capacity int
	table    []linked.List[entry[K, V]]
	hashFn   HashFn[K]
	maxLoad  int
}

// data structure invariant
func (h *HashDict[K, V]) isHashDict() bool {
	return h != nil && 0 <= h.size && 0 < h.capacity &&
		len(h.table) == h.capacity && 0 < h.maxLoad
}

func NewHashDict[K comparable, V comparable](capacity int, hashFn HashFn[K], maxLoad int) (result HashDict[K, V]) {
	assertion.Require(0 < capacity, "capacity is positive")
	assertion.Require(hashFn != nil, "hash function is not nil")
	defer func() {
		assertion.Ensure(result.isHashDict(), "hash dict invariant holds")
	}()

	table := make([]linked.List[entry[K, V]], capacity)
	return HashDict[K, V]{
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

func (h *HashDict[K, V]) indexOfKey(key K) (result int) {
	assertion.Require(h.isHashDict(), "hash dict invariant holds")
	defer func() {
		assertion.Ensure(0 <= result && result < h.capacity, "result is within bound")
	}()

	return abs(h.hashFn(key) % h.capacity)
}

func (h *HashDict[K, V]) Get(key K) (V, bool) {
	assertion.Require(h.isHashDict(), "hash dict invariant holds")

	index := h.indexOfKey(key)
	l := &h.table[index]

	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Data.Key() == key {
			return curr.Data.Value(), true
		}
	}

	return *new(V), false
}

func (h *HashDict[K, V]) Put(key K, value V) {
	assertion.Require(h.isHashDict(), "hash dict invariant holds")
	defer func() {
		assertion.Ensure(h.isHashDict(), "hash dict invariant holds")
		v, ok := h.Get(key)
		assertion.Ensure(ok && v == value, "Get() returns value for key")
	}()

	index := h.indexOfKey(key)

	l := &h.table[index]

	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Data.Key() == key {
			curr.Data.value = value
			return
		}
	}

	newHead := linked.NewNode(entry[K, V]{key: key, value: value})
	newHead.Next = l.Head
	l.Head = newHead
	h.size++

	if h.size >= h.capacity*h.maxLoad {
		h.resize(h.capacity * 2)
	}
}

func (h *HashDict[K, V]) Delete(key K) {
	assertion.Require(h.isHashDict(), "hash dict invariant holds")
	defer func() {
		assertion.Ensure(h.isHashDict(), "hash dict invariant holds")
		_, ok := h.Get(key)
		assertion.Ensure(!ok, "Get() returns no value for key")
	}()

	index := h.indexOfKey(key)

	l := &h.table[index]
	if l.Head == nil {
		return
	}

	isDeleted := false
	for curr := &l.Head; *curr != nil; curr = &(*curr).Next {
		if (*curr).Data.Key() == key {
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

func (h *HashDict[K, V]) Size() (result int) {
	assertion.Require(h.isHashDict(), "hash dict invariant holds")
	defer func() {
		assertion.Ensure(result >= 0, "result is non-negative")
	}()

	return h.size
}

func (h *HashDict[K, V]) resize(newSize int) {
	assertion.Require(h.isHashDict(), "hash dict invariant holds")
	defer func(oldSize int) {
		assertion.Ensure(h.isHashDict(), "hash dict invariant holds")
		assertion.Ensure(oldSize == h.size, "resize does not affect count of elements")
		assertion.Ensuref(newSize == h.capacity, "resize changes capacity, newSize=%v,capacity=%v", newSize, h.capacity)
	}(h.size)

	if newSize == h.capacity {
		return
	}

	oldTable := h.table
	h.table = make([]linked.List[entry[K, V]], newSize)
	h.capacity = newSize
	h.size = 0

	for _, l := range oldTable {
		for curr := l.Head; curr != nil; curr = curr.Next {
			h.Put(curr.Data.Key(), curr.Data.Value())
		}
	}
}
