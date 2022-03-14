package leetcode

import "algorithm/pkg/link"

func treeToDoublyList(root *link.Node) *link.Node {
	if root == nil {
		return nil
	}
	//深度优先搜索,将二叉搜索树转换为排序链表,并返回该链表的头尾节点
	var dfs func(node *link.Node) (*link.Node, *link.Node)
	dfs = func(node *link.Node) (*link.Node, *link.Node) {
		if node == nil {
			//树为空,返回空链表
			return nil, nil
		}
		//递归转换左右子树
		var hl, tl, hr, tr *link.Node
		if node.Left != nil {
			hl, tl = dfs(node.Left)
		}
		if node.Right != nil {
			hr, tr = dfs(node.Right)
		}
		//拼接左子树链表-当前节点-右子树链表
		if hl != nil {
			tl.Right = node
			node.Left = tl
		} else {
			hl = node
		}
		if hr != nil {
			node.Right = hr
			hr.Left = node
		} else {
			tr = node
		}
		//返回新的头尾节点
		return hl, tr
	}
	head, tail := dfs(root)
	//将头尾节点连成环
	tail.Right = head
	head.Left = tail
	//返回头结点
	return head
}
