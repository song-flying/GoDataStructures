package list

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/assertion"
)

func LinearSearch[T comparable](x T, l linked.List[T]) *linked.Node[T] {
	assertion.Require(l.IsList(), "list invariant holds")

	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Data == x {
			return curr
		}
	}

	return nil
}
