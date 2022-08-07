package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryNode_ToArrayOrder(t *testing.T) {
	root := &BinaryNode[int]{Data: 1}
	l := &BinaryNode[int]{Data: 2}
	r := &BinaryNode[int]{Data: 3}
	ll := &BinaryNode[int]{Data: 4}
	lr := &BinaryNode[int]{Data: 5}
	rl := &BinaryNode[int]{Data: 6}
	rr := &BinaryNode[int]{Data: 7}
	root.Left = l
	root.Right = r
	l.Left = ll
	l.Right = lr
	r.Left = rl
	r.Right = rr
	// tree looks like:
	//        1
	//   2         3
	// 4   5     6   7

	assert.Equal(t, []int{4, 2, 5, 1, 6, 3, 7}, root.ToArrayPreorder())
	assert.Equal(t, []int{1, 2, 4, 5, 3, 6, 7}, root.ToArrayInorder())
	assert.Equal(t, []int{4, 5, 2, 6, 7, 3, 1}, root.ToArrayPostorder())
}
