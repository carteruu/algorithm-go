package offer

import (
	"container/heap"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type listNodeHeap []*ListNode

func (l listNodeHeap) Len() int {
	return len(l)
}
func (l listNodeHeap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l listNodeHeap) Less(i, j int) bool {
	return l[i].Val < l[j].Val
}
func (l *listNodeHeap) Push(list interface{}) {
	*l = append(*l, list.(*ListNode))
}
func (l *listNodeHeap) Pop() interface{} {
	x := (*l)[l.Len()-1]
	*l = (*l)[:l.Len()-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	l := make(listNodeHeap, 0, len(lists))
	heap.Init(&l)
	for _, list := range lists {
		if list != nil {
			heap.Push(&l, list)
		}
	}
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	for l.Len() > 0 {
		pop := heap.Pop(&l).(*ListNode)
		dummyTail.Next = pop
		dummyTail = dummyTail.Next
		pop = pop.Next
		if pop != nil {
			heap.Push(&l, pop)
		}
	}
	return dummyHead.Next
}
func mergeKLists1(lists []*ListNode) *ListNode {
	l := make([]*ListNode, 0, len(lists))
	for _, list := range lists {
		if list != nil {
			l = append(l, list)
		}
	}
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	less := func(i, j int) bool {
		return l[i].Val < l[j].Val
	}
	for len(l) > 0 {
		sort.Slice(l, less)
		dummyTail.Next = l[0]
		dummyTail = dummyTail.Next
		l[0] = l[0].Next
		if l[0] == nil {
			l = l[1:]
		}
	}
	return dummyHead.Next
}
