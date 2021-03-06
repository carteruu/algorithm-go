package offer

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre := head
	cur := pre.Next
	pre.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}
