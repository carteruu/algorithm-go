package leetcode

func oddEvenList11(head *ListNode) *ListNode {
	if head == nil && head.Next == nil {
		return head
	}
	evenHead := &ListNode{}
	evenTail := evenHead
	curr := head
	var prev *ListNode
	for curr != nil && curr.Next != nil {
		evenTail.Next = curr.Next
		evenTail = evenTail.Next

		prev = curr
		curr.Next = curr.Next.Next
		curr = curr.Next
	}
	evenTail.Next = nil
	if prev.Next != nil {
		prev = prev.Next
	}
	prev.Next = evenHead.Next
	return head
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
