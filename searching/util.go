package searching

import "github.com/song-flying/GoDataStructures/pkg/assertion"

// specification function
func isIn(x int, a []int, low, high int) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bounds")

	for i := low; i < high; i++ {
		assertion.Invariant(low <= i && i <= high, "i is within bound")
		if x == a[i] {
			return true
		}
	}

	return false
}

// specification function
func isSorted(a []int, low, high int) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")

	for i := low; i < high-1; i++ {
		assertion.Invariant(low <= i, "i is within lower bound")
		assertion.Check(i < high-1, "i is within upper bound")
		if a[i] > a[i+1] {
			return false
		}
	}

	return true
}
