package graph

import (
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/song-flying/GoDataStructures/sorting"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUndirectedGraph(t *testing.T) {
	vertices := []string{"A", "B", "C", "D", "E", "F"}
	g := NewUndirectedGraph(vertices)

	assert.Equal(t, len(vertices), g.Size())

	g.AddEdge("A", "B")
	g.AddEdge("B", "C")
	g.AddEdge("C", "A")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")
	g.AddEdge("D", "F")
	g.AddEdge("E", "A")
	// g looks like:
	//    A------E
	//  /  \     |
	// B----C----D----F
	checkNeighbors[string](t, g, order.StringComp, "A", []string{"B", "C", "E"})
	checkNeighbors[string](t, g, order.StringComp, "B", []string{"A", "C"})
	checkNeighbors[string](t, g, order.StringComp, "C", []string{"A", "B", "D"})
	checkNeighbors[string](t, g, order.StringComp, "D", []string{"C", "E", "F"})
	checkNeighbors[string](t, g, order.StringComp, "E", []string{"A", "D"})
	checkNeighbors[string](t, g, order.StringComp, "F", []string{"D"})
}

func checkNeighbors[V comparable](t *testing.T, g Graph[V], comp order.CompareFn[V], v V, expected []V) {
	t.Helper()

	neighbors := g.GetNeighbors(v).ToArray()
	sorting.SelectionSort(neighbors, comp)
	assert.Equal(t, expected, neighbors)
}
