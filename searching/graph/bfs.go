package graph

import (
	"github.com/song-flying/GoDataStructures/graph"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/hash"
	"github.com/song-flying/GoDataStructures/queue"
	"github.com/song-flying/GoDataStructures/set"
)

func BreathFirstSearch[V comparable](g *graph.UndirectedGraph[V], v V, w V) bool {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(v) && g.Contains(w), "g contains v and w")

	marked := set.NewHashSet[V](g.Size(), hash.Universal[V], 1)
	q := queue.NewLinkedQueue[V]()
	marked.Add(v)
	q.Enqueue(v)
	// v is in queue => v is marked

	for !q.IsEmpty() {
		x := q.Dequeue()
		if x == w {
			return true
		}

		neighbors := g.GetNeighbors(x)
		for curr := neighbors.Head; curr != nil; curr = curr.Next {
			y := curr.Data
			if !marked.Contains(y) {
				marked.Add(y)
				q.Enqueue(y)
			}
		}
	}

	return false
}
