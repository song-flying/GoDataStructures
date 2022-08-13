package set

import (
	"github.com/song-flying/GoDataStructures/array"
	"github.com/song-flying/GoDataStructures/pkg/contract"
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/song-flying/GoDataStructures/tree"
)

type BSTSet[E comparable] struct {
	tree *tree.BinaryTree[E]
	comp order.CompareFn[E]
	size int
}

func (t *BSTSet[E]) isOrdered(root *tree.BinaryNode[E], lower, upper *E) bool {
	contract.Require(t.comp != nil, "comparison function is not nil")
	contract.Require(lower == nil || upper == nil || t.comp(*lower, *upper) < 0, "lower < upper")

	if root == nil {
		return true
	}

	return (lower == nil || t.comp(*lower, root.Data) < 0) && t.isOrdered(root.Left, lower, &root.Data) &&
		(upper == nil || t.comp(root.Data, *upper) < 0) && t.isOrdered(root.Right, &root.Data, upper)
}

func (t *BSTSet[E]) isOrderedWithMinMax(root *tree.BinaryNode[E]) (minElement, maxElement *E, isOrdered bool) {
	contract.Require(t.comp != nil, "comparison function is not nil")
	contract.Require(root.IsBinaryTree(), "root is a binary tree")

	if root == nil {
		return nil, nil, true
	}

	if root.Left != nil {
		leftMin, leftMax, leftIsOrdered := t.isOrderedWithMinMax(root.Left)
		if !leftIsOrdered || t.comp(*leftMax, root.Data) >= 0 {
			return nil, nil, false
		} else {
			minElement = leftMin
		}
	} else {
		minElement = &root.Data
	}

	if root.Right != nil {
		rightMin, rightMax, rightIsOrdered := t.isOrderedWithMinMax(root.Right)
		if !rightIsOrdered || t.comp(*rightMin, root.Data) <= 0 {
			return nil, nil, false
		} else {
			maxElement = rightMax
		}
	} else {
		maxElement = &root.Data
	}

	return minElement, maxElement, true
}

func (t *BSTSet[E]) hasSameEntries(a1, a2 []E) bool {
	contract.Require(array.IsSorted(a1, t.comp) && array.IsDistinct(a1, t.comp), "a1 is sorted & distinct")
	contract.Require(array.IsSorted(a2, t.comp) && array.IsDistinct(a2, t.comp), "a2 is sorted & distinct")

	return array.Same(a1, a2)
}

// IsBSTSet data structure invariant
func (t *BSTSet[E]) IsBSTSet() bool {
	return t.tree != nil && t.IsBST(t.tree.Root)
}

func (t *BSTSet[E]) IsBST(root *tree.BinaryNode[E]) bool {
	_, _, isOrdered := t.isOrderedWithMinMax(root)
	return t.comp != nil && root.IsBinaryTree() && isOrdered
}

func NewBSTSet[E comparable](comp order.CompareFn[E]) (result BSTSet[E]) {
	contract.Require(comp != nil, "comparison function is not nil")
	defer func() {
		contract.Ensure(result.IsBSTSet(), "BST invariant holds")
	}()

	t := tree.NewBinaryTree(tree.Nil[E]())

	return BSTSet[E]{
		tree: &t,
		comp: comp,
		size: 0,
	}
}

func (t *BSTSet[E]) lookup(root **tree.BinaryNode[E], element E) **tree.BinaryNode[E] {
	contract.Require(root != nil, "root pointer is not nil")
	contract.Require(t.IsBST(*root), "BST invariant holds")

	if *root == nil {
		return nil
	}

	compResult := t.comp(element, (*root).Data)
	switch {
	case compResult == 0:
		return root
	case compResult < 0:
		return t.lookup(&(*root).Left, element)
	default: //compResult > 0:
		return t.lookup(&(*root).Right, element)
	}
}

func (t *BSTSet[E]) Contains(element E) bool {
	contract.Require(t.IsBSTSet(), "BST invariant holds")

	node := t.lookup(&t.tree.Root, element)
	return node != nil
}

