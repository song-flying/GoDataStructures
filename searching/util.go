package searching

import "github.com/song-flying/GoDataStructures/pkg/assertion"

// specification function
func isIn(x int, a []int, low, high int) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bounds")

	loopInv := func(i int) bool {
		assertion.Invariant(low <= i && i <= high, "i is within bound")
		return true
	}
	for i := low; loopInv(i) && i < high; i++ {
		if x == a[i] {
			return true
		}
	}

	return false
}

// specification function
func isSorted(a []int, low, high int) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")

	loopInv := func(i int) bool {
		assertion.Invariant(low <= i, "i is within lower bound")
		return true
	}
	for i := low; loopInv(i) && i < high-1; i++ {
		assertion.Check(i < high-1, "i is within upper bound")
		if a[i] > a[i+1] {
			return false
		}
	}

	return true
}
