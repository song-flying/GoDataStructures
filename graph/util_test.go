package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestDistances(t *testing.T) {
	vertices := []string{"A", "B", "C", "D", "E", "F"}
	g := NewUndirectedGraph(vertices)
	g.AddEdge("A", "B")
	g.AddEdge("B", "C")
	g.AddEdge("C", "A")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")
	g.AddEdge("D", "F")
	g.AddEdge("E", "A")
	g.AddEdge("E", "F")
	// g looks like:
	//    A--------E
	//  /  \      / \
	// B----C----D---F

	distances := ShortestDistances(&g, "A")

	distToA, ok := distances.Get("A")
	assert.True(t, ok)
	assert.Equal(t, 0, *distToA)

	distToB, ok := distances.Get("B")
	assert.True(t, ok)
	assert.Equal(t, 1, *distToB)

	distToC, ok := distances.Get("C")
	assert.True(t, ok)
	assert.Equal(t, 1, *distToC)

	distToD, ok := distances.Get("D")
	assert.True(t, ok)
	assert.Equal(t, 2, *distToD)

	distToE, ok := distances.Get("E")
	assert.True(t, ok)
	assert.Equal(t, 1, *distToE)

	distToF, ok := distances.Get("F")
	assert.True(t, ok)
	assert.Equal(t, 2, *distToF)
}

func TestHasCycle(t *testing.T) {
	vertices := []string{"A", "B", "C", "D", "E", "F"}

	g := NewUndirectedGraph[string](vertices)
	g.AddEdge("A", "C")
	g.AddEdge("B", "C")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")
	g.AddEdge("F", "D")
	// g looks like:
	//    A      E
	//     \     |
	// B----C----D----F
	assert.False(t, HasCycle[string](&g))

	g.AddEdge("E", "F")
	assert.True(t, HasCycle[string](&g))
}
