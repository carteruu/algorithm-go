package leet

import (
	"algorithm/pkg/link"
)

func treeToDoublyList2(root *link.Node) *link.Node {
	if root == nil {
		return nil
	}
	//虚拟头尾结点
	dummyHead := &link.Node{}
	dummyTail := dummyHead
	//中序遍历
	var dfs func(node *link.Node)
	dfs = func(node *link.Node) {
		if node == nil {
			return
		}
		dfs(node.Left)
		//拼接节点
		dummyTail.Right = node
		node.Left = dummyTail
		dummyTail = dummyTail.Right
		dfs(node.Right)
	}
	dfs(root)
	//拼接头尾节点,形成环
	head := dummyHead.Right
	head.Left = dummyTail
	dummyTail.Right = head
	return head
}
func treeToDoublyList1(root *link.Node) *link.Node {
	if root == nil {
		return nil
	}
	var head *link.Node
	var tail *link.Node
	var dfs func(node *link.Node)
	dfs = func(node *link.Node) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if head == nil {
			head = node
		} else {
			tail.Right = node
			node.Left = tail
		}
		tail = node
		dfs(node.Right)
	}
	dfs(root)
	head.Left = tail
	tail.Right = head
	return head
}
