package off

type MaxQueue struct {
	q []int
	//单调队列
	mq []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.mq) == 0 {
		//单调队列为空时,返回-1
		return -1
	}
	//队头为最大值
	return this.mq[0]
}

func (this *MaxQueue) Push_back(value int) {
	//加入一个值之前,需要循环判断单调队列的队尾元素是否小于当前值,是则删掉队尾的元素
	for size := len(this.mq); size > 0 && this.mq[size-1] < value; size = len(this.mq) {
		this.mq = this.mq[:size-1]
	}
	//加入队列
	this.q = append(this.q, value)
	this.mq = append(this.mq, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		//队列为空,返回-1
		return -1
	}
	value := this.q[0]
	if this.mq[0] == value {
		//如果出队的元素和单调队列队头的元素相同时,单调队列也需要出队
		this.mq = this.mq[1:]
	}
	this.q = this.q[1:]
	return value
}
