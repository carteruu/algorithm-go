package leetcode_2

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			dummyTail.Next = list1
			list1 = list1.Next
		} else {
			dummyTail.Next = list2
			list2 = list2.Next
		}
		dummyTail = dummyTail.Next
	}
	if list1 != nil {
		dummyTail.Next = list1
	}
	if list2 != nil {
		dummyTail.Next = list2
	}
	return dummyHead.Next
}
