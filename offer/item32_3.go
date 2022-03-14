package offer

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	isLeft := true
	for len(queue) > 0 {
		size := len(queue)
		r := make([]int, size, size)
		for i := 0; i < size; i++ {
			var poll *TreeNode
			poll = queue[0]
			queue = queue[1:]
			if isLeft {
				r[i] = poll.Val
			} else {
				r[size-1-i] = poll.Val
			}
			if poll.Left != nil {
				queue = append(queue, poll.Left)
			}
			if poll.Right != nil {
				queue = append(queue, poll.Right)
			}
		}
		res = append(res, r)
		isLeft = !isLeft
	}
	return res
}
