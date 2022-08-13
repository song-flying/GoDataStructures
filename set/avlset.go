package set

import (
	"github.com/song-flying/GoDataStructures/array"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/song-flying/GoDataStructures/tree"
)

type AVLSet[E comparable] struct {
	tree *tree.BinaryTree[E]
	comp order.CompareFn[E]
	size int
}

func (t *AVLSet[E]) isOrdered(root *tree.BinaryNode[E], lower, upper *E) bool {
	contract.Require(root.IsBinaryTree(), "root is binary tree")
	contract.Require(t.comp != nil, "comparison function is not nil")
	contract.Require(lower == nil || upper == nil || t.comp(*lower, *upper) < 0, "lower < upper")

	if root == nil {
		return true
	}

	key := root.Data

	return (lower == nil || t.comp(*lower, key) < 0) && t.isOrdered(root.Left, lower, &key) &&
		(upper == nil || t.comp(key, *upper) < 0) && t.isOrdered(root.Right, &key, upper)
}

func (t *AVLSet[E]) isOrderedWithMinMax(root *tree.BinaryNode[E]) (minKey, maxKey *E, isOrdered bool) {
	contract.Require(t.comp != nil, "comparison function is not nil")
	contract.Require(root.IsBinaryTree(), "root is a binary tree")

	if root == nil {
		return nil, nil, true
	}

	if root.Left != nil {
		leftMinKey, leftMaxKey, leftIsOrdered := t.isOrderedWithMinMax(root.Left)
		if !leftIsOrdered || t.comp(*leftMaxKey, root.Data) >= 0 {
			return nil, nil, false
		} else {
			minKey = leftMinKey
		}
	} else {
		minKey = &root.Data
	}

	if root.Right != nil {
		rightMinKey, rightMaxKey, rightIsOrdered := t.isOrderedWithMinMax(root.Right)
		if !rightIsOrdered || t.comp(*rightMinKey, root.Data) <= 0 {
			return nil, nil, false
		} else {
			maxKey = rightMaxKey
		}
	} else {
		maxKey = &root.Data
	}

	return minKey, maxKey, true
}

func (t *AVLSet[E]) hasSameEntries(a1, a2 []E) bool {
	contract.Require(array.IsSorted(a1, t.comp) && array.IsDistinct(a1, t.comp), "a1 is sorted & distinct")
	contract.Require(array.IsSorted(a2, t.comp) && array.IsDistinct(a2, t.comp), "a2 is sorted & distinct")

	return array.Same(a1, a2)
}

func (t *AVLSet[E]) isHeightOKFrom(root *tree.BinaryNode[E]) bool {
	contract.Require(root.IsBinaryTree(), "root is binary tree")
	return root == nil || t.isHeightOKFrom(root.Left) && t.isHeightOKFrom(root.Right) && root.Height == order.Max(root.Left.GetHeight(), root.Right.GetHeight())+1
}

func (t *AVLSet[E]) isBalancedFrom(root *tree.BinaryNode[E]) bool {
	contract.Require(root.IsBinaryTree(), "root is binary tree")
	return root == nil || t.isBalancedFrom(root.Left) && t.isBalancedFrom(root.Right) && abs(root.Left.GetHeight()-root.Right.GetHeight()) <= 1
}

// IsAVLSet data structure invariant
func (t *AVLSet[E]) IsAVLSet() bool {
	return t.tree != nil && t.IsAVL(t.tree.Root)
}

func (t *AVLSet[E]) IsAVL(root *tree.BinaryNode[E]) bool {
	_, _, isOrdered := t.isOrderedWithMinMax(root)
	return t.comp != nil &&
		root.IsBinaryTree() && isOrdered && t.isHeightOKFrom(root) && t.isBalancedFrom(root)
}

func NewAVLSet[E comparable](comp order.CompareFn[E]) (result AVLSet[E]) {
	contract.Require(comp != nil, "comparison function is not nil")
	defer func() {
		contract.Ensure(result.IsAVLSet(), "AVL invariant holds")
	}()

	t := tree.NewBinaryTree(tree.Nil[E]())

	return AVLSet[E]{
		tree: &t,
		comp: comp,
		size: 0,
	}
}

func (t *AVLSet[E]) lookup(root *tree.BinaryNode[E], element E) *tree.BinaryNode[E] {
	contract.Require(t.IsAVL(root), "AVL invariant holds")

	if root == nil {
		return nil
	}

	compResult := t.comp(element, root.Data)
	switch {
	case compResult == 0:
		return root
	case compResult < 0:
		return t.lookup(root.Left, element)
	default: //compResult > 0:
		return t.lookup(root.Right, element)
	}
}

func (t *AVLSet[E]) Contains(element E) bool {
	contract.Require(t.IsAVLSet(), "AVL invariant holds")

	node := t.lookup(t.tree.Root, element)

	return node != nil
}

