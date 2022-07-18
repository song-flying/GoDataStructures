package array

import "github.com/song-flying/GoDataStructures/pkg/assertion"

func cubes(n int) (result []int) {
	assertion.Requiref(n >= 0, "n=%d must be non-negative", n)
	defer func() {
		assertion.Ensuref(len(result) == n, "result length %d should be equal to %d", len(result), n)
	}()

	result = make([]int, n)

	for i := 0; i < n; i++ {
		assertion.Invariantf(0 <= i && i < n, "i=%d is out of array bound", i)
		result[i] = i * i * i
	}

	return result
}
