package leet

import "algorithm/pkg"

func countUnivalSubtrees(root *pkg.TreeNode) int {
	res := 0
	var postorder func(node *pkg.TreeNode) bool
	postorder = func(node *pkg.TreeNode) bool {
		if node == nil {
			return true
		}
		left := postorder(node.Left)
		right := postorder(node.Right)
		eq := left && right && (node.Left == nil || node.Val == node.Left.Val) && (node.Right == nil || node.Val == node.Right.Val)
		if eq {
			res++
		}
		return eq
	}
	postorder(root)
	return res
}
