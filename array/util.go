package array

import "github.com/song-flying/GoDataStructures/pkg/assertion"

func cubes(n int) (result []int) {
	assertion.Requiref(n >= 0, "n (=%d) is non-negative", n)
	defer func() {
		assertion.Ensuref(len(result) == n, "result's length (=%d) equals n (=%d)", len(result), n)
	}()

	result = make([]int, n)

	for i := 0; i < n; i++ {
		assertion.Invariant(0 <= i && i <= n, "0 <= i <= n")
		result[i] = i * i * i
	}

	return result
}

func copyArray(a []int, n int) (result []int) {
	assertion.Require(n == len(a), "len(a) = n")
	defer func() {
		assertion.Ensure(len(result) == n, "len(result) = n")
	}()
	result = make([]int, n)
	for i := 0; i < n; i++ {
		assertion.Invariant(0 <= i && i < n, "0 <= i < n")
		result[i] = a[i]
	}

	return
}

func subArray(a []int, low, high int) (result []int) {
	assertion.Require(0 <= low && low <= high && high <= len(a), "0 <= low <= high <= len(a)")
	defer func() {
		assertion.Ensure(len(result) == high-low, "len(result) = high - low")
	}()

	result = make([]int, high-low)
	for i, j := low, 0; i < high; i++ {
		assertion.Invariant(low <= i && i < high, "low <= i < high")
		assertion.Invariant(0 <= j && j < high-low, "0 <= j < high - low")
		assertion.Invariant(j-0 == i-low, "j - 0 = i - low")
		result[j] = a[i]
		j++
	}

	return
}

func copyInto(src []int, i, n int, dst []int, j int) int {
	assertion.Require(0 <= i && 0 <= n && i+n <= len(src), "0 <= i && 0 <= n && i+n <= len(src)")
	assertion.Require(0 <= j && j+n <= len(dst), "0 <= j && j + n <= len(dst)")

	var k, l = i, j
	for ; k < i+n; k, l = k+1, l+1 {
		assertion.Invariant(i <= k && k < i+n, "i <= k < i+n")
		assertion.Invariant(j <= l && l < j+n, "j <= l < j+n")
		assertion.Invariant(k-i == l-j, "k-i == l-j")
		dst[l] = src[k]
	}

	return k
}
