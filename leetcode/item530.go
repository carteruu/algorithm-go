package leetcode

import (
	"algorithm/pkg"
	"math"
)

var arr []int

func getMinimumDifference(root *pkg.TreeNode) int {
	rs, pre := math.MaxInt64, -1
	var inorder func(node *pkg.TreeNode)
	inorder = func(node *pkg.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		diff := node.Val - pre
		if pre != -1 && diff < rs {
			rs = diff
		}
		pre = node.Val
		inorder(node.Right)
	}
	inorder(root)
	return rs
}
func getMinimumDifference1(root *pkg.TreeNode) int {
	arr = []int{}
	inorder(root)
	rs := math.MaxInt64
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < rs {
			rs = diff
		}
	}
	return rs
}

func inorder(node *pkg.TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	arr = append(arr, node.Val)
	inorder(node.Right)
}
