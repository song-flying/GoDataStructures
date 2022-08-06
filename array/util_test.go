package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCubes(t *testing.T) {
	a := Cubes(10)
	assert.Equal(t, 10, len(a))

	assert.Panics(t, func() {
		_ = Cubes(-1)
	})
}

func TestCopyArray(t *testing.T) {
	a := []int{1, 2, 3}

	b := CopyArray(a, 3)
	assert.Equal(t, 3, len(b))
	assert.Equal(t, 1, b[0])
	assert.Equal(t, 2, b[1])
	assert.Equal(t, 3, b[2])

	assert.Panics(t, func() {
		_ = CopyArray(a, 2)
	})
}

func TestSubArray(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	b := SubArray(a, 0, 0)
	assert.Equal(t, 0, len(b))

	c := SubArray(a, 0, 1)
	assert.Equal(t, 1, len(c))
	assert.Equal(t, 1, c[0])

	d := SubArray(a, 3, 5)
	assert.Equal(t, 2, len(d))
	assert.Equal(t, 4, d[0])
	assert.Equal(t, 5, d[1])

	assert.Panics(t, func() {
		_ = SubArray(a, 3, 2)
	})
}

func TestCopyInto(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{0, 0, 0, 0, 0}
	c := CopyArray(b, len(b))
	d := CopyArray(b, len(b))

	i := CopyInto(a, 0, 0, b, 1)
	assert.Equal(t, []int{0, 0, 0, 0, 0}, b)
	assert.Equal(t, -1, i)

	i = CopyInto(a, 0, 2, c, 1)
	assert.Equal(t, []int{0, 1, 2, 0, 0}, c)
	assert.Equal(t, 2, i)

	i = CopyInto(a, 4, 1, d, 0)
	assert.Equal(t, []int{5, 0, 0, 0, 0}, d)
	assert.Equal(t, 0, i)

	assert.Panics(t, func() {
		CopyInto(a, 4, 2, d, 0)
	})

	assert.Panics(t, func() {
		CopyInto(a, 0, 2, d, 4)
	})
}

func TestFindMax(t *testing.T) {
	a := []int{5, 3, 7, 1}
	i := FindMax(a, 4)
	assert.Equal(t, 2, i)

	a = []int{1, 2, 3}
	i = FindMax(a, 3)
	assert.Equal(t, 2, i)

	a = []int{3, 2, 1}
	i = FindMax(a, 3)
	assert.Equal(t, 0, i)

	a = []int{1}
	i = FindMax(a, 1)
	assert.Equal(t, 0, i)

	a = []int{}
	assert.Panics(t, func() {
		_ = FindMax(a, 0)
	})
}

func TestShuffle(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Shuffle(a)
	t.Log(a)
}
