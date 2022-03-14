package leetcode

import "algorithm/pkg"

func findSecondMinimumValue(root *pkg.TreeNode) int {
	var ord func(node *pkg.TreeNode) int
	ord = func(node *pkg.TreeNode) int {
		if node == nil {
			return -1
		}
		if node.Val > root.Val {
			return node.Val
		}
		left := ord(node.Left)
		right := ord(node.Right)
		if left == -1 {
			return right
		}
		if right == -1 {
			return left
		}
		if left < right {
			return left
		} else {
			return right
		}
	}

	return ord(root)
}
