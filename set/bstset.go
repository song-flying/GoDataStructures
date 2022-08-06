package set

import (
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"github.com/song-flying/GoDataStructures/tree"
	"golang.org/x/exp/constraints"
)

type CompareFn[E constraints.Ordered] func(x, y E) int

type BSTSet[E constraints.Ordered] struct {
	tree *tree.BinaryTree[E]
	comp CompareFn[E]
	size int
}

func (b *BSTSet[E]) isOrdered(root *tree.BinaryNode[E], lower, upper *E) bool {
	assertion.Require(b.comp != nil, "comparison function is not nil")
	assertion.Require(lower == nil || upper == nil || b.comp(*lower, *upper) < 0, "lower < upper")

	if root == nil {
		return true
	}

	key := root.Data

	return (lower == nil || b.comp(*lower, key) < 0) && b.isOrdered(root.Left, lower, &key) &&
		(upper == nil || b.comp(key, *upper) < 0) && b.isOrdered(root.Right, &key, upper)
}

// IsBSTSet data structure invariant
func (b *BSTSet[E]) IsBSTSet() bool {
	return b.tree != nil && b.comp != nil && b.tree.IsBinaryTree() && b.isOrdered(b.tree.Root, nil, nil)
}

func NewBSTSet[E constraints.Ordered](comp CompareFn[E]) (result BSTSet[E]) {
	assertion.Require(comp != nil, "comparison function is not nil")
	defer func() {
		assertion.Ensure(result.IsBSTSet(), "BST invariant holds")
	}()

	t := tree.NewBinaryTree(tree.Nil[E]())
	return BSTSet[E]{
		tree: &t,
		comp: comp,
		size: 0,
	}
}

func (b *BSTSet[E]) lookup(root **tree.BinaryNode[E], key E) **tree.BinaryNode[E] {
	assertion.Require(b.IsBSTSet(), "BST invariant holds")

	if *root == nil {
		return nil
	}

	compResult := b.comp(key, (*root).Data)
	switch {
	case compResult == 0:
		return root
	case compResult < 0:
		return b.lookup(&(*root).Left, key)
	default: //compResult > 0:
		return b.lookup(&(*root).Right, key)
	}
}

func (b *BSTSet[E]) Contains(key E) bool {
	assertion.Require(b.IsBSTSet(), "BST invariant holds")

	node := b.lookup(&b.tree.Root, key)
	return node != nil
}

func (b *BSTSet[E]) insert(root *tree.BinaryNode[E], key E) *tree.BinaryNode[E] {
	assertion.Require(root.IsBinaryTree(), "root is valid binary tree")
	assertion.Require(b.isOrdered(root, nil, nil), "root is ordered")
	defer func() {
		assertion.Ensure(root.IsBinaryTree(), "root is valid binary tree")
		assertion.Ensure(b.isOrdered(root, nil, nil), "root is ordered")
	}()

	if root == nil {
		node := tree.NewBinaryNode[E](key)
		b.size++
		return &node
	}

	compResult := b.comp(key, root.Data)
	switch {
	case compResult == 0:
		root.Data = key
	case compResult < 0:
		root.Left = b.insert(root.Left, key)
	default: //compResult > 0:
		root.Right = b.insert(root.Right, key)
	}

	return root
}

func (b *BSTSet[E]) Add(key E) {
	assertion.Require(b.IsBSTSet(), "BST invariant holds")
	defer func() {
		assertion.Ensure(b.IsBSTSet(), "BST invariant holds")
		assertion.Ensure(b.Contains(key), "Get(key) returns value")
	}()

	b.tree.Root = b.insert(b.tree.Root, key)
}

func (b *BSTSet[E]) remove(root **tree.BinaryNode[E]) {
	assertion.Require(root != nil, "root node is not nil")
	assertion.Require(root != nil, "root is not nil")
	assertion.Require((*root).IsBinaryTree(), "root is valid binary tree")
	assertion.Require(b.isOrdered(*root, nil, nil), "root is ordered")
	defer func() {
		assertion.Ensure((*root).IsBinaryTree(), "root is valid binary tree")
		assertion.Ensure(b.isOrdered(*root, nil, nil), "root is ordered")
	}()

	switch {
	case (*root).Left == nil && (*root).Right == nil:
		*root = nil
		return
	case (*root).Left != nil:
		curr := &(*root).Left
		for (*curr).Right != nil {
			curr = &(*curr).Right
		}
		(*root).Data = (*curr).Data
		*curr = (*curr).Left
	case (*root).Right != nil:
		curr := &(*root).Right
		for (*curr).Left != nil {
			curr = &(*curr).Left
		}
		(*root).Data = (*curr).Data
		*curr = (*curr).Right
	}
}

func (b *BSTSet[E]) Delete(key E) {
	assertion.Require(b.IsBSTSet(), "BST invariant holds")
	defer func() {
		assertion.Ensure(b.IsBSTSet(), "BST invariant holds")
		assertion.Ensure(!b.Contains(key), "Get(key) returns no value")
	}()

	target := b.lookup(&b.tree.Root, key)
	if target == nil {
		return
	}

	b.remove(target)
	b.size--
}

func (b *BSTSet[E]) Size() (result int) {
	assertion.Require(b.IsBSTSet(), "BST invariant holds")
	defer func() {
		assertion.Ensure(0 <= result, "result is non-negative")
	}()

	return b.size
}
