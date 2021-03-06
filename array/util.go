package array

import "github.com/song-flying/GoDataStructures/pkg/assertion"

func cubes(n int) (result []int) {
	assertion.Require(n >= 0, "n is non-negative")
	defer func() {
		assertion.Ensure(len(result) == n, "len(result) = n")
	}()

	result = make([]int, n)

	loopInv := func(i int) bool {
		assertion.Invariant(0 <= i && i <= n, "0 <= i <= n")
		return true
	}
	for i := 0; loopInv(i) && i < n; i++ {
		result[i] = i * i * i
	}

	return result
}

func copyArray(a []int, n int) (result []int) {
	assertion.Require(n == len(a), "len(a) = n")
	defer func() {
		assertion.Ensure(len(result) == n, "len(result) = n")
		assertion.Ensure(same(a, 0, n, result, 0, n), "a[0,n) = result[0,n)")
	}()

	result = make([]int, n)

	loopInv := func(i int) bool {
		assertion.Invariant(0 <= i && i <= n, "0 <= i <= n")
		assertion.Invariant(same(a, 0, i, result, 0, i), "result[0,i] = a[0,i]")
		return true
	}
	for i := 0; loopInv(i) && i < n; i++ {
		result[i] = a[i]
	}

	return
}

func same(a []int, lowA, highA int, b []int, lowB, highB int) bool {
	assertion.Require(0 <= lowA && lowA <= highA && highA <= len(a), "a's low and high within bound")
	assertion.Require(0 <= lowB && lowB <= highB && highB <= len(a), "b's low and high within bound")
	assertion.Require(highA-lowA == highB-lowB, "a and b's segment's length are the same")

	loopInv := func(i, j int) bool {
		assertion.Invariant(lowA <= i && i <= highA, "i is within bound")
		assertion.Invariant(lowB <= j && j <= highB, "j is within bound")
		assertion.Invariant(i-lowA == j-lowB, "i - lowA = j - lowB")
		return true
	}
	for i, j := lowA, lowB; loopInv(i, j) && i < highA; i, j = i+1, j+1 {
		if a[i] != b[j] {
			return false
		}
	}

	return true
}

func subArray(a []int, low, high int) (result []int) {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low & high are in range")
	defer func() {
		assertion.Ensure(same(a, low, high, result, 0, len(result)), "result[0, len) = a[low,high)")
	}()

	result = make([]int, high-low)

	loopInv := func(i, j int) bool {
		assertion.Invariant(low <= i && i <= high, "i is within bound")
		assertion.Invariant(0 <= j && j <= high-low, "j is within bound")
		assertion.Invariant(j-0 == i-low, "i and j moves at same speed")
		assertion.Invariant(same(a, low, i, result, 0, j), "result[0,j] = a[low,i]")
		return true
	}
	for i, j := low, 0; loopInv(i, j) && i < high; i++ {
		result[j] = a[i]
		j++
	}

	return
}

func copyInto(src []int, i, n int, dst []int, j int) (result int) {
	assertion.Require(0 <= n, "n >= 0")
	assertion.Require(0 <= i && i+n <= len(src), "0 <= i && i+n <= len(src)")
	assertion.Require(0 <= j && j+n <= len(dst), "0 <= j && j+n <= len(dst)")
	defer func() {
		assertion.Ensure(same(src, i, i+n, dst, j, j+n), "src[i, i+n) = dst[j,j+n)")
		assertion.Ensure(n == 0 && (result == -1) || n > 0 && (result == j+n-1), "result OK")
	}()

	if n == 0 {
		return -1
	}

	var k, l = i, j
	loopInv := func(k, l int) bool {
		assertion.Invariant(i <= k && k <= i+n, "i <= k <= i+n")
		assertion.Invariant(j <= l && l <= j+n, "j <= l <= j+n")
		assertion.Invariant(k-i == l-j, "k-i == l-j")
		assertion.Invariant(same(src, i, k, dst, j, l), "src[i,k] = dst[j,l]")
		return true
	}
	for ; loopInv(k, l) && k < i+n; k, l = k+1, l+1 {
		dst[l] = src[k]
	}

	return l - 1
}

// specification function
func isMax(maxIndex int, a []int, n int) bool {
	assertion.Require(0 <= n && n <= len(a), "0 <= n <= len(a)")
	assertion.Require(0 <= maxIndex && maxIndex < n, "maxIndex is within bound")

	max := a[maxIndex]
	for i := 0; i < n; i++ {
		if a[i] > max {
			return false
		}
	}

	return true
}

func findMax(a []int, n int) (result int) {
	assertion.Require(0 < n && n == len(a), "0 < n = len(a)")
	defer func() {
		assertion.Ensure(0 <= result && result < n, "result is within bound")
		assertion.Ensure(isMax(result, a, n), "result is index of max element")
	}()

	maxIndex := 0
	maxVal := a[0]

	loopInv := func(i int) bool {
		assertion.Invariant(1 <= i && i <= n, "i is within bound")
		assertion.Invariant(isMax(maxIndex, a, i), "maxIndex is index of max element for a[0,i)")
		return true
	}
	for i := 1; loopInv(i) && i < n; i++ {
		if a[i] > maxVal {
			maxIndex = i
			maxVal = a[i]
		}
	}

	return maxIndex
}
