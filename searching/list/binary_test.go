package list

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	n := linked.Node[int]{
		Data: 1,
		Next: &linked.Node[int]{
			Data: 3,
			Next: &linked.Node[int]{
				Data: 5,
				Next: &linked.Node[int]{
					Data: 7,
				},
			},
		},
	}

	l := linked.NewList[int](&n)

	i := BinarySearch(5, l)
	assert.Equal(t, 2, i)

	i = BinarySearch(4, l)
	assert.Equal(t, -1, i)

	l = linked.NewList[int](&linked.Node[int]{Data: 42})
	i = BinarySearch(1, l)
	assert.Equal(t, -1, i)

	i = BinarySearch(42, l)
	assert.Equal(t, 0, i)

	l = linked.NewEmptyList[int]()
	i = BinarySearch(1, l)
	assert.Equal(t, -1, i)
}
