package dict

import (
	"github.com/song-flying/GoDataStructures/array"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/song-flying/GoDataStructures/tree"
)

type BSTDict[K comparable, V comparable] struct {
	tree      *tree.BinaryTree[entry[K, V]]
	keyComp   order.CompareFn[K]
	entryComp order.CompareFn[entry[K, V]]
	size      int
}

func (t *BSTDict[K, V]) isOrdered(root *tree.BinaryNode[entry[K, V]], lower, upper *K) bool {
	contract.Require(t.keyComp != nil, "comparison function is not nil")
	contract.Require(lower == nil || upper == nil || t.keyComp(*lower, *upper) < 0, "lower < upper")

	if root == nil {
		return true
	}

	key := root.Data.Key

	return (lower == nil || t.keyComp(*lower, key) < 0) && t.isOrdered(root.Left, lower, &key) &&
		(upper == nil || t.keyComp(key, *upper) < 0) && t.isOrdered(root.Right, &key, upper)
}

func (t *BSTDict[K, V]) isOrderedWithMinMax(root *tree.BinaryNode[entry[K, V]]) (minKey, maxKey *K, isOrdered bool) {
	contract.Require(t.keyComp != nil, "comparison function is not nil")
	contract.Require(root.IsBinaryTree(), "root is a binary tree")

	if root == nil {
		return nil, nil, true
	}

	if root.Left != nil {
		leftMinKey, leftMaxKey, leftIsOrdered := t.isOrderedWithMinMax(root.Left)
		if !leftIsOrdered || t.keyComp(*leftMaxKey, root.Data.Key) >= 0 {
			return nil, nil, false
		} else {
			minKey = leftMinKey
		}
	} else {
		minKey = &root.Data.Key
	}

	if root.Right != nil {
		rightMinKey, rightMaxKey, rightIsOrdered := t.isOrderedWithMinMax(root.Right)
		if !rightIsOrdered || t.keyComp(*rightMinKey, root.Data.Key) <= 0 {
			return nil, nil, false
		} else {
			maxKey = rightMaxKey
		}
	} else {
		maxKey = &root.Data.Key
	}

	return minKey, maxKey, true
}

func (t *BSTDict[K, V]) hasSameEntries(a1, a2 []entry[K, V]) bool {
	contract.Require(array.IsSorted(a1, t.entryComp) && array.IsDistinct(a1, t.entryComp), "a1 is sorted & distinct")
	contract.Require(array.IsSorted(a2, t.entryComp) && array.IsDistinct(a2, t.entryComp), "a2 is sorted & distinct")

	return array.Same(a1, a2)
}

// IsBSTDict data structure invariant
func (t *BSTDict[K, V]) IsBSTDict() bool {
	return t.tree != nil && t.IsBST(t.tree.Root)
}

func (t *BSTDict[K, V]) IsBST(root *tree.BinaryNode[entry[K, V]]) bool {
	_, _, isOrdered := t.isOrderedWithMinMax(root)
	return t.keyComp != nil && t.entryComp != nil && root.IsBinaryTree() && isOrdered
}

func NewBSTDict[K comparable, V comparable](comp order.CompareFn[K]) (result *BSTDict[K, V]) {
	contract.Require(comp != nil, "comparison function is not nil")
	defer func() {
		contract.Ensure(result.IsBSTDict(), "BST invariant holds")
	}()

	t := tree.NewBinaryTree(tree.Nil[entry[K, V]]())
	entryComp := func(e1, e2 entry[K, V]) int {
		return comp(e1.Key, e2.Key)
	}

	return &BSTDict[K, V]{
		tree:      t,
		keyComp:   comp,
		entryComp: entryComp,
		size:      0,
	}
}

func (t *BSTDict[K, V]) lookup(root **tree.BinaryNode[entry[K, V]], key K) **tree.BinaryNode[entry[K, V]] {
	contract.Require(root != nil, "root pointer is not nil")
	contract.Require(t.IsBST(*root), "BST invariant holds")

	if *root == nil {
		return nil
	}

	compResult := t.keyComp(key, (*root).Data.Key)
	switch {
	case compResult == 0:
		return root
	case compResult < 0:
		return t.lookup(&(*root).Left, key)
	default: //compResult > 0:
		return t.lookup(&(*root).Right, key)
	}
}

func (t *BSTDict[K, V]) Get(key K) (V, bool) {
	contract.Require(t.IsBSTDict(), "BST invariant holds")

	node := t.lookup(&t.tree.Root, key)
	if node != nil {
		return (*node).Data.Value, true
	}
	return *new(V), false
}

func (t *BSTDict[K, V]) ToArray(root *tree.BinaryNode[entry[K, V]]) (result []entry[K, V]) {
	contract.Require(t.IsBST(root), "BST invariant holds")
	defer func() {
		contract.Ensure(array.IsSorted(result, t.entryComp), "result entries are sorted")
	}()
	if root == nil {
		return
	}

	if root.Left != nil {
		result = append(result, t.ToArray(root.Left)...)
	}
	result = append(result, root.Data)
	if root.Right != nil {
		result = append(result, t.ToArray(root.Right)...)
	}

	return
}

