package sorting

import "github.com/song-flying/GoDataStructures/pkg/assertion"

func MergeSort(a []int) {
	MergeSortRange(a, 0, len(a))
}

func MergeSortRange(a []int, low, high int) {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within range")
	defer func() {
		assertion.Ensure(isSorted(a, low, high), "result is sorted")
	}()

	if high-low <= 1 {
		return
	}

	mid := low + (high-low)/2

	MergeSortRange(a, low, mid)

	MergeSortRange(a, mid, high)

	merge(a, low, mid, high)
}

func merge(a []int, low int, mid int, high int) {
	assertion.Require(0 <= low && low <= mid && mid < high && high <= len(a), "low, mid and high are within bound")
	assertion.Require(isSorted(a, low, mid), "a[low, mid) is sorted")
	assertion.Require(isSorted(a, mid, high), "a[mid, high) is sorted")
	defer func() {
		assertion.Ensure(isSorted(a, low, high), "a[low, high) is sorted")
	}()

	b := make([]int, high-low)
	curr := 0
	i, j := low, mid
	loopInv := func() bool {
		assertion.Invariant(low <= i && i <= mid, "i is within bound")
		assertion.Invariant(mid <= j && j <= high, "j is within bound")
		assertion.Invariant(isSorted(b, 0, curr), "b[0, curr) is sorted")
		assertion.Invariant(isSorted(a, i, mid), "a[i, mid) is sorted")
		assertion.Invariant(i == mid || rangeLessOrEqual(b, 0, curr, a, i, mid), "b[0, curr) <= a[i, mid)")
		assertion.Invariant(isSorted(a, j, high), "a[j, high) is sorted")
		assertion.Invariant(j == high || rangeLessOrEqual(b, 0, curr, a, j, high), "b[0, curr) <= a[j, high)")
		assertion.Invariant(curr == (i-low)+(j-mid), "curr = (i-low) + (j-mid)")
		return true
	}
	for ; loopInv() && i < mid && j < high; curr++ {
		if a[i] <= a[j] {
			b[curr] = a[i]
			i++
		} else {
			b[curr] = a[j]
			j++
		}
	}

	if i == mid {
		for k := j; k < high; k++ {
			b[curr] = a[k]
			curr++
		}
	} else if j == high {
		for k := i; k < mid; k++ {
			b[curr] = a[k]
			curr++
		}
	}

	for i := low; i < high; i++ {
		a[i] = b[i-low]
	}
}
