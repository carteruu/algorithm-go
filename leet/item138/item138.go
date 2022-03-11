package item138

import (
	"algorithm/pkg/random"
)

func copyRandomList(head *random.Node) *random.Node {
	m := make(map[*random.Node]*random.Node)
	var newHead *random.Node
	var newPre *random.Node
	for node := head; node != nil; {
		newNode := &random.Node{
			Val:    node.Val,
			Random: node.Random,
		}
		if newHead == nil {
			newHead = newNode
		}

		m[node] = newNode
		if newPre != nil {
			newPre.Next = newNode
		}
		newPre = newNode
		node = node.Next
	}
	newNode := newHead
	for newNode != nil {
		newNode.Random = m[newNode.Random]
		newNode = newNode.Next
	}
	return newHead
}
