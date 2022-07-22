package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := []int{3, 9, 5, 7, 1}
	InsertionSort(a)
	assert.Equal(t, []int{1, 3, 5, 7, 9}, a)

	a = []int{42}
	InsertionSort(a)
	assert.Equal(t, []int{42}, a)
}
