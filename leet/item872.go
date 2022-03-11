package leet

import "algorithm/pkg"

func leafSimilar(root1 *pkg.TreeNode, root2 *pkg.TreeNode) bool {
	leaf := make([]int, 0, 500)
	idx := 0
	var inorder func(node *pkg.TreeNode, isFirst bool) bool
	inorder = func(node *pkg.TreeNode, isFirst bool) bool {
		if node == nil {
			return true
		}
		if node.Left == nil && node.Right == nil {
			if isFirst {
				leaf = append(leaf, node.Val)
				return true
			} else {
				if len(leaf) < idx+1 || leaf[idx] != node.Val {
					return false
				}
				idx++
			}
		}

		return inorder(node.Left, isFirst) && inorder(node.Right, isFirst)
	}
	inorder(root1, true)
	if !inorder(root2, false) {
		return false
	}
	return len(leaf) == idx
}
