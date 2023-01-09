package leetcode_2

func middleNode(head *ListNode) *ListNode {
	h := head
	for h != nil && h.Next != nil {
		h = h.Next.Next
		head = head.Next
	}
	return head
}
