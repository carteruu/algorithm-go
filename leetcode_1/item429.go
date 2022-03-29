package leetcode_1

//dfs
func levelOrder(root *Node) [][]int {
	var ans [][]int
	var dfs func(node *Node, level int)
	dfs = func(node *Node, level int) {
		if node == nil {
			return
		}
		if len(ans)-1 < level {
			ans = append(ans, make([]int, 0, len(node.Children)))
		}
		ans[level] = append(ans[level], node.Val)
		for _, child := range node.Children {
			dfs(child, level+1)
		}
	}
	dfs(root, 0)
	return ans
}

//bfs
func levelOrder_1(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	q := []*Node{root}
	for len(q) > 0 {
		size := len(q)
		ansItem := make([]int, size)
		for i := 0; i < size; i++ {
			pop := q[0]
			q = q[1:]
			ansItem[i] = pop.Val
			for _, child := range pop.Children {
				if child == nil {
					continue
				}
				q = append(q, child)
			}
		}
		ans = append(ans, ansItem)
	}
	return ans
}
