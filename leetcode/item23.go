package leetcode

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	var nodeHeap listNodeHeap
	heap.Init(&nodeHeap)
	for _, node := range lists {
		if node != nil {
			heap.Push(&nodeHeap, node)
		}
	}
	dummyHead := &ListNode{}
	dummyTail := dummyHead
	for nodeHeap.Len() > 0 {
		dummyTail.Next = heap.Pop(&nodeHeap).(*ListNode)
		if dummyTail.Next.Next != nil {
			heap.Push(&nodeHeap, dummyTail.Next.Next)
		}
		dummyTail = dummyTail.Next
	}
	return dummyHead.Next
}

type listNodeHeap []*ListNode

func (h listNodeHeap) Len() int {
	return len(h)
}
func (h listNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h listNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *listNodeHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func (h *listNodeHeap) Push(node interface{}) {
	*h = append(*h, node.(*ListNode))
}
