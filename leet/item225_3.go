package leet

type MyStack struct {
	q, q1 []int
}

func ConstructorMyStack() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	size := len(this.q)
	this.q = append(this.q, x)
	for i := 0; i < size; i++ {
		this.q = append(this.q, this.q[0])
		this.q = this.q[1:]
	}
}

func (this *MyStack) Pop() int {
	if len(this.q) == 0 {
		return -1
	}
	pop := this.q[0]
	this.q = this.q[1:]
	return pop
}

func (this *MyStack) Top() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.q[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}
