/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil { return 0 }

	depth := 0
	q := Queue[*TreeNode]{items: []*TreeNode{root}}
	for !q.IsEmpty() {
		for range len(q.items) {
			node, _ := q.Dequeue()
			if node.Left != nil { q.Enqueue(node.Left) }
			if node.Right != nil { q.Enqueue(node.Right) }
		}
		depth++
	}
	return depth
}

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(val T) {
	q.items = append(q.items, val)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.items) == 0 { return zero, false }
	val := q.items[0]
	q.items = q.items[1:]
	return val, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}