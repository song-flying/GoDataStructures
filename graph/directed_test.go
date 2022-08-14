package graph

import (
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectedGraph(t *testing.T) {
	vertices := []string{"A", "B", "C", "D", "E", "F"}
	g := NewDirectedGraph(vertices)

	assert.Equal(t, len(vertices), g.Size())

	g.AddEdge("A", "B")
	g.AddEdge("B", "C")
	g.AddEdge("C", "A")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")
	g.AddEdge("D", "F")
	g.AddEdge("E", "A")
	// g looks like:
	// -----A<----E
	// |   /\    /\
	//\/    |     |
	// B--->C---->D---->F
	checkNeighbors[string](t, &g, order.StringComp, "A", []string{"B"})
	checkNeighbors[string](t, &g, order.StringComp, "B", []string{"C"})
	checkNeighbors[string](t, &g, order.StringComp, "C", []string{"A", "D"})
	checkNeighbors[string](t, &g, order.StringComp, "D", []string{"E", "F"})
	checkNeighbors[string](t, &g, order.StringComp, "E", []string{"A"})
	checkNeighbors[string](t, &g, order.StringComp, "F", nil)
}
