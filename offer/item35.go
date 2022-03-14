package offer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for curr := head; curr != nil; curr = curr.Next.Next {
		n := &Node{Val: curr.Val}
		n.Next = curr.Next
		curr.Next = n
	}
	for curr := head; curr != nil; curr = curr.Next.Next {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
	}
	nHead := head.Next
	for curr := head; curr != nil; curr = curr.Next {
		n := curr.Next.Next
		if curr.Next.Next != nil {
			curr.Next.Next = curr.Next.Next.Next
		}
		curr.Next = n
	}
	return nHead
}
