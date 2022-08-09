package linked

import (
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"golang.org/x/exp/constraints"
)

func HasCycle[T comparable](l *Node[T]) bool {
	if l == nil {
		return false
	}

	slow := l
	fast := l.Next

	loopInv := func(i int) bool {
		contract.Invariant(fast == nil || slow != nil, "fast /= nil => slow /= nil")
		contract.Invariant(fast == nil || isReachableWith(l, slow, i) && isReachableWith(l, fast, 2*i+1), "speeds of fast and slow are OK")
		return true
	}
	for i := 0; loopInv(i) && fast != nil && slow != fast; i++ {
		slow = next(slow, 1)
		contract.Assert(slow != nil, "slow node is not nil")

		fast = next(fast, 2)
	}

	if fast == nil {
		return false
	} else {
		contract.Assert(fast != nil && slow != nil, "fast and slow are not nil")
		contract.Assert(fast == slow, "fast and slow points to same node")
		return true
	}
}

func next[T comparable](start *Node[T], n int) (result *Node[T]) {
	contract.Require(n >= 0, "n is non-negative")
	defer func() {
		contract.Ensure(result == nil || isReachableWith(start, result, n), "result is reachable from start with n steps")
	}()

	if start == nil {
		return nil
	} else if n == 0 {
		return start
	} else {
		return next(start.Next, n-1)
	}
}

// specification function
func isReachableWith[T comparable](src, dst *Node[T], hops int) bool {
	contract.Require(hops >= 0, "hops is non-negative")

	curr := src
	for i := 0; i < hops; i++ {
		if curr == nil {
			return false
		}
		curr = curr.Next
	}

	return curr == dst
}

func IsSegment[T comparable](start, end *Node[T]) bool {
	contract.Require(!HasCycle(start), "start node leads to no cycle")

	for curr := start; curr != nil; curr = curr.Next {
		if curr == end {
			return true
		}
	}

	return start != nil && end == nil
}

func LengthOfSegment[T comparable](start, end *Node[T]) int {
	contract.Require(IsSegment(start, end), "start and end forms a segment")

	count := 0
	for curr := start; curr != end; curr = curr.Next {
		count++
	}

	return count
}

func IsInSegment[T comparable](x T, start, end *Node[T]) (result bool) {
	contract.Require(IsSegment(start, end), "start and end forms a segment")

	for curr := start; curr != end; curr = curr.Next {
		if x == curr.Data {
			return true
		}
	}

	return false
}

func IsSegmentSorted[T constraints.Ordered](start, end *Node[T]) bool {
	contract.Require(IsSegment(start, end), "start and ends forms a segment")

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

func IthSegment[T comparable](start *Node[T], i int) T {
	contract.Require(0 <= i && i < LengthOfSegment(start, nil), "i is within bound")

	curr := start
	for count := 0; count < i; count++ {
		contract.Assert(curr != nil, "current node is not nil")
		curr = curr.Next
	}

	return curr.Data
}
