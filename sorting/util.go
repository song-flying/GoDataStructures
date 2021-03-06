package sorting

import "github.com/song-flying/GoDataStructures/pkg/assertion"

func isSorted(a []int, low, high int) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")

	for i := low; i < high-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}

	return true
}

func swap(a []int, i, j int) {
	assertion.Require(0 <= i && i < len(a), "i is within bound")
	assertion.Require(0 <= j && j < len(a), "j is within bound")
	defer func(oldAi, oldAj int) {
		assertion.Ensure(a[i] == oldAj && a[j] == oldAi, "elements at i, j are swapped")
	}(a[i], a[j])

	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp

	return
}

type compareFn = func(x, y int) bool

func elementCompareTo(x int, a []int, low, high int, comparer compareFn) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")

	for i := low; i < high; i++ {
		if !comparer(x, a[i]) {
			return false
		}
	}

	return true
}

func elementLessThanOrEqual(x int, a []int, low, high int) bool {
	return elementCompareTo(x, a, low, high, func(x, y int) bool {
		return x <= y
	})
}

func elementLessThan(x int, a []int, low, high int) bool {
	return elementCompareTo(x, a, low, high, func(x, y int) bool {
		return x < y
	})
}

func elementGreaterThanOrEqual(x int, a []int, low, high int) bool {
	return elementCompareTo(x, a, low, high, func(x, y int) bool {
		return x >= y
	})
}

func elementGreaterThan(x int, a []int, low, high int) bool {
	return elementCompareTo(x, a, low, high, func(x, y int) bool {
		return x > y
	})
}

func rangeLessOrEqual(a []int, lowA, highA int, b []int, lowB, highB int) bool {
	assertion.Require(0 <= lowA && lowA <= highA && highA <= len(a), "lowA and highA are within bound")
	assertion.Require(0 <= lowB && lowB <= highB && highB <= len(b), "lowB and highB are within bound")

	for i := lowA; i < highA; i++ {
		if !elementLessThanOrEqual(a[i], b, lowB, highB) {
			return false
		}
	}

	return true
}
