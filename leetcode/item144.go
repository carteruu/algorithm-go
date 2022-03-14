package leetcode

import "algorithm/pkg"

func preorderTraversal(root *pkg.TreeNode) (vals []int) {
	var stack []*pkg.TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

func preorderTraversal2(root *pkg.TreeNode) (ans []int) {
	if root == nil {
		return
	}
	var stack []*pkg.TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		topIdx := len(stack) - 1
		pop := stack[topIdx]
		stack = stack[:topIdx]
		if pop == nil {
			continue
		}
		ans = append(ans, pop.Val)
		stack = append(stack, pop.Right, pop.Left)
	}
	return
}
func preorderTraversal1(root *pkg.TreeNode) (ans []int) {
	if root == nil {
		return
	}
	ans = append(ans, root.Val)
	ans = append(ans, preorderTraversal(root.Left)...)
	ans = append(ans, preorderTraversal(root.Right)...)
	return ans
}
