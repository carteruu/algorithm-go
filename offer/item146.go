package offer

type LRUCache struct {
	capacity int
	count    int
	cache    map[int]*node
	head     *node
	tail     *node
}

type node struct {
	key  int
	val  int
	pre  *node
	next *node
}

func Constructor(capacity int) LRUCache {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		count:    0,
		cache:    make(map[int]*node, 0),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.cache[key]; ok {
		this.moveToHead(n)
		return n.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.cache[key]; ok {
		//存在
		n.val = value
		this.moveToHead(n)
		return
	}
	//不存在
	this.count++
	if this.count > this.capacity {
		//删除尾
		this.delete(this.tail.pre)
	}
	n := &node{
		key: key,
		val: value,
	}
	this.cache[key] = n
	n.next = this.head.next
	this.head.next.pre = n
	this.head.next = n
	n.pre = this.head
}

func (this *LRUCache) moveToHead(n *node) {
	//从原来的位置删除
	n.pre.next = n.next
	n.next.pre = n.pre
	//加入头
	oldNext := this.head.next
	this.head.next = n
	n.pre = this.head
	n.next = oldNext
	oldNext.pre = n
}

func (this *LRUCache) delete(n *node) {
	//从原来的位置删除
	n.pre.next = n.next
	n.next.pre = n.pre
	delete(this.cache, n.key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
