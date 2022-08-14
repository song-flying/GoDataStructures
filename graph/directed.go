package graph

import (
	"github.com/song-flying/GoDataStructures/dict"
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/hash"
)

type DirectedGraph[V comparable] struct {
	adjDict dict.HashDict[V, linked.List[V]]
}

func (g *DirectedGraph[V]) hasVertex(v V) bool {
	_, ok := g.adjDict.Get(v)
	return ok
}

func (g *DirectedGraph[V]) IsDirectedGraph() bool {
	keys := g.adjDict.Keys().Iterator()
	for keys.HasNext() {
		v := keys.Next()
		neighbors, ok := g.adjDict.Get(v)
		contract.Assert(ok, "vertices are initialized with neighbor list")
		if neighbors.Contains(v) { // self loop
			return false
		}
		if !neighbors.IsDistinct() { // duplicate edges
			return false
		}
	}

	return true
}

func NewDirectedGraph[V comparable](vertices []V) (result DirectedGraph[V]) {
	contract.Require(len(vertices) > 0, "vertices is not empty")
	defer func() {
		contract.Ensure(result.IsDirectedGraph(), "graph invariant holds")
	}()

	adjDict := dict.NewHashDict[V, linked.List[V]](1, hash.Universal[V], 1)
	for _, v := range vertices {
		adjDict.Put(v, linked.NewEmptyList[V]())
	}

	return DirectedGraph[V]{
		adjDict: adjDict,
	}
}

func (g *DirectedGraph[V]) ContainsEdge(v, w V) bool {
	contract.Require(g.IsDirectedGraph(), "graph invariant holds")
	contract.Require(g.hasVertex(v) && g.hasVertex(w), "g contains v and w")

	vNeighbors, _ := g.adjDict.Get(v)

	return vNeighbors.Contains(w)
}

func (g *DirectedGraph[V]) AddEdge(v, w V) {
	contract.Require(g.IsDirectedGraph(), "graph invariant holds")
	contract.Require(g.hasVertex(v) && g.hasVertex(w), "g contains v and w")
	contract.Require(v != w && !g.ContainsEdge(v, w), "g does not contain edge (v,w)")
	defer func() {
		contract.Ensure(g.IsDirectedGraph(), "graph invariant holds")
		contract.Ensure(g.ContainsEdge(v, w), "g contains edge (v,w)")
	}()

	vNeighbors, _ := g.adjDict.Get(v)
	vNeighbors.Add(w)
}

func (g *DirectedGraph[V]) GetNeighbors(v V) (result *linked.List[V]) {
	contract.Require(g.IsDirectedGraph(), "graph invariant holds")
	contract.Require(g.hasVertex(v), "g contains v")

	neighbors, _ := g.adjDict.Get(v)

	return neighbors
}

func (g *DirectedGraph[V]) Vertices() *linked.List[V] {
	contract.Require(g.IsDirectedGraph(), "graph invariant holds")

	return g.adjDict.Keys()
}

func (g *DirectedGraph[V]) Size() int {
	contract.Require(g.IsDirectedGraph(), "graph invariant holds")

	return g.adjDict.Size()
}

func (g *DirectedGraph[V]) Contains(v V) bool {
	contract.Require(g.IsDirectedGraph(), "graph invariant holds")

	_, ok := g.adjDict.Get(v)
	return ok
}

func (g *DirectedGraph[V]) Reverse() Graph[V] {
	contract.Require(g.IsDirectedGraph(), "graph invariant holds")

	gReverse := NewDirectedGraph(g.Vertices().ToArray())

	vertices := g.Vertices().Iterator()
	for vertices.HasNext() {
		v := vertices.Next()
		neighbors := g.GetNeighbors(v).Iterator()
		for neighbors.HasNext() {
			w := neighbors.Next()
			gReverse.AddEdge(w, v)
		}
	}

	return &gReverse
}
