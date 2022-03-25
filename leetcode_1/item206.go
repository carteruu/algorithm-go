package leetcode_1

func reverseList_1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}
	return prev
}
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
