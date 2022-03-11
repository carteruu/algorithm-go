package off

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		r := make([]int, 0, size)
		for i := 0; i < size; i++ {
			poll := queue[0]
			r = append(r, poll.Val)
			queue = queue[1:]
			if poll.Left != nil {
				queue = append(queue, poll.Left)
			}
			if poll.Right != nil {
				queue = append(queue, poll.Right)
			}
		}
		res = append(res, r)
	}
	return res
}
