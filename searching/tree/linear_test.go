package tree

import (
	"github.com/song-flying/GoDataStructures/tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	binTree := tree.NewBinaryTree[int](&tree.BinaryNode[int]{
		Data: 5,
		Left: &tree.BinaryNode[int]{
			Data: 4,
			Left: &tree.BinaryNode[int]{
				Data: 6,
			},
			Right: &tree.BinaryNode[int]{
				Data: 7,
			},
		},
		Right: &tree.BinaryNode[int]{
			Data: 3,
			Left: &tree.BinaryNode[int]{
				Data: 2,
			},
			Right: &tree.BinaryNode[int]{
				Data: 9,
			},
		},
	})

	assert.Nil(t, LinearSearch(1, binTree))
	assert.Nil(t, LinearSearch(8, binTree))
	assert.NotNil(t, LinearSearch(6, binTree))
	assert.NotNil(t, LinearSearch(2, binTree))
}
