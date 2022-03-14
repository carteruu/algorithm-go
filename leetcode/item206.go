package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	next := head.Next
	head.Next = nil
	for next != nil {
		nextNext := next.Next

		next.Next = head
		head = next
		next = nextNext
	}
	return head
}
