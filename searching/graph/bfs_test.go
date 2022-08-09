package graph

import (
	"github.com/song-flying/GoDataStructures/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreathFirstSearch(t *testing.T) {
	vertices := []string{"A", "B", "C", "D", "E", "F"}
	g := graph.NewUndirectedGraph(vertices)

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

	assert.True(t, BreathFirstSearch(&g, "A", "F"))
}
