package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	p := dummyHead
	for i := 0; i < left-1; i++ {
		p = p.Next
	}
	cur := p.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next

		next.Next = p.Next
		p.Next = next
	}
	return dummyHead.Next
}

func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	idx := 1
	for head != nil {
		dummyTail.Next = head
		if idx == left {
			for idx != right {
				head = head.Next
				idx++
			}
			tail := head
			head = head.Next
			idx++
			newHead, newTail := rev(dummyTail.Next, tail)
			dummyTail.Next = newHead
			dummyTail = newTail
			continue
		}
		dummyTail = dummyTail.Next
		head = head.Next
		idx++
	}
	return dummyHead.Next
}

func rev(head, tail *ListNode) (*ListNode, *ListNode) {
	cur := head
	end := tail.Next
	var pre *ListNode
	for cur != end {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	return tail, head
}
