package leetcode

import "algorithm/pkg"

func postorderTraversal2(root *pkg.TreeNode) (res []int) {
	stk := []*pkg.TreeNode{}
	var prev *pkg.TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}

func postorderTraversal(root *pkg.TreeNode) []int {
	var ans []int
	var stack []*pkg.TreeNode
	node := root
	var prev *pkg.TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		topIdx := len(stack) - 1
		node = stack[topIdx]
		stack = stack[:topIdx]
		if node.Right == nil || node == prev {
			ans = append(ans, node.Val)
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}
	return ans
}
func postorderTraversal1(root *pkg.TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	ans = append(ans, postorderTraversal(root.Left)...)
	ans = append(ans, postorderTraversal(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}
