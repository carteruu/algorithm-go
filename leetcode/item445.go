package leetcode

func addTwoNumbers445(l1 *ListNode, l2 *ListNode) *ListNode {
	var arr1, arr2 []int
	for l1 != nil {
		arr1 = append(arr1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		arr2 = append(arr2, l2.Val)
		l2 = l2.Next
	}
	var dummyHead *ListNode
	sign := 0
	for i := 0; i < len(arr1) || i < len(arr2) || sign > 0; i++ {
		val1 := 0
		val2 := 0
		if i < len(arr1) {
			val1 = arr1[len(arr1)-1-i]
		}
		if i < len(arr2) {
			val2 = arr2[len(arr2)-1-i]
		}
		val := val1 + val2 + sign
		sign = val / 10
		dummyHead = &ListNode{
			Val:  val % 10,
			Next: dummyHead,
		}
	}
	return dummyHead
}
func addTwoNumbers445_2(l1 *ListNode, l2 *ListNode) *ListNode {
	var num1, num2 int
	for l1 != nil {
		num1 *= 10
		num1 += l1.Val
		l1 = l1.Next
	}
	for l2 != nil {
		num2 *= 10
		num2 += l2.Val
		l2 = l2.Next
	}
	num1 += num2
	if num1 == 0 {
		return &ListNode{
			Val: 0,
		}
	}
	var dummyHead *ListNode
	for num1 > 0 {
		cur := &ListNode{
			Val:  num1 % 10,
			Next: dummyHead,
		}
		num1 /= 10
		dummyHead = cur
	}
	return dummyHead
}
func addTwoNumbers445_1(l1 *ListNode, l2 *ListNode) *ListNode {
	var rev func(l *ListNode) *ListNode
	rev = func(l *ListNode) *ListNode {
		if l == nil {
			return nil
		}
		ll := l.Next
		l.Next = nil
		for ll != nil {
			lll := ll.Next
			ll.Next = l

			l = ll
			ll = lll
		}
		return l
	}
	var addTwoNumbers func(l1 *ListNode, l2 *ListNode) *ListNode
	addTwoNumbers = func(l1 *ListNode, l2 *ListNode) *ListNode {
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
	l1 = rev(l1)
	l2 = rev(l2)
	ans := addTwoNumbers(l1, l2)
	return rev(ans)
}
