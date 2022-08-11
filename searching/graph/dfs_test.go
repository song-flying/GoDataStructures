package graph

import (
	"github.com/song-flying/GoDataStructures/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDepthFirstSearchR(t *testing.T) {
	vertices := []string{"A", "B", "C", "D", "E", "F"}
	g := graph.NewUndirectedGraph(vertices)

	g.AddEdge("A", "C")
	g.AddEdge("A", "B")
	g.AddEdge("B", "D")
	g.AddEdge("B", "C")
	g.AddEdge("C", "E")
	// g looks like:
	//      A
	//    /  \
	//   B----C----E
	//  /
	// D

	assert.True(t, DepthFirstSearchR(&g, "A", "E"))
}

func TestDepthFirstSearchX(t *testing.T) {
	vertices := []string{"A", "B", "C", "D", "E", "F"}
	g := graph.NewUndirectedGraph(vertices)

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("B", "C")
	g.AddEdge("C", "E")
	// g looks like:
	//      A
	//    /  \
	//   B----C----E
	//  /
	// D
	assert.True(t, DepthFirstSearchX(&g, "A", "E"))
}

func TestDepthFirstSearch(t *testing.T) {
	vertices := []string{"A", "B", "C", "D", "E", "F"}
	g := graph.NewUndirectedGraph(vertices)

	g.AddEdge("A", "C")
	g.AddEdge("A", "B")
	g.AddEdge("B", "D")
	g.AddEdge("B", "C")
	g.AddEdge("C", "E")
	// g looks like:
	//      A
	//    /  \
	//   B----C----E
	//  /
	// D

	assert.True(t, DepthFirstSearch(&g, "A", "E"))
}
