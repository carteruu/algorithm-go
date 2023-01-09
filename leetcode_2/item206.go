package leetcode_2

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	var s []*ListNode
	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	s[0].Next = nil
	for i := len(s) - 1; i >= 0; i-- {
		dummyTail.Next = s[i]
		dummyTail = dummyTail.Next
	}
	return dummyHead.Next
}

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
