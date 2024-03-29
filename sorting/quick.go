package sorting

import (
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/order"
)

func QuickSort[T comparable](a []T, comp order.CompareFn[T]) {
	QuickSortRange(a, 0, len(a), comp)
}

func QuickSortRange[T comparable](a []T, low, high int, comp order.CompareFn[T]) {
	contract.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")
	defer func() {
		contract.Ensure(isSorted(a, low, high, comp), "result is sorted")
	}()

	if high-low <= 1 {
		return
	}

	pIndex := low + (high-low)/2
	mid := partition(a, low, pIndex, high, comp)

	QuickSortRange(a, low, mid, comp)
	contract.Assert(isSorted(a, low, mid, comp), "a[low,mid) is sorted")

	QuickSortRange(a, mid+1, high, comp)
	contract.Assert(isSorted(a, mid+1, high, comp), "a[mid+1,high) is sorted")
}

func partition[T comparable](a []T, low, pIndex, high int, comp order.CompareFn[T]) (result int) {
	contract.Require(0 <= low && low < high && high <= len(a), "low and high are within bound")
	contract.Require(low <= pIndex && pIndex < high, "pivot index is within bound")
	defer func() {
		contract.Ensure(low <= result && result < high, "result is within bound")
		contract.Ensure(elementGreaterThanOrEqual(a[result], a, low, result, comp), "a[result] >= any element from a[low, result)")
		contract.Ensure(elementLessThanOrEqual(a[result], a, result+1, high, comp), "a[result] <= any element from a[result+1, high)")
	}()

	pivot := a[pIndex]
	swap(a, low, pIndex)

	left := low + 1
	right := high

	loopInv := func() bool {
		contract.Invariant(low+1 <= left && left <= right && right <= high, "left and right are within bound")
		contract.Invariant(elementGreaterThanOrEqual(pivot, a, low+1, left, comp), "pivot >= any element from a[low, left)")
		contract.Invariant(elementLessThanOrEqual(pivot, a, right, high, comp), "pivot <= any element from a[right, high)")
		return true
	}
	for loopInv() && left < right {
		if comp(a[left], pivot) <= 0 {
			left++
		} else { // a[left] > pivot
			swap(a, left, right-1)
			right--
		}
	}

	swap(a, low, left-1)

	return left - 1
}
