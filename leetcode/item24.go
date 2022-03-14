package leetcode

func swapPairs11(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	for head != nil && head.Next != nil {
		three := head.Next.Next
		dummyTail.Next = head.Next
		dummyTail.Next.Next = head
		dummyTail = dummyTail.Next.Next
		head = three
	}
	dummyTail.Next = head
	return dummyHead.Next
}

func swapPairs1(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs11(newHead.Next)
	newHead.Next = head
	return newHead
}
