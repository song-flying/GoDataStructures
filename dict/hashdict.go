package dict

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/contract"
)

type HashFn[K comparable] func(K) int

type HashDict[K comparable, V comparable] struct {
	size     int
	capacity int
	table    []linked.List[entry[K, V]]
	hashFn   HashFn[K]
	maxLoad  int
}

func (h *HashDict[K, V]) listOK() bool {
	for _, l := range h.table {
		if !l.IsList() {
			return false
		}
	}

	return true
}

func (h *HashDict[K, V]) hashOK() bool {
	for i, l := range h.table {
		for curr := l.Head; curr != nil; curr = curr.Next {
			hashIndex := h.indexOfKey(curr.Data.Key)
			if i != hashIndex {
				return false
			}
		}
	}

	return true
}

func (h *HashDict[K, V]) sizeOK() bool {
	size := 0
	for _, l := range h.table {
		size += l.Length()
	}

	return h.size == size
}

// IsHashDict data structure invariant
func (h *HashDict[K, V]) IsHashDict() bool {
	return h != nil && 0 <= h.size && 0 < h.capacity && len(h.table) == h.capacity &&
		0 < h.maxLoad && h.listOK() && h.hashOK() && h.sizeOK()
}

func NewHashDict[K comparable, V comparable](capacity int, hashFn HashFn[K], maxLoad int) (result HashDict[K, V]) {
	contract.Require(0 < capacity, "capacity is positive")
	contract.Require(0 < maxLoad, "maxLoad is positive")
	contract.Require(hashFn != nil, "hash function is not nil")
	defer func() {
		contract.Ensure(result.IsHashDict(), "hash dict invariant holds")
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
		contract.Ensure(result >= 0, "result is non-negative")
	}()

	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func (h *HashDict[K, V]) indexOfKey(key K) (result int) {
	contract.Require(h.hashFn != nil, "hash function is not nil")
	contract.Require(h.capacity > 0, "capacity is positive")
	defer func() {
		contract.Ensure(0 <= result && result < h.capacity, "result is within bound")
	}()

	return abs(h.hashFn(key) % h.capacity)
}

func (h *HashDict[K, V]) Get(key K) (result V, found bool) {
	contract.Require(h.IsHashDict(), "hash dict invariant holds")

	index := h.indexOfKey(key)
	l := &h.table[index]

	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Data.Key == key {
			return curr.Data.Value, true
		}
	}

	return *new(V), false
}

func (h *HashDict[K, V]) Put(key K, value V) {
	contract.Require(h.IsHashDict(), "hash dict invariant holds")
	defer func() {
		contract.Ensure(h.IsHashDict(), "hash dict invariant holds")
		v, ok := h.Get(key)
		contract.Ensure(ok && v == value, "Get(key) returns value")
	}()

	index := h.indexOfKey(key)
	l := &h.table[index]
	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Data.Key == key {
			curr.Data.Value = value
			return
		}
	}

	newHead := linked.NewNode(entry[K, V]{Key: key, Value: value})
	newHead.Next = l.Head
	l.Head = &newHead
	h.size++

	if h.size >= h.capacity*h.maxLoad {
		h.resize(h.capacity * 2)
	}
}

func (h *HashDict[K, V]) Delete(key K) {
	contract.Require(h.IsHashDict(), "hash dict invariant holds")
	defer func() {
		contract.Ensure(h.IsHashDict(), "hash dict invariant holds")
		_, ok := h.Get(key)
		contract.Ensure(!ok, "Get() returns no value for key")
	}()

	index := h.indexOfKey(key)
	l := &h.table[index]
	if l.Head == nil {
		return
	}

	isDeleted := false
	for curr := &l.Head; *curr != nil; curr = &(*curr).Next {
		if (*curr).Data.Key == key {
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
	contract.Require(h.IsHashDict(), "hash dict invariant holds")
	defer func() {
		contract.Ensure(result >= 0, "result is non-negative")
	}()

	return h.size
}

func (h *HashDict[K, V]) resize(newCapacity int) {
	contract.Require(h.IsHashDict(), "hash dict invariant holds")
	defer func(oldSize int) {
		contract.Ensure(h.IsHashDict(), "hash dict invariant holds")
		contract.Ensure(oldSize == h.size, "resize does not change count of entries")
		contract.Ensure(newCapacity == h.capacity, "resize changes capacity")
	}(h.size)

	if newCapacity == h.capacity {
		return
	}

	oldTable := h.table
	h.table = make([]linked.List[entry[K, V]], newCapacity)
	h.capacity = newCapacity
	h.size = 0

	for _, l := range oldTable {
		for curr := l.Head; curr != nil; curr = curr.Next {
			h.Put(curr.Data.Key, curr.Data.Value)
		}
	}
}
