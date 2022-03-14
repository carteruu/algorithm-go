package leetcode

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sign := 0
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	for l1 != nil || l2 != nil {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		val := val1 + val2 + sign
		sign = val / 10
		dummyTail.Next = &ListNode{Val: val % 10}
		dummyTail = dummyTail.Next
	}
	if sign > 0 {
		dummyTail.Next = &ListNode{Val: sign}
	}
	return dummyHead.Next
}
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	cur1 := l1
	cur2 := l2
	sign := 0
	for cur1 != nil && cur2 != nil {
		val := cur1.Val + cur2.Val + sign
		sign = val / 10
		cur1.Val = val % 10

		if cur1.Next == nil && cur2.Next == nil {
			if sign > 0 {
				cur1.Next = &ListNode{Val: sign}
			}
			break
		}
		if cur1.Next == nil {
			if sign > 0 {
				cur1.Next = &ListNode{Val: sign}
				sign = 0
			} else {
				cur1.Next = cur2.Next
				break
			}
		}
		if cur2.Next == nil {
			if sign > 0 {
				cur2.Next = &ListNode{Val: sign}
				sign = 0
			} else {
				break
			}
		}

		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return l1
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sign := 0
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	for l1 != nil || l2 != nil {
		val := sign
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		sign = val / 10
		dummyTail.Next = &ListNode{Val: val % 10}
		dummyTail = dummyTail.Next
	}
	if sign > 0 {
		dummyTail.Next = &ListNode{Val: sign}
	}
	return dummyHead.Next
}
