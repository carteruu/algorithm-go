package leet

import "algorithm/pkg"

func flatten(root *pkg.TreeNode) {
	if root == nil {
		return
	}
	stack := []*pkg.TreeNode{root}
	prev := &pkg.TreeNode{}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		prev.Left, prev.Right = nil, curr
		prev = curr

		left, right := curr.Left, curr.Right
		if right != nil {
			stack = append(stack, right)
		}
		if left != nil {
			stack = append(stack, left)
		}
	}
}

func flatten2(root *pkg.TreeNode) {
	if root == nil {
		return
	}
	right := root.Right
	flatten(root.Left)
	root.Right = root.Left
	root.Left = nil

	flatten(right)
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
}
func flatten1(root *pkg.TreeNode) {
	var dfs func(node *pkg.TreeNode) *pkg.TreeNode
	dfs = func(node *pkg.TreeNode) *pkg.TreeNode {
		if node == nil {
			return nil
		}
		//记录右节点
		right := node.Right
		//将左子树移动到右子树,并清空左子树
		node.Right = dfs(node.Left)
		node.Left = nil
		//找到最后一个非空右子节点,并将右子树拼上去
		tail := node
		for tail.Right != nil {
			tail = tail.Right
		}
		tail.Right = dfs(right)
		//返回当前节点
		return node
	}
	dfs(root)
}
