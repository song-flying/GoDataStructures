package sorting

import (
	"github.com/song-flying/GoDataStructures/pkg/assertion"
)

func SelectionSort(a []int) {
	SelectionSortRange(a, 0, len(a))
}

func SelectionSortRange(a []int, low, high int) {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")
	defer func() {
		assertion.Ensure(isSorted(a, low, high), "a is sorted")
	}()

	for i := low; i < high; i++ {
		assertion.Invariant(low <= i && i <= high, "i is within bound")
		assertion.Invariant(isSorted(a, low, i), "a[low,i) is sorted")
		assertion.Invariant(rangeLessOrEqual(a, low, i, a, i, high), "a[low,i) <= a[i,high)")
		minIndex := findMin(a, i, high)
		swap(a, minIndex, i)
	}
}

func isSorted(a []int, low, high int) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")

	for i := low; i < high-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}

	return true
}

func elementLessOrEqual(x int, a []int, low, high int) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")

	for i := low; i < high; i++ {
		if a[i] < x {
			return false
		}
	}

	return true
}

func rangeLessOrEqual(a []int, lowA, highA int, b []int, lowB, highB int) bool {
	assertion.Require(0 <= lowA && lowA <= highA && highA <= len(a), "lowA and highA are within bound")
	assertion.Require(0 <= lowB && lowB <= highB && highB <= len(a), "lowB and highB are within bound")

	for i := lowA; i < highA; i++ {
		if !elementLessOrEqual(a[i], b, lowB, highB) {
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

func findMin(a []int, low, high int) (result int) {
	assertion.Require(0 <= low && low < high && high <= len(a), "low and high are within bound")
	defer func() {
		assertion.Ensure(low <= result && result < high, "result is within bound")
		assertion.Ensure(elementLessOrEqual(a[result], a, low, high), "result is the index of smallest element")
	}()

	minIndex := low
	for i := low + 1; i < high; i++ {
		assertion.Invariant(low+1 <= i && i <= high, "i is within bound")
		assertion.Invariant(low <= minIndex && minIndex < high, "minIndex is within bound")
		assertion.Invariant(elementLessOrEqual(a[minIndex], a, low, i), "a[minIndex] is min value for a[low,i)")
		if a[i] < a[minIndex] {
			minIndex = i
		}
	}

	return minIndex
}
