package leetcode

func isPalindrome4(head *ListNode) bool {
	//将前半部分翻转,pre将为左半部分头节点,slow将为右半部分头结点
	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}
	if fast != nil {
		//节点数为奇数时,需要跳过正中间的节点
		slow = slow.Next
	}
	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}
	return true
}
func isPalindrome11(head *ListNode) bool {
	var q []int
	for head != nil {
		q = append(q, head.Val)
		head = head.Next
	}
	n := len(q)
	for i := 0; i < n/2; i++ {
		if q[i] != q[n-1-i] {
			return false
		}
	}
	return true
}
