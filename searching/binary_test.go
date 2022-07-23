package searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 7}

	i := BinarySearch(5, a)
	assert.Equal(t, 2, i)

	i = BinarySearch(4, a)
	assert.Equal(t, -1, i)

	a = []int{}
	i = BinarySearch(1, a)
	assert.Equal(t, -1, i)
}
