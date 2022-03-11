package leet

type LRUCache1 struct {
	capacity int
	head     *node
	tail     *node
	cache    map[int]*node
}

func ConstructorLRUCache(capacity int) LRUCache1 {
	if capacity <= 0 {
		panic("容量必须大于0")
	}
	LRUCache1 := LRUCache1{
		capacity: capacity,
		head:     &node{},
		tail:     &node{},
		cache:    make(map[int]*node, capacity),
	}
	LRUCache1.head.next = LRUCache1.tail
	LRUCache1.tail.prev = LRUCache1.head
	return LRUCache1
}

func (this *LRUCache1) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache1) Put(key int, value int) {
	if n, exists := this.cache[key]; exists {
		n.val = value
		this.moveToHead(n)
	} else {
		if len(this.cache) == this.capacity {
			//已满,删除尾部
			this.delete(this.tail.prev)
		}
		n := &node{key: key, val: value}
		this.cache[key] = n
		this.addToHead(n)
	}
}

func (this *LRUCache1) moveToHead(node *node) {
	this.delete(node)
	this.addToHead(node)
}

func (this *LRUCache1) addToHead(node *node) {
	//添加到队头
	node.next = this.head.next
	node.prev = this.head

	this.head.next.prev = node
	this.head.next = node
	this.cache[node.key] = node
}

func (this *LRUCache1) delete(node *node) {
	//删除
	node.prev.next = node.next
	node.next.prev = node.prev
	delete(this.cache, node.key)
}
