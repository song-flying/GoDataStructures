package linked

import (
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"golang.org/x/exp/constraints"
)

func HasCycle[T any](l List[T]) bool {
	if l == nil {
		return false
	}

	slow := l
	fast := l.Next

	loopInv := func(i int) bool {
		assertion.Invariant(fast == nil || slow != nil, "fast /= nil => slow /= nil")
		assertion.Invariant(fast == nil || isReachableWith(l, slow, i) && isReachableWith(l, fast, 2*i+1), "speeds of fast and slow are OK")
		return true
	}
	for i := 0; loopInv(i) && fast != nil && slow != fast; i++ {
		slow = Next(slow, 1)
		assertion.Check(slow != nil, "slow node is not nil")

		fast = Next(fast, 2)
	}

	if fast == nil {
		return false
	} else {
		assertion.Check(fast != nil && slow != nil, "fast and slow are not nil")
		assertion.Check(fast == slow, "fast and slow points to same node")
		return true
	}
}

func Next[T any](start *Node[T], n int) (result *Node[T]) {
	assertion.Require(n >= 0, "n is non-negative")
	defer func() {
		assertion.Ensure(result == nil || isReachableWith(start, result, n), "result is reachable from start with n steps")
	}()

	if start == nil {
		return nil
	} else if n == 0 {
		return start
	} else {
		return Next(start.Next, n-1)
	}
}

// specification function
func isReachableWith[T any](src, dst *Node[T], hops int) bool {
	assertion.Require(hops >= 0, "hops is non-negative")

	curr := src
	for i := 0; i < hops; i++ {
		if curr == nil {
			return false
		}
		curr = curr.Next
	}

	return curr == dst
}

func IsSegment[T any](start, end *Node[T]) bool {
	assertion.Require(!HasCycle(start), "start node leads to no cycle")

	for curr := start; curr != nil; curr = curr.Next {
		if curr == end {
			return true
		}
	}

	return start != nil && end == nil
}

func LengthOfSegment[T any](start, end List[T]) int {
	assertion.Require(IsSegment(start, end), "start and end forms a segment")

	count := 0
	for curr := start; curr != end; curr = curr.Next {
		count++
	}

	return count
}

func Length[T any](l List[T]) int {
	assertion.Require(!HasCycle(l), "l has no cycle")

	length := 0
	for curr := l; curr != nil; curr = curr.Next {
		length++
	}

	return length
}

func Ith[T any](l List[T], i int) T {
	assertion.Require(0 <= i && i < Length(l), "i is within bound")

	curr := l
	for count := 0; count < i; count++ {
		assertion.Check(curr != nil, "current node is not nil")
		curr = curr.Next
	}

	return curr.Data
}

func isInSegment[T comparable](x T, start, end List[T]) (result bool) {
	assertion.Require(IsSegment(start, end), "start and end forms a segment")

	for curr := start; curr != end; curr = curr.Next {
		if x == curr.Data {
			return true
		}
	}

	return false
}

func isSegmentSorted[T constraints.Ordered](start, end List[T]) bool {
	assertion.Require(IsSegment(start, end), "start and ends forms a segment")

	if start == end {
		return true
	}

	for prev, curr := start, start.Next; curr != end; prev, curr = prev.Next, curr.Next {
		if prev.Data >= curr.Data {
			return false
		}
	}

	return true
}

func BinarySearch[T constraints.Ordered](x T, l List[T]) (result int) {
	assertion.Require(!HasCycle(l), "l has no cycle")

	if l == nil {
		return -1
	}

	return BinarySearchSegment(x, l, nil)
}

func BinarySearchSegment[T constraints.Ordered](x T, start, end List[T]) (result int) {
	assertion.Require(isSegmentSorted(start, end), "segment [start,end) is sorted")
	defer func() {
		assertion.Ensure(
			result == -1 && !isInSegment(x, start, end) ||
				0 <= result && result < LengthOfSegment(start, end) && Ith(start, result) == x,
			"result is OK")
	}()

	low := 0
	high := LengthOfSegment(start, end)

	for low < high {
		mid := low + (high-low)/2
		assertion.Check(low <= mid && mid < high, "mid is within [low, high)")
		midVal := Ith(start, mid)

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
