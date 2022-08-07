package dict

import (
	"github.com/song-flying/GoDataStructures/array"
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/song-flying/GoDataStructures/tree"
)

type BSTDict[K comparable, V comparable] struct {
	tree      *tree.BinaryTree[entry[K, V]]
	keyComp   order.CompareFn[K]
	entryComp order.CompareFn[entry[K, V]]
	size      int
}

func (b *BSTDict[K, V]) isOrdered(root *tree.BinaryNode[entry[K, V]], lower, upper *K) bool {
	assertion.Require(b.keyComp != nil, "comparison function is not nil")
	assertion.Require(lower == nil || upper == nil || b.keyComp(*lower, *upper) < 0, "lower < upper")

	if root == nil {
		return true
	}

	key := root.Data.Key

	return (lower == nil || b.keyComp(*lower, key) < 0) && b.isOrdered(root.Left, lower, &key) &&
		(upper == nil || b.keyComp(key, *upper) < 0) && b.isOrdered(root.Right, &key, upper)
}

func (b *BSTDict[K, V]) isOrderedHasMinMax(root *tree.BinaryNode[entry[K, V]]) (minKey, maxKey *K, isOrdered bool) {
	assertion.Require(b.keyComp != nil, "comparison function is not nil")
	assertion.Require(root.IsBinaryTree(), "root is a binary tree")

	if root == nil {
		return nil, nil, true
	}

	if root.Left != nil {
		leftMinKey, leftMaxKey, leftIsOrdered := b.isOrderedHasMinMax(root.Left)
		if !leftIsOrdered || b.keyComp(*leftMaxKey, root.Data.Key) >= 0 {
			return nil, nil, false
		} else {
			minKey = leftMinKey
		}
	} else {
		minKey = &root.Data.Key
	}

	if root.Right != nil {
		rightMinKey, rightMaxKey, rightIsOrdered := b.isOrderedHasMinMax(root.Right)
		if !rightIsOrdered || b.keyComp(*rightMinKey, root.Data.Key) <= 0 {
			return nil, nil, false
		} else {
			maxKey = rightMaxKey
		}
	} else {
		maxKey = &root.Data.Key
	}

	return minKey, maxKey, true
}

func (b *BSTDict[K, V]) hasSameEntries(a1, a2 []entry[K, V]) bool {
	assertion.Require(array.IsSorted(a1, b.entryComp) && array.IsDistinct(a1, b.entryComp), "a1 is sorted & distinct")
	assertion.Require(array.IsSorted(a2, b.entryComp) && array.IsDistinct(a2, b.entryComp), "a2 is sorted & distinct")

	return array.Same(a1, a2)
}

// IsBSTDict data structure invariant
func (b *BSTDict[K, V]) IsBSTDict() bool {
	_, _, isOrdered := b.isOrderedHasMinMax(b.tree.Root)
	return b.tree != nil && b.keyComp != nil && b.entryComp != nil && b.tree.IsBinaryTree() && isOrdered
}

func NewBSTDict[K comparable, V comparable](comp order.CompareFn[K]) (result BSTDict[K, V]) {
	assertion.Require(comp != nil, "comparison function is not nil")
	defer func() {
		assertion.Ensure(result.IsBSTDict(), "BST invariant holds")
	}()

	t := tree.NewBinaryTree(tree.Nil[entry[K, V]]())
	entryComp := func(e1, e2 entry[K, V]) int {
		return comp(e1.Key, e2.Key)
	}

	return BSTDict[K, V]{
		tree:      &t,
		keyComp:   comp,
		entryComp: entryComp,
		size:      0,
	}
}

