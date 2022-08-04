package set

type Set[T comparable] interface {
	Contains(x T) bool
	Add(x T)
}
