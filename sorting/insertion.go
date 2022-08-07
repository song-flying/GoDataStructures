package sorting

import (
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"github.com/song-flying/GoDataStructures/pkg/order"
)

func InsertionSort[T comparable](a []T, comp order.CompareFn[T]) {
	InsertionSortRange(a, 0, len(a), comp)
}

func InsertionSortRange[T comparable](a []T, low, high int, comp order.CompareFn[T]) {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")
	defer func() {
		assertion.Ensure(isSorted(a, low, high, comp), "a[low,high) is sorted")
	}()

	loopInv := func(i int) bool {
		assertion.Invariant(low <= i && i <= high, "i is within bound")
		assertion.Invariant(isSorted(a, low, i, comp), "a[low, i) is sorted")
		return true
	}
	for i := low; loopInv(i) && i < high; i++ {
		insert(a, low, i, comp)
	}
}

func insert[T comparable](a []T, low int, curr int, comp order.CompareFn[T]) {
	assertion.Require(0 <= low && low <= curr && curr < len(a), "low and i are within bound")
	assertion.Require(isSorted(a, low, curr, comp), "a[low,curr) is sorted")
	defer func() {
		assertion.Ensure(isSorted(a, low, curr+1, comp), "a[low,curr] is sorted")
	}()

	loopInv := func(i int) bool {
		assertion.Invariant(low <= i && i <= curr, "i is within bound")
		assertion.Invariant(isSorted(a, low, i, comp), "a[low,i) is sorted")
		assertion.Invariant(isSorted(a, i, curr, comp), "a[i,curr] is sorted")
		assertion.Invariant(rangeLessOrEqual(a, low, i, a, i+1, curr+1, comp), "a[low, i) <= a(i, curr]")
		return true
	}
	for i := curr; loopInv(i) && i > low; i-- {
		if comp(a[i-1], a[i]) > 0 {
			swap(a, i-1, i)
		} else {
			break
		}
	}
}