func (t *BSTDict[K, V]) insert(root *tree.BinaryNode[entry[K, V]], key K, value V) (result *tree.BinaryNode[entry[K, V]]) {
	contract.Require(t.IsBST(root), "BST invariant holds for root")
	defer func(oldEntries []entry[K, V]) {
		contract.Ensure(t.IsBST(result), "BST invariant holds for result")
		newEntries := t.ToArray(result)
		extract := func(e entry[K, V]) K { return e.Key }
		contract.Ensure(array.Contains(newEntries, key, extract), "new root should contain new entry")
		filter := func(e entry[K, V]) bool { return e.Key != key }
		oldEntries = array.Filter(oldEntries, filter)
		newEntries = array.Filter(newEntries, filter)
		contract.Ensure(t.hasSameEntries(oldEntries, newEntries), "new root should contain same entries as old root, except for new entry")
	}(t.ToArray(root))

	if root == nil {
		node := tree.NewBinaryNode[entry[K, V]](entry[K, V]{Key: key, Value: value})
		t.size++
		return &node
	}

	compResult := t.keyComp(key, root.Data.Key)
	switch {
	case compResult == 0:
		root.Data.Value = value
	case compResult < 0:
		root.Left = t.insert(root.Left, key, value)
	default: //compResult > 0:
		root.Right = t.insert(root.Right, key, value)
	}

	return root
}

func (t *BSTDict[K, V]) Put(key K, value V) {
	contract.Require(t.IsBSTDict(), "BST invariant holds")
	defer func() {
		contract.Ensure(t.IsBSTDict(), "BST invariant holds")
		v, ok := t.Get(key)
		contract.Ensure(ok && value == v, "Get(key) returns value")
	}()

	t.tree.Root = t.insert(t.tree.Root, key, value)
}

func (t *BSTDict[K, V]) maxFrom(root **tree.BinaryNode[entry[K, V]]) (result **tree.BinaryNode[entry[K, V]]) {
	contract.Require(t.keyComp != nil && t.entryComp != nil, "comparison function is not nil")
	contract.Require(root != nil && *root != nil, "root points to some node")
	_, maxKey, isOrdered := t.isOrderedWithMinMax(*root)
	contract.Require(isOrdered, "root node is ordered")
	defer func(maxKey *K) {
		contract.Ensure(result != nil && (*result) != nil, "result points to some node")
		contract.Ensure((*result).Right == nil, "result node is the right most one")
		contract.Ensure((*result).Data.Key == *maxKey, "result node has max key")
	}(maxKey)

	curr := root
	for (*curr).Right != nil {
		curr = &(*curr).Right
	}
	return curr
}

func (t *BSTDict[K, V]) minFrom(root **tree.BinaryNode[entry[K, V]]) (result **tree.BinaryNode[entry[K, V]]) {
	contract.Require(t.keyComp != nil && t.entryComp != nil, "comparison function is not nil")
	contract.Require(root != nil && *root != nil, "root points to some node")
	minKey, _, isOrdered := t.isOrderedWithMinMax(*root)
	contract.Require(isOrdered, "root node is ordered")
	defer func(minKey *K) {
		contract.Ensure(result != nil && (*result) != nil, "result points to some node")
		contract.Ensure((*result).Left == nil, "result node is the left most one")
		contract.Ensure((*result).Data.Key == *minKey, "result node has min key")
	}(minKey)

	curr := root
	for (*curr).Left != nil {
		curr = &(*curr).Left
	}
	return curr
}

func (t *BSTDict[K, V]) remove(pRoot **tree.BinaryNode[entry[K, V]]) {
	contract.Require(pRoot != nil && *pRoot != nil, "root is not nil")
	contract.Require(t.IsBST(*pRoot), "BST invariant holds for root")
	defer func(key K, oldEntries []entry[K, V]) {
		contract.Ensure(t.IsBST(*pRoot), "BST invariant holds for root")
		newEntries := t.ToArray(*pRoot)
		extract := func(e entry[K, V]) K { return e.Key }
		contract.Ensure(!array.Contains(newEntries, key, extract), "root tree does not contain removed entry")
		filter := func(e entry[K, V]) bool { return e.Key != key }
		oldEntries = array.Filter(oldEntries, filter)
		contract.Ensure(t.hasSameEntries(oldEntries, newEntries), "new root should contain same entries as old root excluding removed entry")
	}((*pRoot).Data.Key, t.ToArray(*pRoot))

	switch {
	case (*pRoot).Left == nil && (*pRoot).Right == nil:
		*pRoot = nil
		return
	case (*pRoot).Left != nil:
		pMax := t.maxFrom(&(*pRoot).Left)
		(*pRoot).Data = (*pMax).Data
		*pMax = (*pMax).Left
	case (*pRoot).Right != nil:
		pMin := t.minFrom(&(*pRoot).Right)
		(*pRoot).Data = (*pMin).Data
		*pMin = (*pMin).Right
	}
}

func (t *BSTDict[K, V]) Delete(key K) {
	contract.Require(t.IsBSTDict(), "BST invariant holds")
	defer func() {
		contract.Ensure(t.IsBSTDict(), "BST invariant holds")
		_, ok := t.Get(key)
		contract.Ensure(!ok, "Get(key) returns no value")
	}()

	target := t.lookup(&t.tree.Root, key)
	if target == nil {
		return
	}

	t.remove(target)
	t.size--
}

func (t *BSTDict[K, V]) Size() (result int) {
	contract.Require(t.IsBSTDict(), "BST invariant holds")
	defer func() {
		contract.Ensure(0 <= result, "result is non-negative")
	}()

	return t.size
}
