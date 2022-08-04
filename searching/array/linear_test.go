package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	a := []int{1, 3, 5, 7}

	i := LinearSearch(5, a)
	assert.Equal(t, 2, i)

	i = LinearSearch(4, a)
	assert.Equal(t, -1, i)

	a = []int{}
	i = LinearSearch(1, a)
	assert.Equal(t, -1, i)
}
