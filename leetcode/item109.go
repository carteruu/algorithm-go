package leetcode

import "algorithm/pkg"

func sortedListToBST(head *ListNode) *pkg.TreeNode {
	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	var builtBst func(l, r int) *pkg.TreeNode
	builtBst = func(l, r int) *pkg.TreeNode {
		if l > r {
			return nil
		}
		node := &pkg.TreeNode{}
		mid := (l + r) / 2
		node.Left = builtBst(l, mid-1)
		node.Val = head.Val
		head = head.Next
		node.Right = builtBst(mid+1, r)
		return node
	}
	return builtBst(0, length-1)
}
func sortedListToBST1(head *ListNode) *pkg.TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &pkg.TreeNode{Val: head.Val}
	}
	s, f := head, head
	var prev *ListNode
	for f != nil && f.Next != nil {
		prev = s
		s = s.Next
		f = f.Next.Next
	}
	prev.Next = nil
	return &pkg.TreeNode{
		Val:   s.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(s.Next),
	}
}

var globalHead *ListNode

func sortedListToBST2(head *ListNode) *pkg.TreeNode {
	globalHead = head
	length := getLength(head)
	return buildTree(0, length-1)
}

func getLength(head *ListNode) int {
	ret := 0
	for ; head != nil; head = head.Next {
		ret++
	}
	return ret
}

func buildTree(left, right int) *pkg.TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2
	root := &pkg.TreeNode{}
	root.Left = buildTree(left, mid-1)
	root.Val = globalHead.Val
	globalHead = globalHead.Next
	root.Right = buildTree(mid+1, right)
	return root
}
