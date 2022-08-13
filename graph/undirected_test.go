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
	checkNeighbors(t, &g, order.StringComp, "A", []string{"B", "C", "E"})
	checkNeighbors(t, &g, order.StringComp, "B", []string{"A", "C"})
	checkNeighbors(t, &g, order.StringComp, "C", []string{"A", "B", "D"})
	checkNeighbors(t, &g, order.StringComp, "D", []string{"C", "E", "F"})
	checkNeighbors(t, &g, order.StringComp, "E", []string{"A", "D"})
	checkNeighbors(t, &g, order.StringComp, "F", []string{"D"})

	assert.True(t, g.HasCycle())

	g = NewUndirectedGraph(vertices)
	assert.Equal(t, len(vertices), g.Size())

	g.AddEdge("A", "C")
	g.AddEdge("B", "C")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")
	g.AddEdge("D", "F")
	// g looks like:
	//    A      E
	//     \     |
	// B----C----D----F

	assert.False(t, g.HasCycle())
}

func checkNeighbors[V comparable](t *testing.T, g *UndirectedGraph[V], comp order.CompareFn[V], v V, expected []V) {
	t.Helper()

	neighbors := g.GetNeighbors(v).ToArray()
	sorting.SelectionSort(neighbors, comp)
	assert.Equal(t, expected, neighbors)
}
