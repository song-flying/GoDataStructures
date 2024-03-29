package array

import (
	"github.com/song-flying/GoDataStructures/array"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/order"
)

func BinarySearch[T comparable](x T, a []T, comp order.CompareFn[T]) (result int) {
	contract.Require(array.IsRangeSorted(a, 0, len(a), comp), "a is sorted")
	defer func() {
		contract.Ensure(!array.IsIn(x, a, 0, len(a)) && result == -1 || 0 <= result && result < len(a) && x == a[result], "result is OK")
	}()

	low := 0
	high := len(a)

	loopInv := func(low, high int) bool {
		contract.Invariant(0 <= low && low <= high && high <= len(a), "low and high are within bound")
		contract.Invariant(low == 0 || comp(a[low-1], x) < 0, "x is larger than any element from a[0, low)")
		contract.Invariant(high == len(a) || comp(a[high], x) > 0, "x is smaller than any element from a[high,len(a))")
		return true
	}
	for loopInv(low, high) && low < high {
		mid := low + (high-low)/2
		contract.Assert(low <= mid && mid < high, "mid is within [low, high)")

		if a[mid] == x {
			return mid
		} else if comp(a[mid], x) < 0 {
			low = mid + 1
		} else { // a[mid] > x
			high = mid
		}
	}

	return -1
}
