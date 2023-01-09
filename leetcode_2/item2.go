package leetcode_2

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	a := 0
	for l1 != nil || l2 != nil {
		val := a
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		a = 0
		if val >= 10 {
			a = 1
			val -= 10
		}
		dummy.Next = &ListNode{Val: val}
		dummy = dummy.Next
	}
	if a > 0 {
		dummy.Next = &ListNode{Val: a}
	}
	return head.Next
}
