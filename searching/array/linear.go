package array

import (
	"github.com/song-flying/GoDataStructures/array"
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"github.com/song-flying/GoDataStructures/pkg/order"
)

func LinearSortedSearch[T comparable](x T, a []T, comp order.CompareFn[T]) (result int) {
	assertion.Require(array.IsRangeSorted(a, 0, len(a), comp), "a is sorted")
	defer func() {
		assertion.Ensure(0 <= result && result < len(a) && x == a[result] || !array.IsIn(x, a, 0, len(a)) && result == -1, "result OK")
	}()

	loopInv := func(i int) bool {
		assertion.Invariant(0 <= i && i <= len(a), "i is within bound")
		assertion.Invariant(i == 0 || comp(x, a[i-1]) > 0, "x is larger than any element from a[0, i)")
		return true
	}
	for i := 0; loopInv(i) && i < len(a); i++ {
		if x == a[i] {
			return i
		} else if comp(x, a[i]) < 0 {
			return -1
		}
	}

	return -1
}
