package graph

import (
	"github.com/song-flying/GoDataStructures/linked"
)

type Graph[V comparable] interface {
	ContainsEdge(v, w V) bool
	AddEdge(v, w V)
	GetNeighbors(v V) (result *linked.List[V])
	Vertices() *linked.List[V]
	Contains(v V) bool
	Size() int
	Reverse() Graph[V]
}
