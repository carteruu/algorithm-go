package leetcode_2

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		depth++
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return depth
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
