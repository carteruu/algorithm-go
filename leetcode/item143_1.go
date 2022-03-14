package leetcode

func reorderList4(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	s, f := head, head.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	var prev *ListNode
	curr := s.Next
	s.Next = nil
	for curr != nil {
		n := curr.Next
		curr.Next = prev
		prev = curr
		curr = n
	}
	dummy := &ListNode{}
	for head != nil && prev != nil {
		dummy.Next = head
		dummy = dummy.Next
		head = head.Next

		dummy.Next = prev
		dummy = dummy.Next
		prev = prev.Next
	}
	if head != nil {
		dummy.Next = head
	}
	if prev != nil {
		dummy.Next = prev
	}
}
