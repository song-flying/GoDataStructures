package tree

import (
	"encoding/json"
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"github.com/song-flying/GoDataStructures/queue"
	"github.com/song-flying/GoDataStructures/stack"
	"strconv"
)

var (
	unvisited     = 0
	visitingLeft  = 1
	visitingRight = 2
	visited       = 3
)

type BinaryNode[T comparable] struct {
	id    int
	state int
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

func (n *BinaryNode[T]) ToArray() (result []T) {
	assertion.Require(n.IsBinaryTree(), "n is valid binary tree")

	if n == nil {
		return
	}

	if n.Left != nil {
		result = append(result, n.Left.ToArray()...)
	}
	result = append(result, n.Data)
	if n.Right != nil {
		result = append(result, n.Right.ToArray()...)
	}

	return
}

func (n *BinaryNode[T]) ToArrayPreorder() (result []T) {
	assertion.Require(n.IsBinaryTree(), "n is valid binary tree")

	var nodesVisited []*BinaryNode[T]
	defer func() {
		for _, node := range nodesVisited {
			node.state = unvisited
		}
	}()

	if n == nil {
		return
	}

	s := stack.NewLinkedStack[*BinaryNode[T]]()
	s.Push(n)

	for !s.IsEmpty() {
		node := s.Peek()
		switch node.state {
		case unvisited:
			if node.Left != nil {
				s.Push(node.Left)
				node.state = visitingLeft
				break
			}
			nodesVisited = append(nodesVisited, node)
			result = append(result, node.Data)
			node.state = visited
			s.Pop()
			if node.Right != nil {
				s.Push(node.Right)
			}
		case visitingLeft:
			nodesVisited = append(nodesVisited, node)
			result = append(result, node.Data)
			node.state = visited
			s.Pop()
			if node.Right != nil {
				s.Push(node.Right)
			}
		default:
			panic("unexpected state " + strconv.Itoa(node.state))
		}
	}

	return
}

func (n *BinaryNode[T]) ToArrayInorder() (result []T) {
	assertion.Require(n.IsBinaryTree(), "n is valid binary tree")

	var nodesVisited []*BinaryNode[T]
	defer func() {
		for _, node := range nodesVisited {
			node.state = unvisited
		}
	}()

	if n == nil {
		return
	}

	s := stack.NewLinkedStack[*BinaryNode[T]]()
	s.Push(n)

	for !s.IsEmpty() {
		node := s.Peek()
		switch node.state {
		case unvisited:
			nodesVisited = append(nodesVisited, node)
			result = append(result, node.Data)
			node.state = visited
			s.Pop()
			if node.Right != nil {
				s.Push(node.Right)
			}
			if node.Left != nil {
				s.Push(node.Left)
			}
		default:
			panic("unexpected state " + strconv.Itoa(node.state))
		}
	}

	return
}

func (n *BinaryNode[T]) ToArrayPostorder() (result []T) {
	assertion.Require(n.IsBinaryTree(), "n is valid binary tree")

	var nodesVisited []*BinaryNode[T]
	defer func() {
		for _, node := range nodesVisited {
			node.state = unvisited
		}
	}()

	if n == nil {
		return
	}

	s := stack.NewLinkedStack[*BinaryNode[T]]()
	s.Push(n)

	for !s.IsEmpty() {
		node := s.Peek()
		switch node.state {
		case unvisited:
			if node.Left != nil {
				s.Push(node.Left)
				node.state = visitingLeft
				break
			}
			if node.Right != nil {
				s.Push(node.Right)
				node.state = visitingRight
				break
			}
			nodesVisited = append(nodesVisited, node)
			result = append(result, node.Data)
			node.state = visited
			s.Pop()
		case visitingLeft:
			if node.Right != nil {
				s.Push(node.Right)
				node.state = visitingRight
				break
			}
			nodesVisited = append(nodesVisited, node)
			result = append(result, node.Data)
			node.state = visited
			s.Pop()
		case visitingRight:
			nodesVisited = append(nodesVisited, node)
			result = append(result, node.Data)
			node.state = visited
			s.Pop()
		default:
			panic("unexpected state " + strconv.Itoa(node.state))
		}
	}

	return
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
	return t.Root.IsBinaryTree()
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
