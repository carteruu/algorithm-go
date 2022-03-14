package offer

func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var path func(node *TreeNode, t int, p []int)
	path = func(node *TreeNode, t int, p []int) {
		//选择
		p = append(p, node.Val)
		if t == node.Val && node.Left == nil && node.Right == nil {
			rs := make([]int, len(p))
			copy(rs, p)
			res = append(res, rs)
			return
		}
		tn := t - node.Val
		if node.Left != nil {
			path(node.Left, tn, p)
		}
		if node.Right != nil {
			path(node.Right, tn, p)
		}
		//撤销选择
		p = p[:len(p)-1]
	}

	if root == nil {
		return res
	}
	path(root, target, []int{})
	return res
}
