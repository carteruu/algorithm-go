package leet

import "algorithm/pkg"

func upsideDownBinaryTree(root *pkg.TreeNode) *pkg.TreeNode {
	left := root
	for {
		if left.Left == nil {
			break
		}
		left = left.Left
	}

	var odr func(node *pkg.TreeNode) *pkg.TreeNode
	odr = func(node *pkg.TreeNode) *pkg.TreeNode {
		if node == nil {
			return nil
		}
		left := odr(node.Left)
		right := odr(node.Right)
		if left != nil {
			left.Right = node
			left.Left = right
		}
		return node
	}
	odr(root)

	return left
}
func upsideDownBinaryTree2(root *pkg.TreeNode) *pkg.TreeNode {
	var left, right *pkg.TreeNode
	for root != nil {
		temp := root.Left

		root.Left = right
		right = root.Right
		root.Right = left

		left = root
		root = temp
	}
	return left
}
func upsideDownBinaryTree1(root *pkg.TreeNode) *pkg.TreeNode {
	var father, left, right *pkg.TreeNode
	for root != nil {
		left = root.Left
		root.Left = right
		right = root.Right
		root.Right = father
		father = root
		root = left
	}
	return father
}
