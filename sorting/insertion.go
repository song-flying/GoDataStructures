package sorting

import "github.com/song-flying/GoDataStructures/pkg/assertion"

func InsertionSort(a []int) {
	InsertionSortRange(a, 0, len(a))
}

func InsertionSortRange(a []int, low, high int) {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")
	defer func() {
		assertion.Ensure(isSorted(a, low, high), "a[low,high) is sorted")
	}()

	loopInv := func(i int) bool {
		assertion.Invariant(low <= i && i <= high, "i is within bound")
		assertion.Invariant(isSorted(a, low, i), "a[low, i) is sorted")
		return true
	}
	for i := low; loopInv(i) && i < high; i++ {
		insert(a, low, i)
	}
}

func insert(a []int, low int, curr int) {
	assertion.Require(0 <= low && low <= curr && curr < len(a), "low and i are within bound")
	assertion.Require(isSorted(a, low, curr), "a[low,curr) is sorted")
	defer func() {
		assertion.Ensure(isSorted(a, low, curr+1), "a[low,curr] is sorted")
	}()

	loopInv := func(i int) bool {
		assertion.Invariant(low <= i && i <= curr, "i is within bound")
		assertion.Invariant(isSorted(a, low, i), "a[low,i) is sorted")
		assertion.Invariant(isSorted(a, i, curr), "a[i,curr] is sorted")
		assertion.Invariant(rangeLessOrEqual(a, low, i, a, i+1, curr+1), "a[low, i) <= a(i, curr]")
		return true
	}
	for i := curr; loopInv(i) && i > low; i-- {
		if a[i-1] > a[i] {
			swap(a, i-1, i)
		} else {
			break
		}
	}
}
