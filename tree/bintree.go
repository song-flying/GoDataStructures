package tree

import (
	"encoding/json"
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"github.com/song-flying/GoDataStructures/queue"
)

type BinaryNode[T comparable] struct {
	id    int
	Data  T
	Left  *BinaryNode[T] `json:",omitempty"`
	Right *BinaryNode[T] `json:",omitempty"`
}

func NewBinaryNode[T comparable](data T) BinaryNode[T] {
	return BinaryNode[T]{
		Data: data,
	}
}

func Nil[T comparable]() *BinaryNode[T] {
	return nil
}

// IsBinaryTree data structure invariant
func (n *BinaryNode[T]) IsBinaryTree() bool {
	return !hasCycle(n)
}

type BinaryTree[T comparable] struct {
	Root *BinaryNode[T]
}

func hasCycle[T comparable](root *BinaryNode[T]) bool {
	var visited []*BinaryNode[T]
	defer func() {
		for _, node := range visited {
			node.id = 0
		}
	}()

	if root == nil {
		return false
	}

	q := queue.NewLinkedQueue[*BinaryNode[T]]()
	root.id = 1
	q.Enqueue(root)

	for !q.IsEmpty() {
		node := q.Dequeue()
		visited = append(visited, node)

		if node.Left != nil {
			if node.Left.id != 0 && node.Left.id != node.id*2 {
				return true
			}
			node.Left.id = node.id * 2
			q.Enqueue(node.Left)
		}

		if node.Right != nil {
			if node.Right.id != 0 && node.Right.id != node.id*2+1 {
				return true
			}
			node.Right.id = node.id*2 + 1
			q.Enqueue(node.Right)
		}
	}

	return false
}

// IsBinaryTree data structure invariant
func (t *BinaryTree[T]) IsBinaryTree() bool {
	return !hasCycle(t.Root)
}

func NewBinaryTree[T comparable](root *BinaryNode[T]) (result BinaryTree[T]) {
	defer func() {
		assertion.Ensure(result.IsBinaryTree(), "binary tree invariant holds")
	}()

	return BinaryTree[T]{
		Root: root,
	}
}

func (t *BinaryTree[T]) String() string {
	treeJson, _ := json.Marshal(t)
	return string(treeJson)
}
