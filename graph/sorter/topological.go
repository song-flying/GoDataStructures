package sorter

import (
	"github.com/song-flying/GoDataStructures/dict"
	"github.com/song-flying/GoDataStructures/graph"
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/song-flying/GoDataStructures/pkg/hash"
	"github.com/song-flying/GoDataStructures/set"
)

type TopologicalSorter[V comparable] struct {
	graph    *graph.DirectedGraph[V]
	visited  *set.HashSet[V]
	dfsStack *set.HashSet[V]
	sorted   *linked.List[V]
	hasCycle bool
	cycle    *linked.List[V]
	from     dict.Dict[V, V]
}

func NewTopologicalSorter[V comparable](g *graph.DirectedGraph[V]) TopologicalSorter[V] {
	return TopologicalSorter[V]{
		graph:    g,
		visited:  set.NewHashSet[V](1, hash.Universal[V], 1),
		dfsStack: set.NewHashSet[V](1, hash.Universal[V], 1),
		sorted:   linked.NewEmptyList[V](),
		hasCycle: false,
		cycle:    linked.NewEmptyList[V](),
		from:     dict.NewHashDict[V, V](1, hash.Universal[V], 1),
	}
}

func (s *TopologicalSorter[V]) Sort() bool {
	vertices := s.graph.Vertices().Iterator()
	for vertices.HasNext() {
		v := vertices.Next()
		if !s.visited.Contains(v) {
			if s.dfs(v) {
				return true
			}
		}
	}

	return false
}

func (s *TopologicalSorter[V]) Sorted() (result []V) {
	return s.sorted.ToArray()
}

func (s TopologicalSorter[V]) Cycle() (result []V) {
	return s.cycle.ToArray()
}

func (s *TopologicalSorter[V]) dfs(v V) bool {
	s.dfsStack.Add(v)
	defer s.dfsStack.Delete(v)

	s.visited.Add(v)

	neighbors := s.graph.GetNeighbors(v).Iterator()
	for neighbors.HasNext() {
		w := neighbors.Next()
		if !s.visited.Contains(w) {
			s.from.Put(w, v)
			if s.dfs(w) {
				return true
			}
		} else if s.dfsStack.Contains(w) {
			s.from.Put(w, v)
			s.hasCycle = true
			s.buildCycle(w)
			return true
		}
	}

	s.sorted.Add(v)

	return false
}

func (s *TopologicalSorter[V]) buildCycle(w V) {
	var (
		ok bool
		v  V
	)
	v = w
	s.cycle.Add(v)
	for v, ok = s.from.Get(v); ok; {
		s.cycle.Add(v)
		if v == w {
			break
		}
		v, ok = s.from.Get(v)
	}
}
