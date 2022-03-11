package leet

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	dummyTail := head
	for cur != nil {
		if cur.Val != dummyTail.Val {
			dummyTail.Next = cur
			dummyTail = dummyTail.Next
		}
		cur = cur.Next
	}
	dummyTail.Next = nil
	return head
}
