package set

import (
	"github.com/stretchr/testify/assert"
	"hash/maphash"
	"strconv"
	"testing"
)

var h maphash.Hash
var hashFn = func(s string) int {
	_, _ = h.WriteString(s)
	defer h.Reset()
	return int(h.Sum64())
}

func TestHashSet(t *testing.T) {
	set := NewHashSet[string](1, hashFn, 1)
	assert.False(t, set.Contains("hello"))
	assert.Equal(t, 0, set.Size())

	set.Add("hello")
	assert.True(t, set.Contains("hello"))
	assert.Equal(t, 1, set.Size())

	set.Add("world")
	assert.True(t, set.Contains("world"))
	assert.Equal(t, 2, set.Size())

	set.Add("hello")
	assert.True(t, set.Contains("hello"))
	assert.Equal(t, 2, set.Size())

	set.Remove("world")
	assert.False(t, set.Contains("world"))
	assert.Equal(t, 1, set.Size())

	set.Remove("hello")
	assert.Equal(t, 0, set.Size())

	for i := 0; i < 8; i++ {
		set.Add(strconv.Itoa(i))
	}
	assert.Equal(t, 16, set.capacity)

	for i := 0; i < 6; i++ {
		set.Remove(strconv.Itoa(i))
	}
	assert.Equal(t, 4, set.capacity)
}
