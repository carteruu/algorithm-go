package leetcode

import "algorithm/pkg"

func diameterOfBinaryTree(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}
	rs := 0
	var depthLen func(node *pkg.TreeNode) int
	depthLen = func(node *pkg.TreeNode) int {
		if node == nil {
			return 0
		}
		leftLen := depthLen(node.Left)
		rightLen := depthLen(node.Right)

		diameter := leftLen + rightLen
		if diameter > rs {
			rs = diameter
		}

		if leftLen > rightLen {
			return leftLen + 1
		} else {
			return rightLen + 1
		}
	}
	depthLen(root)
	return rs
}
