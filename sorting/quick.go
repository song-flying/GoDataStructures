package sorting

import "github.com/song-flying/GoDataStructures/pkg/assertion"

func QuickSort(a []int) {
	QuickSortRange(a, 0, len(a))
}

func QuickSortRange(a []int, low, high int) {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")
	defer func() {
		assertion.Ensure(isSorted(a, low, high), "result is sorted")
	}()

	if high-low <= 1 {
		return
	}

	pIndex := low + (high-low)/2
	mid := partition(a, low, pIndex, high)

	QuickSortRange(a, low, mid)

	QuickSortRange(a, mid+1, high)
}

func partition(a []int, low, pIndex, high int) (result int) {
	assertion.Require(0 <= low && low < high && high <= len(a), "low and high are within bound")
	assertion.Require(low <= pIndex && pIndex < high, "pivot index is within bound")
	defer func() {
		assertion.Ensure(low <= result && result < high, "result is within bound")
		assertion.Ensure(elementGreaterThanOrEqual(a[result], a, low, result), "a[result] >= any element from a[low, result)")
		assertion.Ensure(elementLessThanOrEqual(a[result], a, result+1, high), "a[result] <= any element from a[result+1, high)")
	}()

	pivot := a[pIndex]
	swap(a, low, pIndex)

	left := low + 1
	right := high

	loopInv := func() bool {
		assertion.Invariant(low+1 <= left && left <= right && right <= high, "left and right are within bound")
		assertion.Invariant(elementGreaterThanOrEqual(pivot, a, low+1, left), "pivot >= any element from a[low, left)")
		assertion.Invariant(elementLessThanOrEqual(pivot, a, right, high), "pivot <= any element from a[right, high)")
		return true
	}
	for loopInv() && left < right {
		if a[left] <= pivot {
			left++
		} else { // a[left] > pivot
			swap(a, left, right-1)
			right--
		}
	}

	swap(a, low, left-1)

	return left - 1
}
