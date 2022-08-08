package sorting

import (
	"github.com/song-flying/GoDataStructures/heap"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/order"
)

func HeapSort[T comparable](a []T, comp order.CompareFn[T]) {
	if len(a) == 0 {
		return
	}

	h := heap.NewHeap(len(a), comp)
	for _, e := range a {
		contract.Assert(!h.IsFull(), "heap is not full")
		h.Add(e)
	}

	for i := 0; i < len(a); i++ {
		contract.Assert(!h.IsEmpty(), "heap is not empty")
		a[i] = h.Delete()
	}

	return
}
