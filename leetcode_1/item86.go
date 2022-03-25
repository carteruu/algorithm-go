package leetcode_1

//将小于x的节点挑出来
//需要分两段遍历，第一次，头节点小于x，需要换头；第二次，不需要换头
func partition_1(head *ListNode, x int) *ListNode {
	//存小于x的节点
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	for head != nil && head.Val < x {
		dummyTail.Next = head
		dummyTail = head
		head = head.Next
	}
	var prev, curr *ListNode = nil, head
	for curr != nil {
		if curr.Val < x {
			dummyTail.Next = curr
			dummyTail = curr
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	dummyTail.Next = head
	return dummyHead.Next
}

//分别把节点放进两个链表，再合并
func partition(head *ListNode, x int) *ListNode {
	dummyHeadSmall := &ListNode{}
	dummyTailSmall := dummyHeadSmall

	dummyHeadBig := &ListNode{}
	dummyTailBig := dummyHeadBig

	for head != nil {
		if head.Val < x {
			dummyTailSmall.Next = head
			dummyTailSmall = head
		} else {
			dummyTailBig.Next = head
			dummyTailBig = head
		}
		head = head.Next
	}
	dummyTailSmall.Next = dummyHeadBig.Next
	//注意这里要把最后一个节点的 Next 属性清空
	dummyTailBig.Next = nil
	return dummyHeadSmall.Next
}
