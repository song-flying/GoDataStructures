package dict

import (
	"github.com/stretchr/testify/assert"
	"hash/maphash"
	"testing"
)

func TestHashDict(t *testing.T) {
	var h maphash.Hash
	hashFn := func(s string) int {
		_, _ = h.WriteString(s)
		defer h.Reset()
		return int(h.Sum64())
	}

	dict := NewHashDict[string, int](10, hashFn)
	v, ok := dict.Get("hello")
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, v)

	dict.Put("hello", 1)
	v, ok = dict.Get("hello")
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, v)

	dict.Put("world", 2)
	v, ok = dict.Get("world")
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, v)

	dict.Put("hello", 3)
	v, ok = dict.Get("hello")
	assert.Equal(t, true, ok)
	assert.Equal(t, 3, v)
}
