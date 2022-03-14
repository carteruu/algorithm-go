package leetcode

func plusOne11(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	c := 1
	curr = prev
	prev = nil
	for curr != nil {
		val := curr.Val
		curr.Val = (val + c) % 10
		c = (val + c) / 10

		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	if c == 1 {
		return &ListNode{1, head}
	} else {
		return head
	}
}
