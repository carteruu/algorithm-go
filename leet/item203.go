package leet

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
