package item641_1

type MyCircularDeque struct {
	elements []int
	capacity int
	head     int
	tail     int
	cnt      int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		capacity: k,
		elements: make([]int, k, k),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head = (this.head - 1 + this.capacity) % this.capacity
	this.elements[this.head] = value
	this.cnt++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.elements[this.tail] = value
	this.tail = (this.tail + 1) % this.capacity
	this.cnt++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.capacity
	this.cnt--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail - 1 + this.capacity) % this.capacity
	this.cnt--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.elements[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.elements[(this.tail-1+this.capacity)%this.capacity]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.cnt == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.cnt == this.capacity
}
