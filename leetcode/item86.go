package leetcode

func partition(head *ListNode, x int) *ListNode {
	lh, rh := &ListNode{}, &ListNode{}
	lt, rt := lh, rh
	for head != nil {
		if head.Val < x {
			lt.Next = head
			lt = lt.Next
		} else {
			rt.Next = head
			rt = rt.Next
		}
		head = head.Next
	}
	if rt.Next != nil {
		rt.Next = nil
	}
	lt.Next = rh.Next
	return lh.Next
}
