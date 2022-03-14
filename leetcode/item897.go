package leetcode

import "algorithm/pkg"

func increasingBST(root *pkg.TreeNode) *pkg.TreeNode {
	var head *pkg.TreeNode
	var tail *pkg.TreeNode
	var dfs func(node *pkg.TreeNode)
	dfs = func(node *pkg.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if head == nil {
			head = node
		} else {
			tail.Right = node
		}
		tail = node
		tail.Left = nil
		dfs(node.Right)
	}
	dfs(root)
	return head
}
