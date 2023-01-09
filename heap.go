package algorithm

import "container/heap"

func NewBigIntHeap(len int) *BigHeap {
	h := BigHeap(make([]int, 0, len))
	heap.Init(&h)
	return &h
}

type BigHeap []int

func (h BigHeap) Len() int {
	return len(h)
}

func (h BigHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h BigHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *BigHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *BigHeap) Pop() interface{} {
	if h.Len() == 0 {
		panic("is empty")
	}
	pos := len(*h) - 1
	node := (*h)[pos]
	*h = (*h)[:pos]
	return node
}
