package off

func levelOrder2(root *TreeNode) []int {
	var queue []*TreeNode
	queue = append(queue, root)
	var res []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			res = append(res, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}
	return res
}
func levelOrder1(root *TreeNode) []int {
	queue := make(chan *TreeNode, 500)
	queue <- root

	var res []int
	for {
		select {
		case node := <-queue:
			if node == nil {
				continue
			}
			res = append(res, node.Val)
			queue <- node.Left
			queue <- node.Right
		default:
			return res
		}
	}
}
