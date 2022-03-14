package leetcode

import "algorithm/pkg"

func isSymmetric(root *pkg.TreeNode) bool {
	if root == nil {
		return true
	}
	var q []*pkg.TreeNode
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size/2; i++ {
			if q[i] == nil && q[size-1-i] == nil {
				continue
			}
			if q[i] == nil || q[size-1-i] == nil {
				return false
			}
			if q[i].Val != q[size-1-i].Val {
				return false
			}
		}
		for i := 0; i < size; i++ {
			poll := q[0]
			q = q[1:]
			if poll != nil {
				q = append(q, poll.Left, poll.Right)
			}
		}
	}
	return true
}
func isSymmetric1(root *pkg.TreeNode) bool {
	var dfs func(node1, node2 *pkg.TreeNode) bool
	dfs = func(node1, node2 *pkg.TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		return dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
	}
	return dfs(root, root)
}
