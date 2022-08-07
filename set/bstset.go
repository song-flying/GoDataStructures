package set

import (
	"github.com/song-flying/GoDataStructures/array"
	"github.com/song-flying/GoDataStructures/pkg/assertion"
	"github.com/song-flying/GoDataStructures/pkg/order"
	"github.com/song-flying/GoDataStructures/tree"
)

type BSTSet[E comparable] struct {
	tree *tree.BinaryTree[E]
	comp order.CompareFn[E]
	size int
}

func (b *BSTSet[E]) isOrdered(root *tree.BinaryNode[E], lower, upper *E) bool {
	assertion.Require(b.comp != nil, "comparison function is not nil")
	assertion.Require(lower == nil || upper == nil || b.comp(*lower, *upper) < 0, "lower < upper")

	if root == nil {
		return true
	}

	return (lower == nil || b.comp(*lower, root.Data) < 0) && b.isOrdered(root.Left, lower, &root.Data) &&
		(upper == nil || b.comp(root.Data, *upper) < 0) && b.isOrdered(root.Right, &root.Data, upper)
}

func (b *BSTSet[E]) isOrderedHasMinMax(root *tree.BinaryNode[E]) (minElement, maxElement *E, isOrdered bool) {
	assertion.Require(b.comp != nil, "comparison function is not nil")
	assertion.Require(root.IsBinaryTree(), "root is a binary tree")

	if root == nil {
		return nil, nil, true
	}

	if root.Left != nil {
		leftMin, leftMax, leftIsOrdered := b.isOrderedHasMinMax(root.Left)
		if !leftIsOrdered || b.comp(*leftMax, root.Data) >= 0 {
			return nil, nil, false
		} else {
			minElement = leftMin
		}
	} else {
		minElement = &root.Data
	}

	if root.Right != nil {
		rightMin, rightMax, rightIsOrdered := b.isOrderedHasMinMax(root.Right)
		if !rightIsOrdered || b.comp(*rightMin, root.Data) <= 0 {
			return nil, nil, false
		} else {
			maxElement = rightMax
		}
	} else {
		maxElement = &root.Data
	}

	return minElement, maxElement, true
}

func (b *BSTSet[E]) hasSameEntries(a1, a2 []E) bool {
	assertion.Require(array.IsSorted(a1, b.comp) && array.IsDistinct(a1, b.comp), "a1 is sorted & distinct")
	assertion.Require(array.IsSorted(a2, b.comp) && array.IsDistinct(a2, b.comp), "a2 is sorted & distinct")

	return array.Same(a1, a2)
}

// IsBSTSet data structure invariant
func (b *BSTSet[E]) IsBSTSet() bool {
	_, _, isOrdered := b.isOrderedHasMinMax(b.tree.Root)
	return b.tree != nil && b.comp != nil && b.tree.IsBinaryTree() && isOrdered
}

