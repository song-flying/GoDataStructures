package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedStack(t *testing.T) {
	var s Stack[int]

	s = NewLinkedStack[int]()
	assert.True(t, s.IsEmpty())

	s.Push(1)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 1, s.Peek())

	s.Push(2)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 2, s.Peek())

	v := s.Pop()
	assert.Equal(t, 2, v)
	assert.Equal(t, 1, s.Peek())

	v = s.Pop()
	assert.Equal(t, 1, v)
	assert.True(t, s.IsEmpty())
}
