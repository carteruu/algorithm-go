package offer

type CQueue struct {
	q []int
}

func Constructor() CQueue {
	return CQueue{q: make([]int, 0, 8)}
}

func (this *CQueue) AppendTail(value int) {
	this.q = append(this.q, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.q) == 0 {
		return -1
	}
	val := this.q[0]
	this.q = this.q[1:]
	return val
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
