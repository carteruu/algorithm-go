package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	for head != nil && head.Next != nil {
		//保存下一个节点和下下个节点
		three := head.Next.Next
		//连接下个节点
		dummyTail.Next = head.Next
		//连接当前节点
		dummyTail.Next.Next = head
		dummyTail = dummyTail.Next.Next
		//当前节点移到下下个节点
		head = three
	}
	//节点数为单数时,连接最后一个节点;双数时,下下个节点为nil,没有影响
	dummyTail.Next = head
	return dummyHead.Next
}
