package dict

import (
	"github.com/song-flying/GoDataStructures/array"
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

func entryComparator[K constraints.Ordered, V comparable](x, y entry[K, V]) int {
	switch {
	case x.Key == y.Key:
		return 0
	case x.Key < y.Key:
		return -1
	default: //x > y
		return 1
	}
}

func (b *BSTDict[K, V]) hasSameEntries(a1, a2 []entry[K, V]) bool {
	return len(a1) == len(a2) && array.IsSubArrayOf(a1, a2, entryComparator[K, V])
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

func equalKey[K comparable, V comparable](e1, e2 entry[K, V]) bool {
	return e1.Key == e2.Key
}

func (b *BSTDict[K, V]) insert(root *tree.BinaryNode[entry[K, V]], key K, value V) (result *tree.BinaryNode[entry[K, V]]) {
	assertion.Require(root.IsBinaryTree(), "root is valid binary tree")
	assertion.Require(b.isOrdered(root, nil, nil), "root is ordered")
	defer func(oldRootEntries []entry[K, V]) {
		assertion.Ensure(result.IsBinaryTree(), "new root is valid binary tree")
		assertion.Ensure(b.isOrdered(result, nil, nil), "new root is ordered")
		e := entry[K, V]{Key: key, Value: value}
		newRootEntries := result.ToArray()
		assertion.Ensure(array.Contains(e, newRootEntries), "new root should contain new entry")
		oldRootEntries = array.Remove(e, oldRootEntries, equalKey[K, V])
		newRootEntries = array.Remove(e, newRootEntries, equalKey[K, V])
		assertion.Ensure(b.hasSameEntries(oldRootEntries, newRootEntries), "new root should contain same entries as old root, except for new entry")
	}(root.ToArray())

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

func (b *BSTDict[K, V]) remove(pRoot **tree.BinaryNode[entry[K, V]]) {
	assertion.Require(pRoot != nil, "root is not nil")
	assertion.Require((*pRoot).IsBinaryTree(), "root is valid binary tree")
	assertion.Require(b.isOrdered(*pRoot, nil, nil), "root is ordered")
	defer func() {
		assertion.Ensure((*pRoot).IsBinaryTree(), "root is valid binary tree")
		assertion.Ensure(b.isOrdered(*pRoot, nil, nil), "root is ordered")
	}()

	switch {
	case (*pRoot).Left == nil && (*pRoot).Right == nil:
		*pRoot = nil
		return
	case (*pRoot).Left != nil:
		pCurr := &(*pRoot).Left
		for (*pCurr).Right != nil {
			pCurr = &(*pCurr).Right
		}
		(*pRoot).Data = (*pCurr).Data
		*pCurr = (*pCurr).Left
	case (*pRoot).Right != nil:
		pCurr := &(*pRoot).Right
		for (*pCurr).Left != nil {
			pCurr = &(*pCurr).Left
		}
		(*pRoot).Data = (*pCurr).Data
		*pCurr = (*pCurr).Right
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
