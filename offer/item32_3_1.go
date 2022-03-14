package offer

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	byLeft := true
	var ans [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			poll := q[0]
			q = q[1:]
			if byLeft {
				level[i] = poll.Val
			} else {
				level[size-1-i] = poll.Val
			}
			if poll.Left != nil {
				q = append(q, poll.Left)
			}
			if poll.Right != nil {
				q = append(q, poll.Right)
			}
		}
		byLeft = !byLeft
		ans = append(ans, level)
	}
	return ans
}
