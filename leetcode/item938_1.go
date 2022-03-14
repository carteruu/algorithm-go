package leetcode

import "algorithm/pkg"

func rangeSumBST(root *pkg.TreeNode, low int, high int) int {
	var dfs func(node *pkg.TreeNode) int
	dfs = func(node *pkg.TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Val <= low {
			//当前节点比low小,在右子树找
			if node.Val == low {
				return node.Val + dfs(node.Right)
			}
			return dfs(node.Right)
		}
		if node.Val >= high {
			//当前节点比high大.找左子树找
			if node.Val == high {
				return node.Val + dfs(node.Left)
			}
			return dfs(node.Left)
		}
		return node.Val + dfs(node.Left) + dfs(node.Right)
	}
	return dfs(root)
}
