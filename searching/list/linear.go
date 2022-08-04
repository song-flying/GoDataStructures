package list

import "github.com/song-flying/GoDataStructures/linked"

func LinearSearch[T comparable](x T, l linked.List[T]) *linked.Node[T] {
	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Data == x {
			return curr
		}
	}

	return nil
}
