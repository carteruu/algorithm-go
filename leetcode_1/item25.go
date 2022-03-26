package leetcode_1

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	curr := head
	for curr != nil {
		h := curr
		need := true
		for i := 0; i < k; i++ {
			if curr == nil {
				need = false
				break
			}
			curr = curr.Next
		}
		if !need {
			dummyTail.Next = h
			break
		}
		newHead := reverse(h, curr)
		dummyTail.Next = newHead
		dummyTail = h
	}
	return dummyHead.Next
}
func reverse(h, t *ListNode) *ListNode {
	var prev *ListNode
	for h != t {
		h.Next, prev, h = prev, h, h.Next
	}
	return prev
}
