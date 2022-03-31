package leetcode_1

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0, 1e5)
	q = append(q, root)
	depth := 0
	for len(q) > 0 {
		depth++
		size := len(q)
		for i := 0; i < size; i++ {
			poll := q[0]
			q = q[1:]
			if poll.Left == nil && poll.Right == nil {
				return depth
			}
			if poll.Left != nil {
				q = append(q, poll.Left)
			}
			if poll.Right != nil {
				q = append(q, poll.Right)
			}
		}
	}
	return depth
}

/**
DFS
*/
func minDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	rt := math.MaxInt
	if root.Left != nil {
		rt = min(rt, minDepth(root.Left))
	}
	if root.Right != nil {
		rt = min(rt, minDepth(root.Right))
	}
	return rt + 1
}
