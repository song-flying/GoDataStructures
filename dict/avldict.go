package dict

import (
	"github.com/song-flying/GoDataStructures/array"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/song-flying/GoDataStructures/tree"
)

type AVLDict[K comparable, V comparable] struct {
	tree      *tree.BinaryTree[entry[K, V]]
	keyComp   order.CompareFn[K]
	entryComp order.CompareFn[entry[K, V]]
	size      int
}

func (t *AVLDict[K, V]) isOrdered(root *tree.BinaryNode[entry[K, V]], lower, upper *K) bool {
	contract.Require(root.IsBinaryTree(), "root is binary tree")
	contract.Require(t.keyComp != nil, "comparison function is not nil")
	contract.Require(lower == nil || upper == nil || t.keyComp(*lower, *upper) < 0, "lower < upper")

	if root == nil {
		return true
	}

	key := root.Data.Key

	return (lower == nil || t.keyComp(*lower, key) < 0) && t.isOrdered(root.Left, lower, &key) &&
		(upper == nil || t.keyComp(key, *upper) < 0) && t.isOrdered(root.Right, &key, upper)
}

func (t *AVLDict[K, V]) isOrderedWithMinMax(root *tree.BinaryNode[entry[K, V]]) (minKey, maxKey *K, isOrdered bool) {
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

func (t *AVLDict[K, V]) hasSameEntries(a1, a2 []entry[K, V]) bool {
	contract.Require(array.IsSorted(a1, t.entryComp) && array.IsDistinct(a1, t.entryComp), "a1 is sorted & distinct")
	contract.Require(array.IsSorted(a2, t.entryComp) && array.IsDistinct(a2, t.entryComp), "a2 is sorted & distinct")

	return array.Same(a1, a2)
}

func (t *AVLDict[K, V]) isHeightOKFrom(root *tree.BinaryNode[entry[K, V]]) bool {
	contract.Require(root.IsBinaryTree(), "root is binary tree")
	return root == nil || t.isHeightOKFrom(root.Left) && t.isHeightOKFrom(root.Right) && root.Height == order.Max(root.Left.GetHeight(), root.Right.GetHeight())+1
}

func (t *AVLDict[K, V]) isBalancedFrom(root *tree.BinaryNode[entry[K, V]]) bool {
	contract.Require(root.IsBinaryTree(), "root is binary tree")
	return root == nil || t.isBalancedFrom(root.Left) && t.isBalancedFrom(root.Right) && abs(root.Left.GetHeight()-root.Right.GetHeight()) <= 1
}

// IsAVLDict data structure invariant
func (t *AVLDict[K, V]) IsAVLDict() bool {
	return t.tree != nil && t.IsAVL(t.tree.Root)
}

func (t *AVLDict[K, V]) IsAVL(root *tree.BinaryNode[entry[K, V]]) bool {
	_, _, isOrdered := t.isOrderedWithMinMax(root)
	return t.keyComp != nil && t.entryComp != nil &&
		root.IsBinaryTree() && isOrdered && t.isHeightOKFrom(root) && t.isBalancedFrom(root)
}

func NewAVLDict[K comparable, V comparable](comp order.CompareFn[K]) (result AVLDict[K, V]) {
	contract.Require(comp != nil, "comparison function is not nil")
	defer func() {
		contract.Ensure(result.IsAVLDict(), "AVL invariant holds")
	}()

	t := tree.NewBinaryTree(tree.Nil[entry[K, V]]())
	entryComp := func(e1, e2 entry[K, V]) int {
		return comp(e1.Key, e2.Key)
	}

	return AVLDict[K, V]{
		tree:      &t,
		keyComp:   comp,
		entryComp: entryComp,
		size:      0,
	}
}

func (t *AVLDict[K, V]) lookup(root *tree.BinaryNode[entry[K, V]], key K) *tree.BinaryNode[entry[K, V]] {
	contract.Require(t.IsAVL(root), "AVL invariant holds from root")

	if root == nil {
		return nil
	}

	compResult := t.keyComp(key, root.Data.Key)
	switch {
	case compResult == 0:
		return root
	case compResult < 0:
		return t.lookup(root.Left, key)
	default: //compResult > 0:
		return t.lookup(root.Right, key)
	}
}

func (t *AVLDict[K, V]) Get(key K) (*V, bool) {
	contract.Require(t.IsAVLDict(), "AVL invariant holds")

	node := t.lookup(t.tree.Root, key)
	if node != nil {
		return &node.Data.Value, true
	}
	return new(V), false
}

func (t *AVLDict[K, V]) ToArray(root *tree.BinaryNode[entry[K, V]]) (result []entry[K, V]) {
	contract.Require(t.IsAVL(root), "AVL invariant holds")
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

func (t *AVLDict[K, V]) insertFrom(root *tree.BinaryNode[entry[K, V]], key K, value V) (result *tree.BinaryNode[entry[K, V]]) {
	contract.Require(t.IsAVL(root), "AVL invariant holds")
	defer func(oldEntries []entry[K, V]) {
		contract.Ensure(t.IsAVL(result), "AVL invariant holds")
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
		node.Height = 1
		t.size++
		return &node
	}

	compResult := t.keyComp(key, root.Data.Key)
	switch {
	case compResult == 0:
		root.Data.Value = value
	case compResult < 0:
		root.Left = t.insertFrom(root.Left, key, value) // invariant broken (not balanced)
		root = t.rebalance(root)                        // invariant restored (balanced)
	default: //compResult > 0:
		root.Right = t.insertFrom(root.Right, key, value) // invariant broken (not balanced)
		root = t.rebalance(root)                          // invariant restored (balanced)
	}

	return root
}

func (t *AVLDict[K, V]) Put(key K, value V) {
	contract.Require(t.IsAVLDict(), "AVL invariant holds")
	defer func() {
		contract.Ensure(t.IsAVLDict(), "AVL invariant holds")
		v, ok := t.Get(key)
		contract.Ensure(ok && value == *v, "Get(key) returns value")
	}()

	t.tree.Root = t.insertFrom(t.tree.Root, key, value)
}

func (t *AVLDict[K, V]) removeFrom(root *tree.BinaryNode[entry[K, V]], key K) (result *tree.BinaryNode[entry[K, V]]) {
	contract.Require(t.IsAVL(root), "AVL invariant holds")
	defer func(oldEntries []entry[K, V]) {
		contract.Ensure(t.IsAVL(result), "AVL invariant holds")
		newEntries := t.ToArray(result)
		extract := func(e entry[K, V]) K { return e.Key }
		contract.Ensure(!array.Contains(newEntries, key, extract), "root tree does not contain removed entry")
		filter := func(e entry[K, V]) bool { return e.Key != key }
		oldEntries = array.Filter(oldEntries, filter)
		contract.Ensure(t.hasSameEntries(oldEntries, newEntries), "new root should contain same entries as old root excluding removed entry")
	}(t.ToArray(root))

	if root == nil {
		return nil
	}

	compResult := t.keyComp(key, root.Data.Key)

	switch {
	case compResult < 0:
		root.Left = t.removeFrom(root.Left, key) // invariant broken (not balanced)
		root = t.rebalance(root)                 // invariant restored (balanced)
	case compResult > 0:
		root.Right = t.removeFrom(root.Right, key) // invariant broken (not balanced)
		root = t.rebalance(root)                   // invariant restored (balanced)
	default: // compResult == 0
		switch {
		case root.Left != nil:
			root.Left, root.Data = t.removeMax(root.Left) // invariant broken (not balanced)
			root = t.rebalance(root)                      // invariant restored (balanced)
		case root.Right != nil:
			root.Right, root.Data = t.removeMin(root.Right) // invariant broken (not balanced)
			root = t.rebalance(root)                        // invariant restored (balanced)
		default:
			root = nil
		}
		t.size--
	}

	return root
}

func (t *AVLDict[K, V]) removeMax(root *tree.BinaryNode[entry[K, V]]) (result *tree.BinaryNode[entry[K, V]], max entry[K, V]) {
	contract.Require(root != nil, "root is not nil")
	contract.Require(t.IsAVL(root), "AVL invariant holds")
	defer func(oldEntries []entry[K, V]) {
		contract.Ensure(t.IsAVL(result), "AVL invariant holds")
		newEntries := t.ToArray(result)
		extract := func(e entry[K, V]) K { return e.Key }
		contract.Ensure(!array.Contains(newEntries, max.Key, extract), "root tree does not contain removed entry")
		filter := func(e entry[K, V]) bool { return e.Key != max.Key }
		oldEntries = array.Filter(oldEntries, filter)
		contract.Ensure(t.hasSameEntries(oldEntries, newEntries), "new root should contain same entries as old root excluding removed entry")
	}(t.ToArray(root))

	if root.Right == nil {
		return root.Left, root.Data
	}

	root.Right, max = t.removeMax(root.Right) // invariant broken (not balanced)
	root = t.rebalance(root)                  // invariant restored (balanced)

	return root, max
}

func (t *AVLDict[K, V]) removeMin(root *tree.BinaryNode[entry[K, V]]) (result *tree.BinaryNode[entry[K, V]], min entry[K, V]) {
	contract.Require(root != nil, "root is not nil")
	contract.Require(t.IsAVL(root), "AVL invariant holds")
	defer func(oldEntries []entry[K, V]) {
		contract.Ensure(t.IsAVL(result), "AVL invariant holds")
		newEntries := t.ToArray(result)
		extract := func(e entry[K, V]) K { return e.Key }
		contract.Ensure(!array.Contains(newEntries, min.Key, extract), "root tree does not contain removed entry")
		filter := func(e entry[K, V]) bool { return e.Key != min.Key }
		oldEntries = array.Filter(oldEntries, filter)
		contract.Ensure(t.hasSameEntries(oldEntries, newEntries), "new root should contain same entries as old root excluding removed entry")
	}(t.ToArray(root))

	if root.Left == nil {
		return root.Right, root.Data
	}

	root.Left, min = t.removeMin(root.Left) // invariant broken (not balanced)
	root = t.rebalance(root)                // invariant restored (balanced)

	return root, min
}

func (t *AVLDict[K, V]) Delete(key K) {
	contract.Require(t.IsAVLDict(), "AVL invariant holds")
	defer func() {
		contract.Ensure(t.IsAVLDict(), "AVL invariant holds")
		_, ok := t.Get(key)
		contract.Ensure(!ok, "Get(key) returns no value")
	}()

	t.tree.Root = t.removeFrom(t.tree.Root, key)
}

func (t *AVLDict[K, V]) Size() (result int) {
	contract.Require(t.IsAVLDict(), "AVL invariant holds")
	defer func() {
		contract.Ensure(0 <= result, "result is non-negative")
	}()

	return t.size
}

func (t *AVLDict[K, V]) rebalance(root *tree.BinaryNode[entry[K, V]]) (result *tree.BinaryNode[entry[K, V]]) {
	contract.Require(root != nil, "root is not nil")
	contract.Require(t.IsAVL(root.Left), "left child is AVL")
	contract.Require(t.IsAVL(root.Right), "right child is AVL")
	defer func() {
		contract.Ensure(t.IsAVL(result), "result is AVL")
	}()

	diffLR := root.Left.GetHeight() - root.Right.GetHeight()
	var u, v, w *tree.BinaryNode[entry[K, V]]
	switch {
	case diffLR < -1:
		u = root
		contract.Assert(u != nil, "u is not nil")
		v = root.Right
		contract.Assert(v != nil, "v is not nil")
		if v.Left.GetHeight() > v.Right.GetHeight() {
			w = v.Left
			contract.Assert(w != nil, "w is not nil")
			u.Right = w.Left
			v.Left = w.Right
			w.Left = u
			w.Right = v
			u.SetHeight()
			v.SetHeight()
			w.SetHeight()
			root = w
		} else {
			w = v.Right
			contract.Assert(w != nil, "w is not nil")
			u.Right = v.Left
			v.Left = u
			u.SetHeight()
			v.SetHeight()
			root = v
		}

	case diffLR > 1:
		u = root
		contract.Assert(u != nil, "u is not nil")
		v = root.Left
		contract.Assert(v != nil, "v is not nil")
		if v.Right.GetHeight() > v.Left.GetHeight() {
			w = v.Right
			contract.Assert(w != nil, "w is not nil")
			u.Left = w.Right
			v.Right = w.Left
			w.Right = u
			w.Left = v
			u.SetHeight()
			v.SetHeight()
			w.SetHeight()
			root = w
		} else {
			w = v.Left
			contract.Assert(w != nil, "w is not nil")
			u.Left = v.Right
			v.Right = u
			u.SetHeight()
			v.SetHeight()
			root = v
		}

	default: // -1 <= diffLR <= 1
		root.SetHeight()
	}

	return root
}
