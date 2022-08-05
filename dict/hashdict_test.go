package dict

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

func TestHashDict(t *testing.T) {
	dict := NewHashDict[string, int](1, hashFn, 1)
	v, ok := dict.Get("hello")
	assert.False(t, ok)
	assert.Equal(t, 0, v)
	assert.Equal(t, 0, dict.Size())

	dict.Put("hello", 1)
	v, ok = dict.Get("hello")
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	assert.Equal(t, 1, dict.Size())

	dict.Put("world", 2)
	v, ok = dict.Get("world")
	assert.True(t, ok)
	assert.Equal(t, 2, v)
	assert.Equal(t, 2, dict.Size())

	dict.Put("hello", 3)
	v, ok = dict.Get("hello")
	assert.True(t, ok)
	assert.Equal(t, 3, v)
	assert.Equal(t, 2, dict.Size())

	dict.Delete("world")
	v, ok = dict.Get("world")
	assert.False(t, ok)
	assert.Equal(t, 1, dict.Size())

	dict.Delete("hello")
	assert.Equal(t, 0, dict.Size())

	for i := 0; i < 8; i++ {
		dict.Put(strconv.Itoa(i), i)
	}
	assert.Equal(t, 16, dict.capacity)

	for i := 0; i < 6; i++ {
		dict.Delete(strconv.Itoa(i))
	}
	assert.Equal(t, 4, dict.capacity)
}
