package offer

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, isB := maxDepth(root)
	return isB
}

func maxDepth(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	leftD, leftB := maxDepth(node.Left)
	if !leftB {
		return 0, false
	}
	rightD, rightB := maxDepth(node.Right)
	if !rightB {
		return 0, false
	}
	return 1 + max(leftD, rightD), abs(leftD, rightD) <= 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func abs(a, b int) int {
	c := a - b
	if c >= 0 {
		return c
	} else {
		return -c
	}
}
