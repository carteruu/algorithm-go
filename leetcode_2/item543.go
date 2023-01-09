package leetcode_2

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, ans := l(root)
	return ans
}

func l(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	lDepth, lAns := l(node.Left)
	rDepth, rAns := l(node.Right)
	return maxInt(lDepth, rDepth) + 1, maxInt(maxInt(lAns, rAns), lDepth+rDepth-1)
}
