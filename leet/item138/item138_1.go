package item138

import (
	"algorithm/pkg/random"
)

func copyRandomList3(head *random.Node) *random.Node {
	m := make(map[*random.Node]*random.Node)
	dummyHead := &random.Node{}
	dummyTail := dummyHead
	cur := head
	for cur != nil {
		dummyTail.Next = &random.Node{
			Val: cur.Val,
		}
		dummyTail = dummyTail.Next
		m[cur] = dummyTail
		cur = cur.Next
	}

	cur = head
	nCur := dummyHead.Next
	for nCur != nil {
		nCur.Random = m[cur.Random]
		nCur = nCur.Next
		cur = cur.Next
	}
	return dummyHead.Next
}

var cachedNode map[*random.Node]*random.Node

func deepCopy(node *random.Node) *random.Node {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &random.Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList2(head *random.Node) *random.Node {
	cachedNode = map[*random.Node]*random.Node{}
	return deepCopy(head)
}

func copyRandomList1(head *random.Node) *random.Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &random.Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}
