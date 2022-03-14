package leetcode

type MyStack11 struct {
	q1 []int
	q2 []int
}

func ConstructorMyStack11() MyStack11 {
	return MyStack11{}
}

func (this *MyStack11) Push(x int) {
	if len(this.q2) > 0 {
		this.q2 = append(this.q2, x)
		return
	}
	this.q1 = append(this.q1, x)
}

func (this *MyStack11) Pop() int {
	if len(this.q1) > 0 {
		lastIdx := len(this.q1) - 1
		for i := 0; i < lastIdx; i++ {
			this.q2 = append(this.q2, this.q1[0])
			this.q1 = this.q1[1:]
		}
		pop := this.q1[0]
		this.q1 = this.q1[1:]
		return pop
	}
	if len(this.q2) > 0 {
		lastIdx := len(this.q2) - 1
		for i := 0; i < lastIdx; i++ {
			this.q1 = append(this.q1, this.q2[0])
			this.q2 = this.q2[1:]
		}
		pop := this.q2[0]
		this.q2 = this.q2[1:]
		return pop
	}
	return 0
}

func (this *MyStack11) Top() int {
	pop := this.Pop()
	this.Push(pop)
	return pop
}

func (this *MyStack11) Empty() bool {
	return len(this.q1) == 0 && len(this.q2) == 0
}
