package leet

import "algorithm/pkg"

func isSameTree(p *pkg.TreeNode, q *pkg.TreeNode) bool {
	if p == q {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
