package array

import (
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"golang.org/x/exp/constraints"
)

// IsIn specification function
func IsIn[T comparable](x T, a []T, low, high int) bool {
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

// IsSorted specification function
func IsSorted[T constraints.Ordered](a []T, low, high int) bool {
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

type CompareFn[T comparable] func(x, y T) int

// IsSortedBy specification function
func IsSortedBy[T comparable](a []T, comp CompareFn[T], low, high int) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")

	loopInv := func(i int) bool {
		assertion.Invariant(low <= i, "i is within lower bound")
		return true
	}
	for i := low; loopInv(i) && i < high-1; i++ {
		assertion.Check(i < high-1, "i is within upper bound")
		if comp(a[i], a[i+1]) > 0 {
			return false
		}
	}

	return true
}
