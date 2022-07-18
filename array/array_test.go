package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayBounds(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}

	assert.Equal(t, 1, a[0])
	assert.Equal(t, 5, a[4])
	assert.Equal(t, 5, len(a))
	// Go compiler enforces array index to be within bound, i.e. 0 <= index < 5 = len(a),
	// thus uncommenting following lines will cause compilation to fail
	//fmt.Printf("a[-1] = %v", a[-1])
	//fmt.Printf("a[5] = %v", a[5])
}

func TestSliceBounds(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	assert.Equal(t, 1, a[0])
	assert.Equal(t, 5, a[4])
	assert.Equal(t, 5, len(a))
	// For slice, the compiler can only enforce lower bound, as upperbound is not declared,
	// violation of upperbound will result in run time error instead
	//fmt.Printf("a[-1] = %v", a[-1]) // compile time error
	//fmt.Printf("a[5] = %v", a[5])   // run time error
}

func TestArrayAssignment(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}

	a[0] = 10
	assert.Equal(t, 10, a[0])
	a[4] = -10
	assert.Equal(t, -10, a[4])

	// similarly assignment at an out of bound index will be rejected by the compiler
	//a[-1] = 42
	//a[5] = 42
}

func TestSliceAssignment(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	a[0] = 10
	assert.Equal(t, 10, a[0])
	a[4] = -10
	assert.Equal(t, -10, a[4])

	// similarly assignment at an out of bound index will be rejected by the compiler or runtime
	//a[-1] = 42  // compile time error
	//a[5] = 42   // run time error
}

func TestArrayAliasing(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	// array assignment creates copy
	b := a
	b[0] = -1
	assert.Equal(t, 1, a[0])
}

func TestSliceAliasing(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	// slice assignment creates alias
	b := a
	b[0] = -1
	assert.Equal(t, -1, a[0])
}
