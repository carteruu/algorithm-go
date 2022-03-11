package leet

import "algorithm/pkg"

func inorderSuccessor1(root *pkg.TreeNode, p *pkg.TreeNode) *pkg.TreeNode {
	find := false
	var ans *pkg.TreeNode
	var dfs func(node *pkg.TreeNode)
	dfs = func(node *pkg.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if find {
			ans = node
			find = false
		}
		if node == p {
			find = true
		}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func inorderSuccessor(root *pkg.TreeNode, p *pkg.TreeNode) *pkg.TreeNode {
	if root == nil || p == nil {
		return nil
	}
	if root.Val > p.Val {
		node := inorderSuccessor(root.Left, p)
		if node == nil {
			return root
		} else {
			return node
		}
	} else {
		node := inorderSuccessor(root.Right, p)
		return node
	}
}
