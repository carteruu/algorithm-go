package interview

import "math"

func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	var prev *TreeNode = nil
	node := root
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev == nil || node.Val > prev.Val {
			prev = node
		} else {
			return false
		}
		node = node.Right
	}
	return true
}
func isValidBST1(root *TreeNode) bool {
	var bst func(node *TreeNode, min, max int) bool
	bst = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val >= max || node.Val <= min {
			return false
		}
		return bst(node.Left, min, node.Val) && bst(node.Right, node.Val, max)
	}
	return bst(root, math.MinInt64, math.MaxInt64)
}
