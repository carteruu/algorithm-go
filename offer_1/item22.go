package offer_1

func getKthFromEnd(head *ListNode, k int) *ListNode {
	head1, head2 := head, head
	for i := 0; i < k; i++ {
		head1 = head1.Next
	}
	for head1 != nil {
		head1 = head1.Next
		head2 = head2.Next
	}
	return head2
}
