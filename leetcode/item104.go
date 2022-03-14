package leetcode

import "algorithm/pkg"

func maxDepth(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
func maxDepth1(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var q []*pkg.TreeNode
	q = append(q, root)
	for len(q) > 0 {
		ans++
		size := len(q)
		for i := 0; i < size; i++ {
			poll := q[0]
			q = q[1:]
			if poll.Left != nil {
				q = append(q, poll.Left)
			}
			if poll.Right != nil {
				q = append(q, poll.Right)
			}
		}
	}
	return ans
}
