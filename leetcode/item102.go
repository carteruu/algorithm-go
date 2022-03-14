package leetcode

import "algorithm/pkg"

func levelOrder(root *pkg.TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var q []*pkg.TreeNode
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		a := make([]int, 0, size)
		for i := 0; i < size; i++ {
			poll := q[0]
			q = q[1:]
			a = append(a, poll.Val)
			if poll.Left != nil {
				q = append(q, poll.Left)
			}
			if poll.Right != nil {
				q = append(q, poll.Right)
			}
		}
		ans = append(ans, a)
	}
	return ans
}
