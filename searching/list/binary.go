package list

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](x T, l *linked.List[T]) (result int) {
	assertion.Require(!l.HasCycle(), "l has no cycle")

	if l.IsEmpty() {
		return -1
	}

	return BinarySearchSegment(x, l.Head, nil)
}

func BinarySearchSegment[T constraints.Ordered](x T, start, end *linked.Node[T]) (result int) {
	assertion.Require(linked.IsSegmentSorted(start, end), "segment [start,end) is sorted")
	defer func() {
		assertion.Ensure(
			result == -1 && !linked.IsInSegment(x, start, end) ||
				0 <= result && result < linked.LengthOfSegment(start, end) && linked.IthSegment(start, result) == x,
			"result is OK")
	}()

	low := 0
	high := linked.LengthOfSegment(start, end)

	for low < high {
		mid := low + (high-low)/2
		assertion.Check(low <= mid && mid < high, "mid is within [low, high)")
		midVal := linked.IthSegment(start, mid)

		if midVal == x {
			return mid
		} else if midVal < x {
			low = mid + 1
		} else { // midVal > x
			high = mid
		}
	}

	return -1
}
