package heap

import (
	"github.com/song-flying/GoDataStructures/array"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/order"
)

type Heap[E comparable] struct {
	next     int
	capacity int
	data     []E
	comp     order.CompareFn[E]
}

func (h *Heap[E]) isHeapSafe() bool {
	return h != nil && 1 < h.capacity && h.capacity == len(h.data) &&
		1 <= h.next && h.next <= h.capacity && h.comp != nil
}

func (h *Heap[E]) isHeapOrdered() bool {
	contract.Require(h.isHeapSafe(), "heap is safe to access")
	for i := 2; i < h.next; i++ {
		if !h.aboveOK(i) {
			return false
		}
	}

	return true
}

func (h *Heap[E]) Contains(element E) bool {
	contract.Require(h.IsHeap(), "heap invariant holds")
	return array.Contains(element, h.data, 1, h.next)
}

// IsHeap data structure invariant
func (h *Heap[E]) IsHeap() bool {
	return h.isHeapSafe() && h.isHeapOrdered()
}

func up(childIndex int) int {
	return childIndex / 2
}

func left(parentIndex int) int {
	return parentIndex * 2
}

func right(parentIndex int) int {
	return parentIndex*2 + 1
}

func (h *Heap[E]) orderOK(i, j int) bool {
	contract.Require(1 <= i && i < h.next, "i is within bound")
	contract.Require(1 <= j && j < h.next, "j is within bound")
	return h.comp(h.data[i], h.data[j]) <= 0
}

func (h *Heap[E]) aboveOK(i int) bool {
	contract.Require(1 <= i && i < h.next, "i is within bound")
	if i == 1 {
		return true
	}

	return h.orderOK(up(i), i)
}

func (h *Heap[E]) belowOK(i int) bool {
	contract.Require(1 <= i && i < h.next, "i is within bound")

	return (left(i) >= h.next || h.orderOK(i, left(i))) && (right(i) >= h.next || h.orderOK(i, right(i)))
}

func (h *Heap[E]) swapUp(cIndex int) {
	contract.Require(h.isHeapSafe(), "heap is safe to access")
	contract.Require(1 < cIndex && cIndex < h.next, "cIndex is within bound")
	contract.Require(!h.aboveOK(cIndex), "parent is not less than child")
	defer func() {
		contract.Ensure(h.aboveOK(cIndex), "parent is less than child")
	}()

	pIndex := up(cIndex)
	tmp := h.data[pIndex]
	h.data[pIndex] = h.data[cIndex]
	h.data[cIndex] = tmp
}

func NewHeap[E comparable](userCapacity int, lessFn order.CompareFn[E]) (result Heap[E]) {
	contract.Require(0 < userCapacity, "userCapacity is positive")
	contract.Require(lessFn != nil, "priority function is not nil")
	defer func() {
		contract.Ensure(result.IsHeap(), "heap invariant holds")
	}()

	return Heap[E]{
		next:     1,
		capacity: userCapacity + 1,
		data:     make([]E, userCapacity+1),
		comp:     lessFn,
	}
}

func (h *Heap[E]) IsEmpty() bool {
	contract.Require(h.isHeapSafe(), "heap is safe to access")
	return h.next == 1
}

func (h *Heap[E]) IsFull() bool {
	contract.Require(h.isHeapSafe(), "heap is safe to access")
	return h.next == h.capacity
}

func (h *Heap[E]) isHeapExceptUp(exception int) bool {
	contract.Require(h.isHeapSafe(), "heap is safe to access")
	contract.Require(1 <= exception && exception < h.next, "exception index is within bound")

	for i := 2; i < h.next; i++ {
		if i == exception {
			continue
		}
		if !h.aboveOK(i) {
			return false
		}
	}

	return true
}

func (h *Heap[E]) checkAboveAndBelow(i int) bool {
	if i == 1 {
		return true
	}
	if left(i) >= h.next {
		return true
	}
	if right(i) >= h.next {
		return h.orderOK(up(i), left(i))
	}
	return h.orderOK(up(i), left(i)) && h.orderOK(up(i), right(i))
}

func (h *Heap[E]) Add(element E) {
	contract.Require(h.IsHeap(), "heap invariant holds")
	contract.Require(!h.IsFull(), "heap is not full")
	defer func() {
		contract.Ensure(h.IsHeap(), "heap invariant holds")
		contract.Ensure(h.Contains(element), "heap contains element")
	}()

	h.data[h.next] = element
	h.next++

	loopInv := func(i int) bool {
		contract.Invariant(1 <= i && i < h.next, "i is within bound")
		contract.Invariant(h.isHeapExceptUp(i), "heap invariant holds except for node i")
		contract.Invariant(h.checkAboveAndBelow(i), "i's parent has no lower priority than i's children")
		return true
	}
	for i := h.next - 1; loopInv(i) && i > 1 && !h.aboveOK(i); i = i / 2 {
		h.swapUp(i)
	}
}

func (h *Heap[E]) isHeapExceptDown(exception int) bool {
	contract.Require(h.isHeapSafe(), "heap is safe to access")
	contract.Require(1 <= exception && exception < h.next, "exception index is within bound")

	for i := 2; i < h.next; i++ {
		if i == exception {
			continue
		}
		if !h.belowOK(i) {
			return false
		}
	}

	return true
}

func (h *Heap[E]) selectChildToSwap(i int) int {
	if right(i) >= h.next {
		return left(i)
	}

	if h.orderOK(left(i), right(i)) {
		return left(i)
	}

	return right(i)
}

func (h *Heap[E]) Delete() (result E) {
	contract.Require(h.IsHeap(), "heap invariant holds")
	contract.Require(!h.IsEmpty(), "heap is not empty")
	defer func() {
		contract.Ensure(h.IsHeap(), "heap invariant holds")
		contract.Ensure(!h.Contains(result), "heap does not contain element")
	}()

	result = h.data[1]
	h.next--

	if h.next == 1 {
		return
	}

	last := h.data[h.next]
	h.data[1] = last

	loopInv := func(i int) bool {
		contract.Invariant(1 <= i && i < h.next, "i is within bound")
		contract.Invariant(h.isHeapExceptDown(i), "heap invariant holds except for node i")
		contract.Invariant(h.checkAboveAndBelow(i), "i's parent has no lower priority than i's children")
		return true
	}
	for i := 1; loopInv(i) && left(i) < h.next && !h.belowOK(i); {
		child := h.selectChildToSwap(i)
		h.swapUp(child)
		i = child
	}

	return
}

func (h Heap[E]) Size() int {
	return h.next - 1
}
