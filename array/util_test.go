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
