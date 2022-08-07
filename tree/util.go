package tree

import "github.com/song-flying/GoDataStructures/queue"

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
