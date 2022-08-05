package set

type Set[T comparable] interface {
	Contains(x T) bool
	Add(x T)
	Remove(x T)
	Size() int
}
