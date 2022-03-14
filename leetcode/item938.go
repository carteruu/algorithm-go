package leetcode

import "algorithm/pkg"

func rangeSumBST11(root *pkg.TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	} else if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	return root.Val + rangeSumBST(root.Right, low, high) + rangeSumBST(root.Left, low, high)
}
