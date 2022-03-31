package leetcode_1

func preorder(root *Node) []int {
	var ans []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}
	dfs(root)
	return ans
}
