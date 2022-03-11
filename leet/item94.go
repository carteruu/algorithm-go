package leet

import "algorithm/pkg"

func inorderTraversal(root *pkg.TreeNode) []int {
	var ans []int
	var stack []*pkg.TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		topIdx := len(stack) - 1
		pop := stack[topIdx]
		ans = append(ans, pop.Val)
		stack = stack[:topIdx]
		node = pop.Right
	}
	return ans
}
func inorderTraversal1(root *pkg.TreeNode) []int {
	var ans []int
	var dfs func(node *pkg.TreeNode)
	dfs = func(node *pkg.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
