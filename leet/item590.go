package leet

import (
	"algorithm/pkg/node_children"
)

func postorder(root *node_children.Node) []int {
	ans := make([]int, 0)
	var dfs func(node *node_children.Node)
	dfs = func(node *node_children.Node) {
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
