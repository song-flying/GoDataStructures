package graph

import (
	"github.com/song-flying/GoDataStructures/dict"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/hash"
	"github.com/song-flying/GoDataStructures/queue"
)

func ShortestDistances[V comparable](g *UndirectedGraph[V], start V) (result dict.Dict[V, int]) {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(start), "g contains start")

	distances := dict.NewHashDict[V, int](1, hash.Universal[V], 1)
	q := queue.NewLinkedQueue[V]()
	q.Enqueue(start)

	for curr := g.Vertices().Head; curr != nil; curr = curr.Next {
		v := curr.Data
		distances.Put(v, -1)
	}

	distances.Put(start, 0)
	for !q.IsEmpty() {
		v := q.Dequeue()
		neighbors := g.GetNeighbors(v)
		distV, _ := distances.Get(v)

		for curr := neighbors.Head; curr != nil; curr = curr.Next {
			w := curr.Data
			distW, _ := distances.Get(w)
			if *distW == -1 {
				q.Enqueue(w)
				distances.Put(w, *distV+1)
			}
		}
	}

	return &distances
}
