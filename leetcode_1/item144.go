package leetcode_1

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
