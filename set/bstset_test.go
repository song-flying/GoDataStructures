package set

import (
	"github.com/song-flying/GoDataStructures/array"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBSTSet(t *testing.T) {
	set := NewBSTSet[int](func(m, n int) int { return m - n })

	assert.False(t, set.Contains(1))
	assert.Equal(t, 0, set.Size())

	set.Add(1)
	assert.True(t, set.Contains(1))
	assert.Equal(t, 1, set.Size())

	set.Add(2)
	assert.True(t, set.Contains(2))
	assert.Equal(t, 2, set.Size())

	set.Add(1)
	assert.True(t, set.Contains(1))
	assert.Equal(t, 2, set.Size())

	set.Delete(2)
	assert.False(t, set.Contains(2))
	assert.Equal(t, 1, set.Size())

	set.Delete(1)
	assert.Equal(t, 0, set.Size())

	// Test repeated insertion and removal
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	array.Shuffle(a)
	t.Logf("elements to insert = %v", a)
	for i, e := range a {
		set.Add(e)
		t.Logf("tree after insertion of element %d = %s", e, set.tree.String())
		assert.True(t, set.Contains(e))
		assert.Equal(t, i+1, set.Size())
	}

	array.Shuffle(a)
	t.Logf("elements to delete = %v", a)
	for i, e := range a {
		set.Delete(e)
		t.Logf("tree after deletion of element %d = %s", e, set.tree.String())
		assert.False(t, set.Contains(e))
		assert.Equal(t, len(a)-i-1, set.Size())

	}
}
