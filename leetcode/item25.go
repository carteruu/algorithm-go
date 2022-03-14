package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	var reverseK func(h *ListNode, k int) (newHead, newTail *ListNode)
	reverseK = func(head *ListNode, k int) (newHead, newTail *ListNode) {
		cur := head
		for i := 0; i < k; i++ {
			if cur == nil {
				//长度不够,不需要翻转
				return head, nil
			}
			cur = cur.Next
		}

		var pre *ListNode
		cur = head
		for i := 0; i < k; i++ {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		nextHead, nextTail := reverseK(cur, k)
		head.Next = nextHead
		return pre, nextTail
	}

	newHead, _ := reverseK(head, k)
	return newHead
}
