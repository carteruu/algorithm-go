package item102

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		level := make([]int, 0, size)
		ans = append(ans, level)
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
	}
	return ans
}
