package leetcode_1

func postorder(root *Node) []int {
	var ans []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			dfs(child)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return ans
}
