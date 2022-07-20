package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cubes(t *testing.T) {
	a := cubes(10)
	assert.Equal(t, 10, len(a))

	assert.Panics(t, func() {
		_ = cubes(-1)
	})
}

func Test_copyArray(t *testing.T) {
	a := []int{1, 2, 3}

	b := copyArray(a, 3)
	assert.Equal(t, 3, len(b))
	assert.Equal(t, 1, b[0])
	assert.Equal(t, 2, b[1])
	assert.Equal(t, 3, b[2])

	assert.Panics(t, func() {
		_ = copyArray(a, 2)
	})
}

func Test_subArray(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	b := subArray(a, 0, 0)
	assert.Equal(t, 0, len(b))

	c := subArray(a, 0, 1)
	assert.Equal(t, 1, len(c))
	assert.Equal(t, 1, c[0])

	d := subArray(a, 3, 5)
	assert.Equal(t, 2, len(d))
	assert.Equal(t, 4, d[0])
	assert.Equal(t, 5, d[1])

	assert.Panics(t, func() {
		_ = subArray(a, 3, 2)
	})
}

func Test_copyInto(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{0, 0, 0, 0, 0}
	c := copyArray(b, len(b))
	d := copyArray(b, len(b))

	i := copyInto(a, 0, 0, b, 1)
	assert.Equal(t, []int{0, 0, 0, 0, 0}, b)
	assert.Equal(t, -1, i)

	i = copyInto(a, 0, 2, c, 1)
	assert.Equal(t, []int{0, 1, 2, 0, 0}, c)
	assert.Equal(t, 2, i)

	i = copyInto(a, 4, 1, d, 0)
	assert.Equal(t, []int{5, 0, 0, 0, 0}, d)
	assert.Equal(t, 0, i)

	assert.Panics(t, func() {
		copyInto(a, 4, 2, d, 0)
	})

	assert.Panics(t, func() {
		copyInto(a, 0, 2, d, 4)
	})
}

func Test_search(t *testing.T) {
	a := []int{1, 3, 5, 7}

	i := search(5, a, 4)
	assert.Equal(t, 2, i)

	i = search(2, a, 4)
	assert.Equal(t, -1, i)

	b := []int{2, 1, 3}
	assert.Panics(t, func() {
		_ = search(1, b, 3)
	})
}
