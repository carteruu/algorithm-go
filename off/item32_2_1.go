package off

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		var level []int
		for i := 0; i < size; i++ {
			poll := q[0]
			q = q[1:]
			level = append(level, poll.Val)
			if poll.Left != nil {
				q = append(q, poll.Left)
			}
			if poll.Right != nil {
				q = append(q, poll.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}
