package list

import (
	"github.com/song-flying/GoDataStructures/linked"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	head := &linked.Node[int]{
		Data: 5,
		Next: &linked.Node[int]{
			Data: 3,
			Next: &linked.Node[int]{
				Data: 1,
			},
		},
	}

	l := linked.NewList(head)
	assert.Equal(t, nil, LinearSearch(2, l))

	assert.Equal(t, head.Next, LinearSearch(3, l))
}
