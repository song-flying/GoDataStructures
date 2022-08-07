package array

import (
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSortedSearch(t *testing.T) {
	a := []int{1, 3, 5, 7}

	i := LinearSortedSearch(5, a, order.IntComp)
	assert.Equal(t, 2, i)

	i = LinearSortedSearch(4, a, order.IntComp)
	assert.Equal(t, -1, i)

	a = []int{}
	i = LinearSortedSearch(1, a, order.IntComp)
	assert.Equal(t, -1, i)
}
