package heap

import (
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap[int](9, order.IntComp)

	a := []int{9, 5, 8, 7, 1, 4, 3, 6, 2}

	for i, v := range a {
		h.Add(v)
		assert.True(t, h.Contains(v))
		assert.Equal(t, i+1, h.Size())
	}

	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range b {
		assert.Equal(t, v, h.Delete())
		assert.False(t, h.Contains(v))
		assert.Equal(t, len(b)-i-1, h.Size())
	}
}
