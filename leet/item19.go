package leet

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	var pre *ListNode
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}
	if pre == nil {
		return head.Next
	}
	pre.Next = slow.Next
	return head
}
