package leetcode_2

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var in func(node *TreeNode)
	in = func(node *TreeNode) {
		if node == nil {
			return
		}
		in(node.Left)
		ans = append(ans, node.Val)
		in(node.Right)
	}
	in(root)
	return ans
}
