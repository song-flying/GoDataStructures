package sorter

import (
	"github.com/song-flying/GoDataStructures/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopologicalSorter(t *testing.T) {
	vertices := []string{"A", "B", "C", "D", "E", "F", "G"}

	g := graph.NewDirectedGraph[string](vertices)
	g.AddEdge("E", "G")
	g.AddEdge("D", "F")
	g.AddEdge("B", "C")
	g.AddEdge("A", "D")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")
	// g looks like
	//      A     E - G
	//        \ /
	//         D
	//        / \
	//  B - C     F
	sorter := NewTopologicalSorter[string](g)

	assert.False(t, sorter.Sort())
	t.Log(sorter.Sorted())

	g.AddEdge("G", "D")
	g.AddEdge("D", "B")

	sorter = NewTopologicalSorter[string](g)
	assert.True(t, sorter.Sort())
	t.Log(sorter.Cycle())
}
