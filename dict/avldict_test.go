package dict

import (
	"github.com/song-flying/GoDataStructures/array"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAVLDict(t *testing.T) {
	dict := NewAVLDict[int, string](func(m, n int) int { return m - n })

	v, ok := dict.Get(1)
	assert.False(t, ok)
	assert.Equal(t, "", *v)
	assert.Equal(t, 0, dict.Size())

	dict.Put(1, "a")
	v, ok = dict.Get(1)
	assert.True(t, ok)
	assert.Equal(t, "a", *v)
	assert.Equal(t, 1, dict.Size())

	dict.Put(2, "b")
	v, ok = dict.Get(2)
	assert.True(t, ok)
	assert.Equal(t, "b", *v)
	assert.Equal(t, 2, dict.Size())

	dict.Put(1, "aa")
	v, ok = dict.Get(1)
	assert.True(t, ok)
	assert.Equal(t, "aa", *v)
	assert.Equal(t, 2, dict.Size())

	dict.Delete(2)
	v, ok = dict.Get(2)
	assert.False(t, ok)
	assert.Equal(t, 1, dict.Size())

	dict.Delete(1)
	assert.Equal(t, 0, dict.Size())

	// Test repeated insertion and removal
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	var entries []entry[int, string]
	for i := 0; i < len(a); i++ {
		entries = append(entries, entry[int, string]{Key: a[i], Value: b[i]})
	}

	array.Shuffle(entries)
	t.Logf("entries to insert = %v", entries)
	for i, e := range entries {
		dict.Put(e.Key, e.Value)
		t.Logf("tree after insertion of key %d = %s", e.Key, dict.tree.String())
		v, ok := dict.Get(e.Key)
		assert.True(t, ok)
		assert.Equal(t, e.Value, *v)
		assert.Equal(t, i+1, dict.Size())
	}

	array.Shuffle(entries)
	t.Logf("entries to remove = %v", entries)
	for i, e := range entries {
		dict.Delete(e.Key)
		t.Logf("tree after removal of key %d = %s", e.Key, dict.tree.String())
		_, ok := dict.Get(e.Key)
		assert.False(t, ok)
		assert.Equal(t, len(a)-i-1, dict.Size())
	}
}
