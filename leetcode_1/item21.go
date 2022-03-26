package leetcode_1

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
func mergeTwoLists_1(list1 *ListNode, list2 *ListNode) *ListNode {
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
