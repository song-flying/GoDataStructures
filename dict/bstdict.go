package dict

import (
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"github.com/song-flying/GoDataStructures/tree"
	"golang.org/x/exp/constraints"
)

type CompareFn[T constraints.Ordered] func(x, y T) int

type BSTDict[K constraints.Ordered, V comparable] struct {
	tree *tree.BinaryTree[entry[K, V]]
	comp CompareFn[K]
	size int
}

func (b *BSTDict[K, V]) isOrdered(root *tree.BinaryNode[entry[K, V]], lower, upper *K) bool {
	assertion.Require(b.comp != nil, "comparison function is not nil")
	assertion.Require(lower == nil || upper == nil || b.comp(*lower, *upper) < 0, "lower < upper")

	if root == nil {
		return true
	}

	key := root.Data.Key

	return (lower == nil || b.comp(*lower, key) < 0) && b.isOrdered(root.Left, lower, &key) &&
		(upper == nil || b.comp(key, *upper) < 0) && b.isOrdered(root.Right, &key, upper)
}

// IsBSTDict data structure invariant
func (b *BSTDict[K, V]) IsBSTDict() bool {
	return b.tree != nil && b.comp != nil && b.tree.IsBinaryTree() && b.isOrdered(b.tree.Root, nil, nil)
}

func NewBSTDict[K constraints.Ordered, V comparable](comp CompareFn[K]) (result BSTDict[K, V]) {
	assertion.Require(comp != nil, "comparison function is not nil")
	defer func() {
		assertion.Ensure(result.IsBSTDict(), "BST invariant holds")
	}()

	t := tree.NewBinaryTree(tree.Nil[entry[K, V]]())
	return BSTDict[K, V]{
		tree: &t,
		comp: comp,
		size: 0,
	}
}

func (b *BSTDict[K, V]) lookup(root **tree.BinaryNode[entry[K, V]], key K) **tree.BinaryNode[entry[K, V]] {
	assertion.Require(b.IsBSTDict(), "BST invariant holds")

	if *root == nil {
		return nil
	}

	compResult := b.comp(key, (*root).Data.Key)
	switch {
	case compResult == 0:
		return root
	case compResult < 0:
		return b.lookup(&(*root).Left, key)
	default: //compResult > 0:
		return b.lookup(&(*root).Right, key)
	}
}

func (b *BSTDict[K, V]) Get(key K) (V, bool) {
	assertion.Require(b.IsBSTDict(), "BST invariant holds")

	node := b.lookup(&b.tree.Root, key)
	if node != nil {
		return (*node).Data.Value, true
	}
	return *new(V), false
}

func (b *BSTDict[K, V]) insert(root *tree.BinaryNode[entry[K, V]], key K, value V) *tree.BinaryNode[entry[K, V]] {
	assertion.Require(root.IsBinaryTree(), "root is valid binary tree")
	assertion.Require(b.isOrdered(root, nil, nil), "root is ordered")
	defer func() {
		assertion.Ensure(root.IsBinaryTree(), "root is valid binary tree")
		assertion.Ensure(b.isOrdered(root, nil, nil), "root is ordered")
	}()

	if root == nil {
		node := tree.NewBinaryNode[entry[K, V]](entry[K, V]{Key: key, Value: value})
		b.size++
		return &node
	}

	compResult := b.comp(key, root.Data.Key)
	switch {
	case compResult == 0:
		root.Data.Value = value
	case compResult < 0:
		root.Left = b.insert(root.Left, key, value)
	default: //compResult > 0:
		root.Right = b.insert(root.Right, key, value)
	}

	return root
}

func (b *BSTDict[K, V]) Put(key K, value V) {
	assertion.Require(b.IsBSTDict(), "BST invariant holds")
	defer func() {
		assertion.Ensure(b.IsBSTDict(), "BST invariant holds")
		v, ok := b.Get(key)
		assertion.Ensure(ok && value == v, "Get(key) returns value")
	}()

	b.tree.Root = b.insert(b.tree.Root, key, value)
}

func (b *BSTDict[K, V]) remove(root **tree.BinaryNode[entry[K, V]]) {
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

func (b *BSTDict[K, V]) Delete(key K) {
	assertion.Require(b.IsBSTDict(), "BST invariant holds")
	defer func() {
		assertion.Ensure(b.IsBSTDict(), "BST invariant holds")
		_, ok := b.Get(key)
		assertion.Ensure(!ok, "Get(key) returns no value")
	}()

	target := b.lookup(&b.tree.Root, key)
	if target == nil {
		return
	}

	b.remove(target)
	b.size--
}

func (b *BSTDict[K, V]) Size() (result int) {
	assertion.Require(b.IsBSTDict(), "BST invariant holds")
	defer func() {
		assertion.Ensure(0 <= result, "result is non-negative")
	}()

	return b.size
}
