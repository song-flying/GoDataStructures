package array

import (
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"golang.org/x/exp/constraints"
)

func LinearSortedSearch[T constraints.Ordered](x T, a []T) (result int) {
	assertion.Require(IsSorted(a, 0, len(a)), "a is sorted")
	defer func() {
		assertion.Ensure(0 <= result && result < len(a) && x == a[result] || !IsIn(x, a, 0, len(a)) && result == -1, "result OK")
	}()

	loopInv := func(i int) bool {
		assertion.Invariant(0 <= i && i <= len(a), "i is within bound")
		assertion.Invariant(i == 0 || x > a[i-1], "x is larger than any element from a[0, i)")
		return true
	}
	for i := 0; loopInv(i) && i < len(a); i++ {
		if x == a[i] {
			return i
		} else if x < a[i] {
			return -1
		}
	}

	return -1
}

func LinearSearch[T comparable](x T, a []T, comp CompareFn[T]) (result int) {
	for i := 0; i < len(a); i++ {
		if x == a[i] {
			return i
		}
	}

	return -1
}