func (b *BSTDict[K, V]) lookup(root **tree.BinaryNode[entry[K, V]], key K) **tree.BinaryNode[entry[K, V]] {
	assertion.Require(b.IsBSTDict(), "BST invariant holds")

	if *root == nil {
		return nil
	}

	compResult := b.keyComp(key, (*root).Data.Key)
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

func (b *BSTDict[K, V]) ToArray(root *tree.BinaryNode[entry[K, V]]) (result []entry[K, V]) {
	assertion.Require(b.IsBSTDict(), "BST invariant holds")
	defer func() {
		assertion.Ensure(array.IsSorted(result, b.entryComp), "result entries are sorted")
	}()
	if root == nil {
		return
	}

	if root.Left != nil {
		result = append(result, b.ToArray(root.Left)...)
	}
	result = append(result, root.Data)
	if root.Right != nil {
		result = append(result, b.ToArray(root.Right)...)
	}

	return
}

func (b *BSTDict[K, V]) insert(root *tree.BinaryNode[entry[K, V]], key K, value V) (result *tree.BinaryNode[entry[K, V]]) {
	assertion.Require(root.IsBinaryTree(), "root is valid binary tree")
	_, _, isOrdered := b.isOrderedHasMinMax(root)
	assertion.Require(isOrdered, "root tree is ordered")
	defer func(oldEntries []entry[K, V]) {
		assertion.Ensure(result.IsBinaryTree(), "new root is valid binary tree")
		_, _, isOrdered := b.isOrderedHasMinMax(result)
		assertion.Ensure(isOrdered, "new root is ordered")
		e := entry[K, V]{Key: key, Value: value}
		newEntries := b.ToArray(result)
		assertion.Ensure(array.Contains(e, newEntries, b.entryComp), "new root should contain new entry")
		oldEntries = array.Filter(e, oldEntries, b.entryComp)
		newEntries = array.Filter(e, newEntries, b.entryComp)
		assertion.Ensure(b.hasSameEntries(oldEntries, newEntries), "new root should contain same entries as old root, except for new entry")
	}(b.ToArray(root))

	if root == nil {
		node := tree.NewBinaryNode[entry[K, V]](entry[K, V]{Key: key, Value: value})
		b.size++
		return &node
	}

	compResult := b.keyComp(key, root.Data.Key)
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

func (b *BSTDict[K, V]) maxNode(root **tree.BinaryNode[entry[K, V]]) (result **tree.BinaryNode[entry[K, V]]) {
	assertion.Require(b.keyComp != nil && b.entryComp != nil, "comparison function is not nil")
	assertion.Require(root != nil && *root != nil, "root points to some node")
	_, maxKey, isOrdered := b.isOrderedHasMinMax(*root)
	assertion.Require(isOrdered, "root node is ordered")
	defer func(maxKey *K) {
		assertion.Ensure(result != nil && (*result) != nil, "result points to some node")
		assertion.Ensure((*result).Right == nil, "result node is the right most one")
		assertion.Ensure((*result).Data.Key == *maxKey, "result node has max key")
	}(maxKey)

	curr := root
	for (*curr).Right != nil {
		curr = &(*curr).Right
	}
	return curr
}

func (b *BSTDict[K, V]) minNode(root **tree.BinaryNode[entry[K, V]]) (result **tree.BinaryNode[entry[K, V]]) {
	assertion.Require(b.keyComp != nil && b.entryComp != nil, "comparison function is not nil")
	assertion.Require(root != nil && *root != nil, "root points to some node")
	minKey, _, isOrdered := b.isOrderedHasMinMax(*root)
	assertion.Require(isOrdered, "root node is ordered")
	defer func(minKey *K) {
		assertion.Ensure(result != nil && (*result) != nil, "result points to some node")
		assertion.Ensure((*result).Left == nil, "result node is the left most one")
		assertion.Ensure((*result).Data.Key == *minKey, "result node has min key")
	}(minKey)

	curr := root
	for (*curr).Left != nil {
		curr = &(*curr).Left
	}
	return curr
}

func (b *BSTDict[K, V]) remove(pRoot **tree.BinaryNode[entry[K, V]]) {
	assertion.Require(pRoot != nil && *pRoot != nil, "root is not nil")
	assertion.Require((*pRoot).IsBinaryTree(), "root is valid binary tree")
	_, _, isOrdered := b.isOrderedHasMinMax(*pRoot)
	assertion.Require(isOrdered, "root is ordered")
	defer func(targetEntry entry[K, V], oldEntries []entry[K, V]) {
		assertion.Ensure((*pRoot).IsBinaryTree(), "root is valid binary tree")
		_, _, isOrdered := b.isOrderedHasMinMax(*pRoot)
		assertion.Ensure(isOrdered, "root is ordered")
		newEntries := b.ToArray(*pRoot)
		assertion.Ensure(!array.Contains(targetEntry, newEntries, b.entryComp), "root tree does not contain removed entry")
		oldEntries = array.Filter(targetEntry, oldEntries, b.entryComp)
		assertion.Ensure(b.hasSameEntries(oldEntries, newEntries), "new root should contain same entries as old root excluding removed entry")
	}((*pRoot).Data, b.ToArray(*pRoot))

	switch {
	case (*pRoot).Left == nil && (*pRoot).Right == nil:
		*pRoot = nil
		return
	case (*pRoot).Left != nil:
		pMax := b.maxNode(&(*pRoot).Left)
		(*pRoot).Data = (*pMax).Data
		*pMax = (*pMax).Left
	case (*pRoot).Right != nil:
		pMin := b.minNode(&(*pRoot).Right)
		(*pRoot).Data = (*pMin).Data
		*pMin = (*pMin).Right
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
