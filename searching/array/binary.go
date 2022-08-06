package array

import (
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](x T, a []T) (result int) {
	assertion.Require(IsSorted(a, 0, len(a)), "a is sorted")
	defer func() {
		assertion.Ensure(!IsIn(x, a, 0, len(a)) && result == -1 || 0 <= result && result < len(a) && x == a[result], "result is OK")
	}()

	low := 0
	high := len(a)

	loopInv := func(low, high int) bool {
		assertion.Invariant(0 <= low && low <= high && high <= len(a), "low and high are within bound")
		assertion.Invariant(low == 0 || a[low-1] < x, "x is larger than any element from a[0, low)")
		assertion.Invariant(high == len(a) || a[high] > x, "x is smaller than any element from a[high,len(a))")
		return true
	}
	for loopInv(low, high) && low < high {
		mid := low + (high-low)/2
		assertion.Check(low <= mid && mid < high, "mid is within [low, high)")

		if a[mid] == x {
			return mid
		} else if a[mid] < x {
			low = mid + 1
		} else { // a[mid] > x
			high = mid
		}
	}

	return -1
}
