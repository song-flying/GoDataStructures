package graph

import (
	"fmt"
	"github.com/song-flying/GoDataStructures/graph"
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/hash"
	"github.com/song-flying/GoDataStructures/set"
	"github.com/song-flying/GoDataStructures/stack"
)

func DepthFirstSearchR[V comparable](g *graph.UndirectedGraph[V], v V, w V) bool {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(v) && g.Contains(w), "g contains v and w")

	marked := set.NewHashSet[V](g.Size(), hash.Universal[V], 1)
	return DepthFirstSearchRHelper[V](g, v, w, &marked)
}

func DepthFirstSearchRHelper[V comparable](g *graph.UndirectedGraph[V], v V, w V, marked set.Set[V]) bool {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(v) && g.Contains(w), "g contains v and w")

	marked.Add(v)
	fmt.Println(v)
	if v == w {
		return true
	}

	neighbors := g.GetNeighbors(v)
	for curr := neighbors.Head; curr != nil; curr = curr.Next {
		u := curr.Data
		if !marked.Contains(u) && DepthFirstSearchRHelper(g, u, w, marked) {
			return true
		}
	}

	return false
}

// DepthFirstSearchX NOT equivalent to DepthFirstSearchR
func DepthFirstSearchX[V comparable](g *graph.UndirectedGraph[V], v V, w V) bool {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(v) && g.Contains(w), "g contains v and w")

	marked := set.NewHashSet[V](g.Size(), hash.Universal[V], 1)
	s := stack.NewLinkedStack[V]()
	marked.Add(v)
	s.Push(v)
	// v is on stack => v is marked

	for !s.IsEmpty() {
		x := s.Pop()
		fmt.Println(x)
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

func DepthFirstSearch[V comparable](g *graph.UndirectedGraph[V], v V, w V) bool {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(v) && g.Contains(w), "g contains v and w")

	marked := set.NewHashSet[V](g.Size(), hash.Universal[V], 1)
	s := stack.NewLinkedStack[*linked.ListIterator[V]]()
	marked.Add(v)
	vNode := linked.NewNode[V](v)
	vList := linked.NewList(&vNode)
	s.Push(vList.Iterator())
	// v is on stack => v is marked

	for !s.IsEmpty() {
		xIter := s.Peek()
		if !xIter.HasNext() {
			s.Pop() // remove empty iterator
			continue
		}

		x := xIter.Next()
		fmt.Println(x)
		if x == w {
			return true
		}

		neighbors := g.GetNeighbors(x).Iterator()
		for neighbors.HasNext() {
			y := neighbors.Next()
			if !marked.Contains(y) {
				marked.Add(y)

				s.Push(neighbors)

				yNode := linked.NewNode[V](y)
				yList := linked.NewList(&yNode)
				s.Push(yList.Iterator())
				break
			}
		}
	}

	return false
}
