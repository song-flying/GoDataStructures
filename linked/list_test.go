package linked

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	l := NewEmptyList[int]()

	l.Add(5)
	assert.True(t, l.Contains(5))
	l.Add(1)
	assert.True(t, l.Contains(5))
	l.Add(3)
	assert.True(t, l.Contains(5))

	assert.Equal(t, []int{3, 1, 5}, l.ToArray())
	assert.True(t, l.IsDistinct())

	l.Add(1)
	assert.False(t, l.IsDistinct())
}