func (t *AVLSet[E]) ToArray(root *tree.BinaryNode[E]) (result []E) {
	contract.Require(t.IsAVL(root), "AVL invariant holds")
	defer func() {
		contract.Ensure(array.IsSorted(result, t.comp), "result entries are sorted")
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

func (t *AVLSet[E]) insertFrom(root *tree.BinaryNode[E], element E) (result *tree.BinaryNode[E]) {
	contract.Require(t.IsAVL(root), "AVL invariant holds")
	defer func(oldElements []E) {
		contract.Ensure(t.IsAVL(result), "AVL invariant holds")
		newElements := t.ToArray(result)
		extract := func(e E) E { return e }
		contract.Ensure(array.Contains(newElements, element, extract), "new root should contain new element")
		filter := func(e E) bool { return e != element }
		oldElements = array.Filter(oldElements, filter)
		newElements = array.Filter(newElements, filter)
		contract.Ensure(t.hasSameEntries(oldElements, newElements), "new root should contain same elements as old root, except for new element")
	}(t.ToArray(root))

	if root == nil {
		node := tree.NewBinaryNode[E](element)
		node.Height = 1
		t.size++
		return &node
	}

	compResult := t.comp(element, root.Data)
	switch {
	case compResult == 0:
		root.Data = element
	case compResult < 0:
		root.Left = t.insertFrom(root.Left, element) // invariant broken (not balanced)
		root = t.rebalance(root)                     // invariant restored (balanced)
	default: //compResult > 0:
		root.Right = t.insertFrom(root.Right, element) // invariant broken (not balanced)
		root = t.rebalance(root)                       // invariant restored (balanced)
	}

	return root
}

func (t *AVLSet[E]) Add(element E) {
	contract.Require(t.IsAVLSet(), "AVL invariant holds")
	defer func() {
		contract.Ensure(t.IsAVLSet(), "AVL invariant holds")
		contract.Ensure(t.Contains(element), "Contains(element) returns true")
	}()

	t.tree.Root = t.insertFrom(t.tree.Root, element)
}

func (t *AVLSet[E]) removeFrom(root *tree.BinaryNode[E], element E) (result *tree.BinaryNode[E]) {
	contract.Require(t.IsAVL(root), "AVL invariant holds")
	defer func(oldElements []E) {
		contract.Ensure(t.IsAVL(result), "AVL invariant holds")
		newElements := t.ToArray(result)
		extract := func(e E) E { return e }
		contract.Ensure(!array.Contains(newElements, element, extract), "root tree does not contain removed element")
		filter := func(e E) bool { return e != element }
		oldElements = array.Filter(oldElements, filter)
		contract.Ensure(t.hasSameEntries(oldElements, newElements), "new root should contain same element as old root excluding removed element")
	}(t.ToArray(root))

	if root == nil {
		return nil
	}

	compResult := t.comp(element, root.Data)

	switch {
	case compResult < 0:
		root.Left = t.removeFrom(root.Left, element) // invariant broken (not balanced)
		root = t.rebalance(root)                     // invariant restored (balanced)
	case compResult > 0:
		root.Right = t.removeFrom(root.Right, element) // invariant broken (not balanced)
		root = t.rebalance(root)                       // invariant restored (balanced)
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

func (t *AVLSet[E]) removeMax(root *tree.BinaryNode[E]) (result *tree.BinaryNode[E], max E) {
	contract.Require(root != nil, "root is not nil")
	contract.Require(t.IsAVL(root), "AVL invariant holds")
	defer func(oldElements []E) {
		contract.Ensure(t.IsAVL(result), "AVL invariant holds")
		newEntries := t.ToArray(result)
		extract := func(e E) E { return e }
		contract.Ensure(!array.Contains(newEntries, max, extract), "root tree does not contain removed element")
		filter := func(e E) bool { return e != max }
		oldElements = array.Filter(oldElements, filter)
		contract.Ensure(t.hasSameEntries(oldElements, newEntries), "new root should contain same elements as old root excluding removed elements")
	}(t.ToArray(root))

	if root.Right == nil {
		return root.Left, root.Data
	}

	root.Right, max = t.removeMax(root.Right) // invariant broken (not balanced)
	root = t.rebalance(root)                  // invariant restored (balanced)

	return root, max
}

func (t *AVLSet[E]) removeMin(root *tree.BinaryNode[E]) (result *tree.BinaryNode[E], min E) {
	contract.Require(root != nil, "root is not nil")
	contract.Require(t.IsAVL(root), "AVL invariant holds")
	defer func(oldElements []E) {
		contract.Ensure(t.IsAVL(result), "AVL invariant holds")
		newElements := t.ToArray(result)
		extract := func(e E) E { return e }
		contract.Ensure(!array.Contains(newElements, min, extract), "root tree does not contain removed element")
		filter := func(e E) bool { return e != min }
		oldElements = array.Filter(oldElements, filter)
		contract.Ensure(t.hasSameEntries(oldElements, newElements), "new root should contain same elements as old root excluding removed elements")
	}(t.ToArray(root))

	if root.Left == nil {
		return root.Right, root.Data
	}

	root.Left, min = t.removeMin(root.Left) // invariant broken (not balanced)
	root = t.rebalance(root)                // invariant restored (balanced)

	return root, min
}

func (t *AVLSet[E]) Delete(element E) {
	contract.Require(t.IsAVLSet(), "AVL invariant holds")
	defer func() {
		contract.Ensure(t.IsAVLSet(), "AVL invariant holds")
		contract.Ensure(!t.Contains(element), "Contains(element) returns false")
	}()

	t.tree.Root = t.removeFrom(t.tree.Root, element)
}

func (t *AVLSet[E]) Size() (result int) {
	contract.Require(t.IsAVLSet(), "AVL invariant holds")
	defer func() {
		contract.Ensure(0 <= result, "result is non-negative")
	}()

	return t.size
}

func (t *AVLSet[E]) rebalance(root *tree.BinaryNode[E]) (result *tree.BinaryNode[E]) {
	contract.Require(root != nil, "root is not nil")
	contract.Require(t.IsAVL(root.Left), "left child is AVL")
	contract.Require(t.IsAVL(root.Right), "right child is AVL")
	defer func() {
		contract.Ensure(t.IsAVL(result), "result is AVL")
	}()

	diffLR := root.Left.GetHeight() - root.Right.GetHeight()
	var u, v, w *tree.BinaryNode[E]
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
