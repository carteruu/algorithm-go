package leetcode

func middleNode1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList1(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode1(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList1(l2)
	mergeList(l1, l2)
}

func reorderList3(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
func reorderList2(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	h := head
	var stack []*ListNode
	for h != nil {
		stack = append(stack, h)
		h = h.Next
	}
	stack = stack[1:]
	cur := head
	for len(stack) > 0 {
		cur.Next = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Next
		if len(stack) == 0 {
			break
		}
		cur.Next = stack[0]
		stack = stack[1:]
		cur = cur.Next
	}
	cur.Next = nil
}
func reorderList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	n := 0
	h := head
	var stack []*ListNode
	for h != nil {
		n++
		stack = append(stack, h)
		h = h.Next
	}
	cur := head
	for {
		if n == 0 {
			break
		}
		next := cur.Next
		cur.Next = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		n--
		cur = cur.Next

		if n == 0 {
			break
		}
		cur.Next = next
		cur = next
		n--
	}
	cur.Next = nil
	return head
}

//func reorderList1(head *ListNode) {
//	slow, fast := head, head
//	n2 := 0
//	isEven := true
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//		if fast == nil {
//			isEven = false
//		}
//		n2++
//	}
//	stack := make([]*ListNode, n2)
//
//}
