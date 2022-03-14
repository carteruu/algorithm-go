package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	cnt := 1
	for cur.Next != nil {
		cnt++
		cur = cur.Next
	}
	k %= cnt
	if k == 0 {
		return head
	}
	//连成环
	cur.Next = head

	var pre *ListNode
	for i := 0; i < cnt-k; i++ {
		pre = head
		head = head.Next
	}
	pre.Next = nil
	return head
}
func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	//计算列表长度
	cur := head
	cnt := 0
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	//计算实际需要移动的长度
	k %= cnt
	if k == 0 {
		return head
	}
	//使用快慢指针获取需要移动到头部的节点,及其前一个节点,及尾结点
	slow, fast := head, head
	for i := 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head
	}
	var pre *ListNode
	var tail *ListNode
	for fast != nil {
		tail = fast
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}
	//旋转
	pre.Next = nil
	tail.Next = head
	return slow
}
