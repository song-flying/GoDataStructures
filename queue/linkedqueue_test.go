package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedQueue(t *testing.T) {
	var q Queue[int]

	q = NewLinkedQueue[int]()
	assert.True(t, q.IsEmpty())

	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
	assert.Equal(t, 1, q.Head())

	q.Enqueue(2)
	assert.False(t, q.IsEmpty())
	assert.Equal(t, 1, q.Head())

	v := q.Dequeue()
	assert.False(t, q.IsEmpty())
	assert.Equal(t, 1, v)
	assert.Equal(t, 2, q.Head())

	v = q.Dequeue()
	assert.True(t, q.IsEmpty())
	assert.Equal(t, 2, v)
}
