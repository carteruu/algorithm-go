package leet

import "algorithm/pkg"

func mergeTrees(root1 *pkg.TreeNode, root2 *pkg.TreeNode) *pkg.TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 != nil && root2 != nil {
		root1.Val = root1.Val + root2.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
	}
	if root1 != nil {
		return root1
	} else {
		return root2
	}
}
