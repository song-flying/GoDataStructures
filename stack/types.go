package stack

type Stack[T any] interface {
	IsEmpty() bool
	Push(x T)
	Pop() T
	Peek() T
}
