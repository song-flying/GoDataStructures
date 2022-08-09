package tree

import "github.com/song-flying/GoDataStructures/tree"

func LinearSearch[T comparable](x T, t *tree.BinaryTree[T]) *tree.BinaryNode[T] {
	return LinearSearchFrom(x, t.Root)
}

func LinearSearchFrom[T comparable](x T, root *tree.BinaryNode[T]) *tree.BinaryNode[T] {
	if root == nil {
		return nil
	}

	if root.Data == x {
		return root
	}

	node := LinearSearchFrom(x, root.Left)
	if node != nil {
		return node
	}
	return LinearSearchFrom(x, root.Right)
}
