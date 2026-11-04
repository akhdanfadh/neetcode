/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, 0, targetSum)
}

func dfs(node *TreeNode, currSum, targetSum int) bool {
	if node == nil {
		return false
	}

	// only if it is a leaf
	currSum += node.Val
	if currSum == targetSum && node.Left == nil && node.Right == nil {
		return true
	}

	// in any branch, true if either one is true (OR)
	return dfs(node.Left, currSum, targetSum) || dfs (node.Right, currSum, targetSum)
}
