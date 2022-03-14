package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	cnt := 0
	tail := head
	for tail != nil {
		cnt++
		if cnt == 1 {
			head = tail
			dummyTail.Next = tail
		}
		next := tail.Next
		if cnt == k {
			newHead, newTail := reverseK(head, next)
			dummyTail.Next = newHead
			dummyTail = newTail
			cnt = 0
		}
		tail = next
	}
	return dummyHead.Next
}

/**
不包含tail
*/
func reverseK(head, tail *ListNode) (*ListNode, *ListNode) {
	var newHead *ListNode
	cur := head
	for cur != tail {
		next := cur.Next
		cur.Next = newHead

		newHead = cur
		cur = next
	}
	return newHead, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
