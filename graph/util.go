package graph

import (
	"github.com/song-flying/GoDataStructures/dict"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/hash"
	"github.com/song-flying/GoDataStructures/queue"
	"github.com/song-flying/GoDataStructures/set"
)

func ShortestDistances[V comparable](g *UndirectedGraph[V], start V) (result dict.Dict[V, int]) {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(start), "g contains start")

	distances := dict.NewHashDict[V, int](1, hash.Universal[V], 1)
	marked := set.NewHashSet[V](1, hash.Universal[V], 1)
	q := queue.NewLinkedQueue[V]()
	q.Enqueue(start)
	marked.Add(start)

	distances.Put(start, 0)
	for !q.IsEmpty() {
		v := q.Dequeue()
		neighbors := g.GetNeighbors(v)
		dist, ok := distances.Get(v)
		contract.Assert(ok, "v's distance is already computed")
		for curr := neighbors.Head; curr != nil; curr = curr.Next {
			w := curr.Data
			if !marked.Contains(w) {
				q.Enqueue(w)
				marked.Add(w)
				distances.Put(w, *dist+1)
			}
		}
	}

	return &distances
}
