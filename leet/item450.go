package leet

import "algorithm/pkg"

func deleteNode1(root *pkg.TreeNode, key int) *pkg.TreeNode {
	var inorder func(node *pkg.TreeNode) *pkg.TreeNode
	inorder = func(node *pkg.TreeNode) *pkg.TreeNode {
		if node == nil {
			return nil
		}
		if node.Val == key {
			if node.Right == nil {
				return node.Left
			}
			if node.Left == nil {
				return node.Right
			}

			parent := node.Right
			newNode := node.Right.Left
			if newNode == nil {
				node.Right.Left = node.Left
				return parent
			}

			isLeft := true
			for newNode.Left != nil && newNode.Right != nil {
				parent = newNode
				if newNode.Left != nil {
					newNode = newNode.Left
					isLeft = true
				} else {
					newNode = newNode.Right
					isLeft = false
				}
			}
			newNode.Left = node.Left
			newNode.Right = node.Right
			if isLeft {
				parent.Left = nil
			} else {
				parent.Right = nil
			}

			return newNode
		}
		if key > node.Val {
			node.Right = inorder(node.Right)
		} else if key < node.Val {
			node.Left = inorder(node.Left)
		}
		return node
	}
	node := inorder(root)
	return node
}
