package set

import (
	"github.com/stretchr/testify/assert"
	"hash/maphash"
	"testing"
)

func TestHashSet(t *testing.T) {
	var h maphash.Hash
	hashFn := func(s string) int {
		_, _ = h.WriteString(s)
		defer h.Reset()
		return int(h.Sum64())
	}

	set := NewHashSet[string](10, hashFn)
	assert.False(t, set.Contains("hello"))

	set.Add("hello")
	assert.True(t, set.Contains("hello"))

	set.Add("world")
	assert.True(t, set.Contains("world"))
}
