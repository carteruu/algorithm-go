package leetcode_1

import "math"

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, min, max int) bool
	dfs = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		return min < node.Val && node.Val < max && dfs(node.Left, min, node.Val) && dfs(node.Right, node.Val, max)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}
