package leet

func deleteDuplicates11(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			num := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == num {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	dummyTail.Next = head
	head = head.Next
	repeat := false
	for head != nil {
		if head.Val == dummyTail.Next.Val {
			repeat = true
		} else {
			if !repeat {
				dummyTail = dummyTail.Next
			}
			dummyTail.Next = head
			repeat = false
		}
		head = head.Next
	}
	if repeat {
		dummyTail.Next = nil
	}
	return dummyHead.Next
}
func deleteDuplicates12(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	next := head
	cur := head.Next
	rep := false
	for cur != nil {
		if next.Val == cur.Val {
			rep = true
		} else {
			if !rep {
				dummyTail.Next = next
				dummyTail = dummyTail.Next
			}
			next = cur
			rep = false
		}
		cur = cur.Next
	}
	if !rep {
		dummyTail.Next = next
	} else {
		dummyTail.Next = nil
	}
	return dummyHead.Next
}
