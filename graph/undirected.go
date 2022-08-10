package graph

import (
	"github.com/song-flying/GoDataStructures/dict"
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/hash"
)

type UndirectedGraph[V comparable] struct {
	adjDict dict.HashDict[V, linked.List[V]]
}

func (g *UndirectedGraph[V]) hasVertex(v V) bool {
	_, ok := g.adjDict.Get(v)
	return ok
}

func (g *UndirectedGraph[V]) IsUndirectedGraph() bool {
	keys := g.adjDict.Keys()
	for curr := keys.Head; curr != nil; curr = curr.Next {
		v := curr.Data
		neighbors, ok := g.adjDict.Get(v)
		contract.Assert(ok, "vertices are initialized with neighbor list")
		for n := neighbors.Head; n != nil; n = n.Next {
			w := n.Data
			if w == v { // self loop
				return false
			}
			neighbors2, ok2 := g.adjDict.Get(w)
			if !ok2 || !neighbors2.Contains(v) { // relation is asymmetric
				return false
			}
			if !neighbors.IsDistinct() { // duplicate edges
				return false
			}
		}
	}

	return true
}

func NewUndirectedGraph[V comparable](vertices []V) (result UndirectedGraph[V]) {
	contract.Require(len(vertices) > 0, "vertices is not empty")
	defer func() {
		contract.Ensure(result.IsUndirectedGraph(), "graph invariant holds")
	}()

	adjDict := dict.NewHashDict[V, linked.List[V]](1, hash.Universal[V], 1)
	for _, v := range vertices {
		adjDict.Put(v, linked.NewEmptyList[V]())
	}

	return UndirectedGraph[V]{
		adjDict: adjDict,
	}
}

func (g *UndirectedGraph[V]) ContainsEdge(v, w V) bool {
	contract.Require(g.IsUndirectedGraph(), "graph invariant holds")
	contract.Require(g.hasVertex(v) && g.hasVertex(w), "g contains v and w")

	vNeighbors, _ := g.adjDict.Get(v)

	return vNeighbors.Contains(w)
}

func (g *UndirectedGraph[V]) AddEdge(v, w V) {
	contract.Require(g.IsUndirectedGraph(), "graph invariant holds")
	contract.Require(g.hasVertex(v) && g.hasVertex(w), "g contains v and w")
	contract.Require(v != w && !g.ContainsEdge(v, w), "g does not contain edge (v,w)")
	defer func() {
		contract.Ensure(g.IsUndirectedGraph(), "graph invariant holds")
		contract.Ensure(g.ContainsEdge(v, w), "g contains edge (v,w)")
	}()

	vNeighbors, _ := g.adjDict.Get(v)
	vNeighbors.Add(w)

	wNeighbors, _ := g.adjDict.Get(w)
	wNeighbors.Add(v)
}

func (g *UndirectedGraph[V]) GetNeighbors(v V) (result *linked.List[V]) {
	contract.Require(g.IsUndirectedGraph(), "graph invariant holds")
	contract.Require(g.hasVertex(v), "g contains v")

	neighbors, _ := g.adjDict.Get(v)

	return neighbors
}

func (g *UndirectedGraph[V]) Vertices() linked.List[V] {
	contract.Require(g.IsUndirectedGraph(), "graph invariant holds")
	return g.adjDict.Keys()
}

func (g *UndirectedGraph[V]) Size() int {
	contract.Require(g.IsUndirectedGraph(), "graph invariant holds")
	return g.adjDict.Size()
}

func (g *UndirectedGraph[V]) Contains(v V) bool {
	contract.Require(g.IsUndirectedGraph(), "graph invariant holds")

	_, ok := g.adjDict.Get(v)

	return ok
}
