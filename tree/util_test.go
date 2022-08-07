package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCycle(t *testing.T) {
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

	// connect 4 to 1 via left link
	ll.Left = root
	assert.True(t, hasCycle(root))
	ll.Left = nil
	assert.False(t, hasCycle(root))

	// connect 4 to 5 via right link
	ll.Right = lr
	assert.True(t, hasCycle(root))
	ll.Right = nil
	assert.False(t, hasCycle(root))

	// connect 5 to 3 via left link
	lr.Left = r
	assert.True(t, hasCycle(root))
	lr.Left = nil
	assert.False(t, hasCycle(root))

	// connect 5 to 6 via right link
	lr.Right = rl
	assert.True(t, hasCycle(root))
	lr.Right = nil
	assert.False(t, hasCycle(root))

	// connect 7 to 2 via left link
	rr.Right = l
	assert.True(t, hasCycle(root))
	rr.Right = nil
	assert.False(t, hasCycle(root))
}
