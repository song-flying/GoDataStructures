package stack

type Stack[T any] interface {
	IsEmpty() bool
	Push(x T)
	Pop() T
}

type Factory[T any] interface {
	NewStack() Stack[T]
}
