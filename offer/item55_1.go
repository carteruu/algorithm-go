package offer

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
func maxDepth1(root *TreeNode) int {
	arr := make([]*TreeNode, 0)
	arr = append(arr, root)
	res := 0
	for len(arr) > 0 {
		size := len(arr)
		res++
		for i := 0; i < size; i++ {
			poll := arr[0]
			arr = arr[1:]
			if poll.Left != nil {
				arr = append(arr, poll.Left)
			}
			if poll.Right != nil {
				arr = append(arr, poll.Right)
			}
		}
	}
	return res
}
