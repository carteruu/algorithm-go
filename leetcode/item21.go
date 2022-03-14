package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			dummyTail.Next = l1
			l1 = l1.Next
		} else {
			dummyTail.Next = l2
			l2 = l2.Next
		}
		dummyTail = dummyTail.Next
	}
	if l1 != nil {
		dummyTail.Next = l1
	} else {
		dummyTail.Next = l2
	}
	return dummyHead.Next
}
