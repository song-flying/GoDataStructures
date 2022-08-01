package linked

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type List[T any] *Node[T]

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		Data: data,
		Next: nil,
	}
}

func NewDummyNode[T any]() *Node[T] {
	return &Node[T]{}
}

func Nil[T any]() *Node[T] {
	return nil
}
