package leetcode_1

type MyCircularDeque struct {
	capacity int
	cnt      int
	head     *DeNode
	tail     *DeNode
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		capacity: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &DeNode{Val: value}
	if this.IsEmpty() {
		this.tail = node
	} else {
		node.next = this.head
		this.head.prev = node
	}
	this.head = node
	this.cnt++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &DeNode{Val: value}
	if this.IsEmpty() {
		this.head = node
	} else {
		this.tail.next = node
		node.prev = this.tail
	}
	this.tail = node
	this.cnt++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = this.head.next
	this.cnt--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = this.tail.prev
	this.cnt--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.Val
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.Val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.cnt == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.cnt == this.capacity
}
