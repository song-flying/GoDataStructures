package sorting

import (
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"github.com/song-flying/GoDataStructures/pkg/order"
)

func isSorted[T comparable](a []T, low, high int, comp order.CompareFn[T]) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")

	for i := low; i < high-1; i++ {
		if comp(a[i], a[i+1]) > 0 {
			return false
		}
	}

	return true
}

func swap[T comparable](a []T, i, j int) {
	assertion.Require(0 <= i && i < len(a), "i is within bound")
	assertion.Require(0 <= j && j < len(a), "j is within bound")
	defer func(oldAi, oldAj T) {
		assertion.Ensure(a[i] == oldAj && a[j] == oldAi, "elements at i, j are swapped")
	}(a[i], a[j])

	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp

	return
}

func elementLessThanOrEqual[T comparable](x T, a []T, low, high int, comp order.CompareFn[T]) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")

	for i := low; i < high; i++ {
		if comp(x, a[i]) > 0 {
			return false
		}
	}

	return true
}

func elementGreaterThanOrEqual[T comparable](x T, a []T, low, high int, comp order.CompareFn[T]) bool {
	assertion.Require(0 <= low && low <= high && high <= len(a), "low and high are within bound")

	for i := low; i < high; i++ {
		if comp(x, a[i]) < 0 {
			return false
		}
	}

	return true
}

func rangeLessOrEqual[T comparable](a []T, lowA, highA int, b []T, lowB, highB int, comp order.CompareFn[T]) bool {
	assertion.Require(0 <= lowA && lowA <= highA && highA <= len(a), "lowA and highA are within bound")
	assertion.Require(0 <= lowB && lowB <= highB && highB <= len(b), "lowB and highB are within bound")

	for i := lowA; i < highA; i++ {
		if !elementLessThanOrEqual(a[i], b, lowB, highB, comp) {
			return false
		}
	}

	return true
}
