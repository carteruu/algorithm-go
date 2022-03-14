package offer

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dh := &ListNode{}
	dt := dh
	for l1 != nil && l2 != nil {
		if l2.Val < l1.Val {
			dt.Next = l2
			l2 = l2.Next
		} else {
			dt.Next = l1
			l1 = l1.Next
		}
		dt = dt.Next
	}
	if l1 != nil {
		dt.Next = l1
	}
	if l2 != nil {
		dt.Next = l2
	}
	return dh.Next
}
