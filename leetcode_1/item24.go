package leetcode_1

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	for head != nil && head.Next != nil {
		node3 := head.Next.Next
		dummyTail.Next = head.Next
		dummyTail.Next.Next = head
		dummyTail = head
		head = node3
	}
	dummyTail.Next = head
	return dummyHead.Next
}
