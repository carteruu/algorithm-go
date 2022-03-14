package leetcode

import "algorithm/pkg"

func averageOfLevels(root *pkg.TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	var ans []float64
	var q []*pkg.TreeNode
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		sum := 0
		for i := 0; i < size; i++ {
			poll := q[0]
			q = q[1:]
			sum += poll.Val
			if poll.Left != nil {
				q = append(q, poll.Left)
			}
			if poll.Right != nil {
				q = append(q, poll.Right)
			}
		}
		ans = append(ans, float64(sum)/float64(size))
	}
	return ans
}
