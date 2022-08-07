package sorting

import (
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"github.com/song-flying/GoDataStructures/pkg/order"
)

func SelectionSort[T comparable](a []T, comp order.CompareFn[T]) {
	SelectionSortRange(a, 0, len(a), comp)
}

func SelectionSortRange[T comparable](a []T, low, high int, comp order.CompareFn[T]) {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")
	defer func() {
		assertion.Ensure(isSorted(a, low, high, comp), "a is sorted")
	}()

	loopInv := func(i int) bool {
		assertion.Invariant(low <= i && i <= high, "i is within bound")
		assertion.Invariant(isSorted(a, low, i, comp), "a[low,i) is sorted")
		assertion.Invariant(rangeLessOrEqual(a, low, i, a, i, high, comp), "a[low,i) <= a[i,high)")
		return true
	}
	for i := low; loopInv(i) && i < high; i++ {
		minIndex := findMin(a, i, high, comp)
		swap(a, minIndex, i)
	}
}

func findMin[T comparable](a []T, low, high int, comp order.CompareFn[T]) (result int) {
	assertion.Require(0 <= low && low < high && high <= len(a), "low and high are within bound")
	defer func() {
		assertion.Ensure(low <= result && result < high, "result is within bound")
		assertion.Ensure(elementLessThanOrEqual(a[result], a, low, high, comp), "result is the index of smallest element")
	}()

	minIndex := low
	loopInv := func(i int) bool {
		assertion.Invariant(low+1 <= i && i <= high, "i is within bound")
		assertion.Invariant(low <= minIndex && minIndex < high, "minIndex is within bound")
		assertion.Invariant(elementLessThanOrEqual(a[minIndex], a, low, i, comp), "a[minIndex] is min value for a[low,i)")
		return true
	}
	for i := low + 1; loopInv(i) && i < high; i++ {
		if comp(a[i], a[minIndex]) < 0 {
			minIndex = i
		}
	}

	return minIndex
}