func NewBSTSet[E comparable](comp order.CompareFn[E]) (result BSTSet[E]) {
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

func (b *BSTSet[E]) lookup(root **tree.BinaryNode[E], element E) **tree.BinaryNode[E] {
	assertion.Require(b.IsBSTSet(), "BST invariant holds")

	if *root == nil {
		return nil
	}

	compResult := b.comp(element, (*root).Data)
	switch {
	case compResult == 0:
		return root
	case compResult < 0:
		return b.lookup(&(*root).Left, element)
	default: //compResult > 0:
		return b.lookup(&(*root).Right, element)
	}
}

func (b *BSTSet[E]) Contains(key E) bool {
	assertion.Require(b.IsBSTSet(), "BST invariant holds")

	node := b.lookup(&b.tree.Root, key)
	return node != nil
}

func (b *BSTSet[E]) ToArray(root *tree.BinaryNode[E]) (result []E) {
	assertion.Require(b.IsBSTSet(), "BST invariant holds")
	defer func() {
		assertion.Ensure(array.IsSorted(result, b.comp), "result elements are sorted")
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

func (b *BSTSet[E]) insert(root *tree.BinaryNode[E], element E) (result *tree.BinaryNode[E]) {
	assertion.Require(root.IsBinaryTree(), "root is valid binary tree")
	_, _, isOrdered := b.isOrderedHasMinMax(root)
	assertion.Require(isOrdered, "root tree is ordered")
	defer func(oldElements []E) {
		assertion.Ensure(result.IsBinaryTree(), "new root is valid binary tree")
		assertion.Ensure(b.isOrdered(result, nil, nil), "new root is ordered")
		newElements := b.ToArray(result)
		assertion.Ensure(array.Contains(element, newElements, b.comp), "new root should contain new element")
		oldElements = array.Filter(element, oldElements, b.comp)
		newElements = array.Filter(element, newElements, b.comp)
		assertion.Ensure(b.hasSameEntries(oldElements, newElements), "new root should contain same entries as old root, except for new element")
	}(b.ToArray(root))

	if root == nil {
		node := tree.NewBinaryNode[E](element)
		b.size++
		return &node
	}

	compResult := b.comp(element, root.Data)
	switch {
	case compResult == 0:
		root.Data = element
	case compResult < 0:
		root.Left = b.insert(root.Left, element)
	default: //compResult > 0:
		root.Right = b.insert(root.Right, element)
	}

	return root
}

func (b *BSTSet[E]) Add(element E) {
	assertion.Require(b.IsBSTSet(), "BST invariant holds")
	defer func() {
		assertion.Ensure(b.IsBSTSet(), "BST invariant holds")
		assertion.Ensure(b.Contains(element), "Contains(element) returns true")
	}()

	b.tree.Root = b.insert(b.tree.Root, element)
}

func (b *BSTSet[E]) maxNode(root **tree.BinaryNode[E]) (result **tree.BinaryNode[E]) {
	assertion.Require(b.comp != nil, "comparison function is not nil")
	assertion.Require(root != nil && *root != nil, "root points to some node")
	_, maxElement, isOrdered := b.isOrderedHasMinMax(*root)
	assertion.Require(isOrdered, "root node is ordered")
	defer func(maxElement *E) {
		assertion.Ensure(result != nil && (*result) != nil, "result points to some node")
		assertion.Ensure((*result).Right == nil, "result node is the right most one")
		assertion.Ensure((*result).Data == *maxElement, "result node has max element")
	}(maxElement)

	curr := root
	for (*curr).Right != nil {
		curr = &(*curr).Right
	}
	return curr
}

func (b *BSTSet[E]) minNode(root **tree.BinaryNode[E]) (result **tree.BinaryNode[E]) {
	assertion.Require(b.comp != nil, "comparison function is not nil")
	assertion.Require(root != nil && *root != nil, "root points to some node")
	minElement, _, isOrdered := b.isOrderedHasMinMax(*root)
	assertion.Require(isOrdered, "root node is ordered")
	defer func(minElement *E) {
		assertion.Ensure(result != nil && (*result) != nil, "result points to some node")
		assertion.Ensure((*result).Left == nil, "result node is the left most one")
		assertion.Ensure((*result).Data == *minElement, "result node has min element")
	}(minElement)

	curr := root
	for (*curr).Left != nil {
		curr = &(*curr).Left
	}
	return curr
}

func (b *BSTSet[E]) remove(pRoot **tree.BinaryNode[E]) {
	assertion.Require(pRoot != nil && *pRoot != nil, "root is not nil")
	assertion.Require((*pRoot).IsBinaryTree(), "root is valid binary tree")
	_, _, isOrdered := b.isOrderedHasMinMax(*pRoot)
	assertion.Require(isOrdered, "root is ordered")
	defer func(targetElement E, oldElements []E) {
		assertion.Ensure((*pRoot).IsBinaryTree(), "root is valid binary tree")
		_, _, isOrdered := b.isOrderedHasMinMax(*pRoot)
		assertion.Ensure(isOrdered, "root is ordered")
		newElements := b.ToArray(*pRoot)
		assertion.Ensure(!array.Contains(targetElement, newElements, b.comp), "root tree does not contain removed element")
		oldElements = array.Filter(targetElement, oldElements, b.comp)
		assertion.Ensure(b.hasSameEntries(oldElements, newElements), "new root should contain same entries as old root excluding removed element")
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

func (b *BSTSet[E]) Delete(element E) {
	assertion.Require(b.IsBSTSet(), "BST invariant holds")
	defer func() {
		assertion.Ensure(b.IsBSTSet(), "BST invariant holds")
		assertion.Ensure(!b.Contains(element), "Contains(element) returns false")
	}()

	target := b.lookup(&b.tree.Root, element)
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
