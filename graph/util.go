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

func HasCycleUndirected[V comparable](g *UndirectedGraph[V]) bool {
	vertices := g.Vertices().Iterator()
	marked := set.NewHashSet[V](g.Size(), hash.Universal[V], 1)
	for vertices.HasNext() {
		v := vertices.Next()
		if !marked.Contains(v) {
			if dfsHasCycleUndirected[V](g, v, v, &marked) {
				return true
			}
		}
	}

	return false
}

func dfsHasCycleUndirected[V comparable](g *UndirectedGraph[V], v, from V, marked set.Set[V]) bool {
	marked.Add(v)
	neighbors := g.GetNeighbors(v).Iterator()
	for neighbors.HasNext() {
		w := neighbors.Next()
		if !marked.Contains(w) {
			if dfsHasCycleUndirected(g, w, v, marked) {
				return true
			}
		} else if w != from {
			return true
		}
	}

	return false
}

func HasCycleDirected[V comparable](g *DirectedGraph[V]) bool {
	vertices := g.Vertices().Iterator()
	marked := set.NewHashSet[V](g.Size(), hash.Universal[V], 1)
	callStack := set.NewHashSet[V](g.Size(), hash.Universal[V], 1)
	for vertices.HasNext() {
		v := vertices.Next()
		if !marked.Contains(v) {
			if dfsHasCycleDirected[V](g, v, &marked, &callStack) {
				return true
			}
		}
	}

	return false
}

func dfsHasCycleDirected[V comparable](g *DirectedGraph[V], v V, marked set.Set[V], callStack set.Set[V]) bool {
	callStack.Add(v)
	marked.Add(v)

	neighbors := g.GetNeighbors(v).Iterator()
	for neighbors.HasNext() {
		w := neighbors.Next()
		if !marked.Contains(w) && dfsHasCycleDirected(g, w, marked, callStack) {
			return true
		} else if callStack.Contains(w) {
			return true
		}
	}

	callStack.Delete(v)

	return false
}
