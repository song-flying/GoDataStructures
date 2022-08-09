package graph

import (
	"github.com/song-flying/GoDataStructures/graph"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/hash"
	"github.com/song-flying/GoDataStructures/set"
	"github.com/song-flying/GoDataStructures/stack"
)

func DepthFirstSearch[V comparable](g *graph.UndirectedGraph[V], v V, w V) bool {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(v) && g.Contains(w), "g contains v and w")

	marked := set.NewHashSet[V](g.Size(), hash.Universal[V], 1)
	s := stack.NewLinkedStack[V]()
	marked.Add(v)
	s.Push(v)
	// v is on stack => v is marked

	for !s.IsEmpty() {
		x := s.Pop()
		if x == w {
			return true
		}

		neighbors := g.GetNeighbors(x)
		for curr := neighbors.Head; curr != nil; curr = curr.Next {
			y := curr.Data
			if !marked.Contains(y) {
				marked.Add(y)
				s.Push(y)
			}
		}
	}

	return false
}
