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

func DepthFirstSearchR[V comparable](g *graph.UndirectedGraph[V], src V, dst V) bool {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(src) && g.Contains(dst), "g contains src and dst")

	marked := set.NewHashSet[V](g.Size(), hash.Universal[V], 1)
	return DepthFirstSearchRHelper[V](g, src, dst, marked)
}

func Some[T comparable](l *linked.List[T], pred func(T) bool) bool {
	for curr := l.Head; curr != nil; curr = curr.Next {
		if pred(curr.Data) {
			return true
		}
	}

	return false
}

func DepthFirstSearchRHelper[V comparable](g *graph.UndirectedGraph[V], v V, w V, marked set.Set[V]) bool {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(v) && g.Contains(w), "g contains v and w")
	contract.Require(!marked.Contains(v), "v is not marked")
	contract.Require(marked.IsEmpty() || Some(g.GetNeighbors(v), func(u V) bool { return marked.Contains(u) }), "one of v's neighbor is marked")
	defer func() {
		contract.Ensure(marked.Contains(v), "v is marked")
	}()

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
func DepthFirstSearchX[V comparable](g *graph.UndirectedGraph[V], src V, dst V) bool {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(src) && g.Contains(dst), "g contains src and dst")

	marked := set.NewHashSet[V](g.Size(), hash.Universal[V], 1)
	s := stack.NewLinkedStack[V]()
	marked.Add(src)
	s.Push(src)

	for !s.IsEmpty() {
		x := s.Pop()
		fmt.Println(x)
		if x == dst {
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

func DepthFirstSearch[V comparable](g *graph.UndirectedGraph[V], src V, dst V) bool {
	contract.Require(g != nil, "g is not nil")
	contract.Require(g.Contains(src) && g.Contains(dst), "g contains src and dst")

	marked := set.NewHashSet[V](g.Size(), hash.Universal[V], 1)
	s := stack.NewLinkedStack[*linked.ListIterator[V]]()
	marked.Add(src)
	vNode := linked.NewNode[V](src)
	vList := linked.NewList(&vNode)
	s.Push(vList.Iterator())

	for !s.IsEmpty() {
		xIter := s.Peek()
		if !xIter.HasNext() {
			s.Pop() // remove empty iterator
			continue
		}

		x := xIter.Next()
		fmt.Println(x)
		if x == dst {
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
