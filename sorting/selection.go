package sorting

import (
	"github.com/song-flying/GoDataStructures/pkg/assertion"
)

func SelectionSort(a []int) {
	SelectionSortRange(a, 0, len(a))
}

func SelectionSortRange(a []int, low, high int) {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")
	defer func() {
		assertion.Ensure(isSorted(a, low, high), "a is sorted")
	}()

	loopInv := func(i int) bool {
		assertion.Invariant(low <= i && i <= high, "i is within bound")
		assertion.Invariant(isSorted(a, low, i), "a[low,i) is sorted")
		assertion.Invariant(rangeLessOrEqual(a, low, i, a, i, high), "a[low,i) <= a[i,high)")
		return true
	}
	for i := low; loopInv(i) && i < high; i++ {
		minIndex := findMin(a, i, high)
		swap(a, minIndex, i)
	}
}

func findMin(a []int, low, high int) (result int) {
	assertion.Require(0 <= low && low < high && high <= len(a), "low and high are within bound")
	defer func() {
		assertion.Ensure(low <= result && result < high, "result is within bound")
		assertion.Ensure(elementLessOrEqual(a[result], a, low, high), "result is the index of smallest element")
	}()

	minIndex := low
	loopInv := func(i int) bool {
		assertion.Invariant(low+1 <= i && i <= high, "i is within bound")
		assertion.Invariant(low <= minIndex && minIndex < high, "minIndex is within bound")
		assertion.Invariant(elementLessOrEqual(a[minIndex], a, low, i), "a[minIndex] is min value for a[low,i)")
		return true
	}
	for i := low + 1; loopInv(i) && i < high; i++ {
		if a[i] < a[minIndex] {
			minIndex = i
		}
	}

	return minIndex
}
