package sorting

import (
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{3, 9, 5, 7, 1}
	QuickSort(a, order.IntComp)
	assert.Equal(t, []int{1, 3, 5, 7, 9}, a)

	a = []int{42}
	QuickSort(a, order.IntComp)
	assert.Equal(t, []int{42}, a)
}
