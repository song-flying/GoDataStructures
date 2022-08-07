package order

import "golang.org/x/exp/constraints"

type CompareFn[T comparable] func(x, y T) int

func NaturalOrdered[T constraints.Ordered]() CompareFn[T] {
	return func(x, y T) int {
		switch {
		case x < y:
			return -1
		case x > y:
			return 1
		default:
			return 0
		}
	}
}

var IntComp = NaturalOrdered[int]()
