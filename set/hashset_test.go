package set

import (
	"github.com/stretchr/testify/assert"
	"hash/maphash"
	"testing"
)

var h maphash.Hash
var hashFn = func(s string) int {
	_, _ = h.WriteString(s)
	defer h.Reset()
	return int(h.Sum64())
}

func TestHashSet(t *testing.T) {
	set := NewHashSet[string](10, hashFn)
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
}
