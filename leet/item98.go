package leet

import (
	"algorithm/pkg"
	"math"
)

func isValidBST(root *pkg.TreeNode) bool {
	var dfs func(node *pkg.TreeNode, min, max int) bool
	dfs = func(node *pkg.TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return dfs(node.Left, min, node.Val) && dfs(node.Right, node.Val, max)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func isValidBST11(root *pkg.TreeNode) bool {
	stack := []*pkg.TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
