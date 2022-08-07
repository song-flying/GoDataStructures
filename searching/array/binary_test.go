package array

import (
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 7}

	i := BinarySearch(5, a, order.IntComp)
	assert.Equal(t, 2, i)

	i = BinarySearch(4, a, order.IntComp)
	assert.Equal(t, -1, i)

	a = []int{42}
	i = BinarySearch(1, a, order.IntComp)
	assert.Equal(t, -1, i)

	i = BinarySearch(42, a, order.IntComp)
	assert.Equal(t, 0, i)

	a = []int{}
	i = BinarySearch(1, a, order.IntComp)
	assert.Equal(t, -1, i)
}
