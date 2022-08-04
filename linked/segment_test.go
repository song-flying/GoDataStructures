package linked

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCycle(t *testing.T) {
	assert.False(t, HasCycle[any](nil))

	acyclic := &Node[int]{
		Data: 1,
		Next: &Node[int]{
			Data: 2,
			Next: &Node[int]{
				Data: 3,
				Next: nil,
			},
		},
	}
	assert.False(t, HasCycle(acyclic))

	cyclic := &Node[int]{
		Data: 1,
		Next: nil,
	}
	cyclic.Next = cyclic
	assert.True(t, HasCycle(cyclic), "self cycle")

	cyclic = &Node[int]{
		Data: 1,
		Next: &Node[int]{
			Data: 2,
			Next: nil,
		},
	}
	cycle := &Node[int]{
		Data: 3,
		Next: &Node[int]{
			Data: 4,
			Next: nil,
		},
	}
	cycle.Next.Next = cycle
	cyclic.Next.Next = cycle
	assert.True(t, HasCycle(cyclic))
}

func TestIsReachableWith(t *testing.T) {
	assert.True(t, isReachableWith[any](nil, nil, 0))
	assert.True(t, isReachableWith(&Node[int]{Data: 1}, nil, 1))

	src := &Node[int]{
		Data: 1,
		Next: &Node[int]{
			Data: 2,
			Next: nil,
		},
	}
	dst := &Node[int]{
		Data: 3,
		Next: &Node[int]{
			Data: 4,
			Next: nil,
		},
	}
	assert.False(t, isReachableWith(src, dst, 100))

	src.Next.Next = dst
	assert.False(t, isReachableWith(src, dst, 1))
	assert.True(t, isReachableWith(src, dst, 2))
	assert.False(t, isReachableWith(src, dst, 3))
}

func TestNext(t *testing.T) {
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	assert.Equal(t, n1, next(n1, 0))
	assert.Equal(t, n2, next(n1, 1))
	assert.Equal(t, n3, next(n1, 2))
	assert.Equal(t, n4, next(n1, 3))
	assert.Equal(t, Nil[int](), next(n1, 4))
	assert.Equal(t, Nil[int](), next(n4, 2))
}

func TestIsSegment(t *testing.T) {
	assert.False(t, IsSegment[int](nil, nil))

	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)

	assert.True(t, IsSegment(n1, nil))
	assert.False(t, IsSegment(nil, n1))
	assert.True(t, IsSegment(n1, n1))
	assert.False(t, IsSegment(n1, n2))

	n1.Next = n2
	assert.True(t, IsSegment(n1, n2))
	assert.False(t, IsSegment(n1, n3))

	n2.Next = n3
	assert.True(t, IsSegment(n1, n3))
	assert.True(t, IsSegment(n2, n3))
}

func TestLengthOfSegment(t *testing.T) {
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)

	assert.Equal(t, 0, LengthOfSegment(n1, n1))

	n1.Next = n2
	assert.Equal(t, 1, LengthOfSegment(n1, n2))

	n2.Next = n3
	assert.Equal(t, 1, LengthOfSegment(n2, n3))
	assert.Equal(t, 2, LengthOfSegment(n1, n3))
}

func TestIthSegment(t *testing.T) {
	assert.Equal(t, 1, IthSegment(NewNode(1), 0))

	l := &Node[int]{
		Data: 1,
		Next: &Node[int]{
			Data: 2,
			Next: &Node[int]{
				Data: 3,
			},
		},
	}

	assert.Equal(t, 1, IthSegment(l, 0))
	assert.Equal(t, 2, IthSegment(l, 1))
	assert.Equal(t, 3, IthSegment(l, 2))
}
