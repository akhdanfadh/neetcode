/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	qp := []*TreeNode{p}
	qq := []*TreeNode{q}
	for len(qp) > 0 && len(qq) > 0 {
		for i := len(qp); i > 0; i-- {
			np, nq := qp[0], qq[0]
			qp, qq = qp[1:], qq[1:]

			if np == nil && nq == nil { continue }
			if np == nil || nq == nil || np.Val != nq.Val { return false }

			qp = append(qp, np.Left, np.Right)
			qq = append(qq, nq.Left, nq.Right)
		}
	}
    
	return len(qp) == 0 && len(qq) == 0
}