func (t *BSTSet[E]) ToArray(root *tree.BinaryNode[E]) (result []E) {
	contract.Require(t.IsBST(root), "BST invariant holds")
	defer func() {
		contract.Ensure(array.IsSorted(result, t.comp), "result elements are sorted")
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

func (t *BSTSet[E]) insert(root *tree.BinaryNode[E], element E) (result *tree.BinaryNode[E]) {
	contract.Require(t.IsBST(root), "BST invariant holds for root")
	defer func(oldElements []E) {
		contract.Ensure(t.IsBST(result), "BST invariant holds for result")
		newElements := t.ToArray(result)
		extract := func(e E) E { return e }
		contract.Ensure(array.Contains(newElements, element, extract), "new root should contain new element")
		filter := func(e E) bool { return e != element }
		oldElements = array.Filter(oldElements, filter)
		newElements = array.Filter(newElements, filter)
		contract.Ensure(t.hasSameEntries(oldElements, newElements), "new root should contain same entries as old root, except for new element")
	}(t.ToArray(root))

	if root == nil {
		node := tree.NewBinaryNode[E](element)
		t.size++
		return &node
	}

	compResult := t.comp(element, root.Data)
	switch {
	case compResult == 0:
		root.Data = element
	case compResult < 0:
		root.Left = t.insert(root.Left, element)
	default: //compResult > 0:
		root.Right = t.insert(root.Right, element)
	}

	return root
}

func (t *BSTSet[E]) Add(element E) {
	contract.Require(t.IsBSTSet(), "BST invariant holds")
	defer func() {
		contract.Ensure(t.IsBSTSet(), "BST invariant holds")
		contract.Ensure(t.Contains(element), "Contains(element) returns true")
	}()

	t.tree.Root = t.insert(t.tree.Root, element)
}

func (t *BSTSet[E]) maxFrom(root **tree.BinaryNode[E]) (result **tree.BinaryNode[E]) {
	contract.Require(t.comp != nil, "comparison function is not nil")
	contract.Require(root != nil && *root != nil, "root points to some node")
	_, maxElement, isOrdered := t.isOrderedWithMinMax(*root)
	contract.Require(isOrdered, "root node is ordered")
	defer func(maxElement *E) {
		contract.Ensure(result != nil && (*result) != nil, "result points to some node")
		contract.Ensure((*result).Right == nil, "result node is the right most one")
		contract.Ensure((*result).Data == *maxElement, "result node has max element")
	}(maxElement)

	curr := root
	for (*curr).Right != nil {
		curr = &(*curr).Right
	}
	return curr
}

func (t *BSTSet[E]) minFrom(root **tree.BinaryNode[E]) (result **tree.BinaryNode[E]) {
	contract.Require(t.comp != nil, "comparison function is not nil")
	contract.Require(root != nil && *root != nil, "root points to some node")
	minElement, _, isOrdered := t.isOrderedWithMinMax(*root)
	contract.Require(isOrdered, "root node is ordered")
	defer func(minElement *E) {
		contract.Ensure(result != nil && (*result) != nil, "result points to some node")
		contract.Ensure((*result).Left == nil, "result node is the left most one")
		contract.Ensure((*result).Data == *minElement, "result node has min element")
	}(minElement)

	curr := root
	for (*curr).Left != nil {
		curr = &(*curr).Left
	}
	return curr
}

func (t *BSTSet[E]) remove(pRoot **tree.BinaryNode[E]) {
	contract.Require(pRoot != nil && *pRoot != nil, "root is not nil")
	contract.Require(t.IsBST(*pRoot), "BST invariant holds for root")
	defer func(element E, oldElements []E) {
		contract.Ensure(t.IsBST(*pRoot), "BST invariant holds for root")
		newElements := t.ToArray(*pRoot)
		extract := func(e E) E { return e }
		contract.Ensure(!array.Contains(newElements, element, extract), "root tree does not contain removed element")
		filter := func(e E) bool { return e != element }
		oldElements = array.Filter(oldElements, filter)
		contract.Ensure(t.hasSameEntries(oldElements, newElements), "new root should contain same entries as old root excluding removed element")
	}((*pRoot).Data, t.ToArray(*pRoot))

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

func (t *BSTSet[E]) Delete(element E) {
	contract.Require(t.IsBSTSet(), "BST invariant holds")
	defer func() {
		contract.Ensure(t.IsBSTSet(), "BST invariant holds")
		contract.Ensure(!t.Contains(element), "Contains(element) returns false")
	}()

	target := t.lookup(&t.tree.Root, element)
	if target == nil {
		return
	}

	t.remove(target)
	t.size--
}

func (t *BSTSet[E]) Size() (result int) {
	contract.Require(t.IsBSTSet(), "BST invariant holds")
	defer func() {
		contract.Ensure(0 <= result, "result is non-negative")
	}()

	return t.size
}
