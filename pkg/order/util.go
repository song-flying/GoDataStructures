package order

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](x, y T) T {
	if x < y {
		return y
	}

	return x
}
