package list

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/contract"
)

func LinearSearch[T comparable](x T, l linked.List[T]) *linked.Node[T] {
	contract.Require(l.IsList(), "list invariant holds")

	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Data == x {
			return curr
		}
	}

	return nil
}
