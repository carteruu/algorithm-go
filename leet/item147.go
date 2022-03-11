package leet

func insertionSortList(head *ListNode) *ListNode {
	newHead := &ListNode{}
	for cur := head; cur != nil; {
		prev := newHead
		for prev.Next != nil && cur.Val >= prev.Next.Val {
			prev = prev.Next
		}
		next := cur.Next
		cur.Next = prev.Next
		prev.Next = cur
		cur = next
	}
	return newHead.Next
}
func insertionSortList1(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	lastSorted := head
	for curr := lastSorted.Next; curr != nil; {
		prev := newHead
		for curr.Val >= prev.Next.Val {
			prev = prev.Next
		}
		lastSorted.Next = curr.Next
		curr.Next = prev.Next
		prev.Next = curr
		curr = lastSorted.Next
	}
	return newHead.Next
}

func insertionSortList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
