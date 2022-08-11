package array

import (
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/order"
	"math/rand"
	"time"
)

func CopyArray[T comparable](a []T, n int) (result []T) {
	contract.Require(n == len(a), "len(a) = n")
	defer func() {
		contract.Ensure(len(result) == n, "len(result) = n")
		contract.Ensure(sameRange(a, 0, n, result, 0, n), "a[0,n) = result[0,n)")
	}()

	result = make([]T, n)

	loopInv := func(i int) bool {
		contract.Invariant(0 <= i && i <= n, "0 <= i <= n")
		contract.Invariant(sameRange(a, 0, i, result, 0, i), "result[0,i] = a[0,i]")
		return true
	}
	for i := 0; loopInv(i) && i < n; i++ {
		result[i] = a[i]
	}

	return
}

func Same[T comparable](a []T, b []T) bool {
	return sameRange(a, 0, len(a), b, 0, len(b))
}

func sameRange[T comparable](a []T, lowA, highA int, b []T, lowB, highB int) bool {
	contract.Require(0 <= lowA && lowA <= highA && highA <= len(a), "a's low and high within bound")
	contract.Require(0 <= lowB && lowB <= highB && highB <= len(a), "b's low and high within bound")
	contract.Require(highA-lowA == highB-lowB, "a and b's segment's length are the same")

	loopInv := func(i, j int) bool {
		contract.Invariant(lowA <= i && i <= highA, "i is within bound")
		contract.Invariant(lowB <= j && j <= highB, "j is within bound")
		contract.Invariant(i-lowA == j-lowB, "i - lowA = j - lowB")
		return true
	}
	for i, j := lowA, lowB; loopInv(i, j) && i < highA; i, j = i+1, j+1 {
		if a[i] != b[j] {
			return false
		}
	}

	return true
}

func SubArray[T comparable](a []T, low, high int) (result []T) {
	contract.Require(0 <= low && low <= high && high <= len(a), "low & high are in range")
	defer func() {
		contract.Ensure(sameRange(a, low, high, result, 0, len(result)), "result[0, len) = a[low,high)")
	}()

	result = make([]T, high-low)

	loopInv := func(i, j int) bool {
		contract.Invariant(low <= i && i <= high, "i is within bound")
		contract.Invariant(0 <= j && j <= high-low, "j is within bound")
		contract.Invariant(j-0 == i-low, "i and j moves at same speed")
		contract.Invariant(sameRange(a, low, i, result, 0, j), "result[0,j] = a[low,i]")
		return true
	}
	for i, j := low, 0; loopInv(i, j) && i < high; i++ {
		result[j] = a[i]
		j++
	}

	return
}

func CopyInto[T comparable](src []T, i, n int, dst []T, j int) (result int) {
	contract.Require(0 <= n, "n >= 0")
	contract.Require(0 <= i && i+n <= len(src), "0 <= i && i+n <= len(src)")
	contract.Require(0 <= j && j+n <= len(dst), "0 <= j && j+n <= len(dst)")
	defer func() {
		contract.Ensure(sameRange(src, i, i+n, dst, j, j+n), "src[i, i+n) = dst[j,j+n)")
		contract.Ensure(n == 0 && (result == -1) || n > 0 && (result == j+n-1), "result OK")
	}()

	if n == 0 {
		return -1
	}

	var k, l = i, j
	loopInv := func(k, l int) bool {
		contract.Invariant(i <= k && k <= i+n, "i <= k <= i+n")
		contract.Invariant(j <= l && l <= j+n, "j <= l <= j+n")
		contract.Invariant(k-i == l-j, "k-i == l-j")
		contract.Invariant(sameRange(src, i, k, dst, j, l), "src[i,k] = dst[j,l]")
		return true
	}
	for ; loopInv(k, l) && k < i+n; k, l = k+1, l+1 {
		dst[l] = src[k]
	}

	return l - 1
}

// specification function
func isMax[T comparable](maxIndex int, a []T, n int, comp order.CompareFn[T]) bool {
	contract.Require(0 <= n && n <= len(a), "0 <= n <= len(a)")
	contract.Require(0 <= maxIndex && maxIndex < n, "maxIndex is within bound")

	max := a[maxIndex]
	for i := 0; i < n; i++ {
		if comp(a[i], max) > 0 {
			return false
		}
	}

	return true
}

func FindMax[T comparable](a []T, n int, comp order.CompareFn[T]) (result int) {
	contract.Require(0 < n && n == len(a), "0 < n = len(a)")
	defer func() {
		contract.Ensure(0 <= result && result < n, "result is within bound")
		contract.Ensure(isMax(result, a, n, comp), "result is index of max element")
	}()

	maxIndex := 0
	maxVal := a[0]

	loopInv := func(i int) bool {
		contract.Invariant(1 <= i && i <= n, "i is within bound")
		contract.Invariant(isMax(maxIndex, a, i, comp), "maxIndex is index of max element for a[0,i)")
		return true
	}
	for i := 1; loopInv(i) && i < n; i++ {
		if comp(a[i], maxVal) > 0 {
			maxIndex = i
			maxVal = a[i]
		}
	}

	return maxIndex
}

func swap[T comparable](a []T, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func Shuffle[T comparable](a []T) {
	for i := 0; i < len(a); i++ {
		swap(a, i, randRange(i, len(a)))
	}
}

func randRange(m, n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(n-m) + m
}

func IsDistinct[T comparable](a []T, comp order.CompareFn[T]) bool {
	contract.Require(IsSorted(a, comp), "a is sorted")

	if len(a) <= 1 {
		return true
	}

	for prev, curr := 0, 1; curr < len(a); prev, curr = prev+1, curr+1 {
		if prev == curr {
			return false
		}
	}

	return true
}

func Contains[T comparable](x T, a []T, low, high int) bool {
	for i := low; i < high; i++ {
		if x == a[i] {
			return true
		}
	}

	return false
}

func ContainsBy[T comparable](x T, a []T, comp order.CompareFn[T]) bool {
	for _, y := range a {
		if comp(x, y) == 0 {
			return true
		}
	}

	return false
}

func Filter[T comparable](x T, a []T, comp order.CompareFn[T]) []T {
	var b []T
	for _, y := range a {
		if comp(x, y) != 0 {
			b = append(b, y)
		}
	}

	return b
}

// IsIn specification function
func IsIn[T comparable](x T, a []T, low, high int) bool {
	contract.Require(0 <= low && low <= high && high <= len(a), "low and high are within bounds")

	loopInv := func(i int) bool {
		contract.Invariant(low <= i && i <= high, "i is within bound")
		return true
	}
	for i := low; loopInv(i) && i < high; i++ {
		if x == a[i] {
			return true
		}
	}

	return false
}

func IsSorted[T comparable](a []T, comp order.CompareFn[T]) bool {
	return IsRangeSorted(a, 0, len(a), comp)
}

// IsRangeSorted specification function
func IsRangeSorted[T comparable](a []T, low, high int, comp order.CompareFn[T]) bool {
	contract.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")

	loopInv := func(i int) bool {
		contract.Invariant(low <= i, "i is within lower bound")
		return true
	}
	for i := low; loopInv(i) && i < high-1; i++ {
		contract.Assert(i < high-1, "i is within upper bound")
		if comp(a[i], a[i+1]) > 0 {
			return false
		}
	}

	return true
}
