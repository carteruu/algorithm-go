package leet

//节点
type node struct {
	key, val, times int
	//双向链表,删除节点时,时间复杂度O(1)
	prev, next *node
}
type LFUCache struct {
	key2Node map[int]*node
	//每个次数链表的虚拟头尾节点
	times2Endpoint map[int][2]*node
	capacity       int
	minTimes       int
}

func ConstructorLFUCache(capacity int) LFUCache {
	return LFUCache{
		capacity:       capacity,
		key2Node:       make(map[int]*node, capacity),
		times2Endpoint: make(map[int][2]*node, capacity),
	}
}

func (this *LFUCache) Get(key int) int {
	if n, ok := this.key2Node[key]; !ok {
		//不存在
		return -1
	} else {
		//存在,
		this.removeTimesLink(n)
		//判断访问次数是否是最少访问次数,且移动后,次数链表为空(头尾相连)时,需要将最少访问次数+1
		if endpoint := this.times2Endpoint[n.times]; n.times == this.minTimes && endpoint[0].next == endpoint[1] {
			this.minTimes++
		}
		this.addOneTimes(n)
		return n.val
	}
}

func (this *LFUCache) Put(key int, value int) {
	if n, ok := this.key2Node[key]; ok {
		n.val = value
		this.removeTimesLink(n)
		//判断访问次数是否是最少访问次数,且移动后,次数链表为空(头尾相连)时,需要将最少访问次数+1
		if endpoint := this.times2Endpoint[n.times]; n.times == this.minTimes && endpoint[0].next == endpoint[1] {
			this.minTimes++
		}
		this.addOneTimes(n)
	} else {
		if this.capacity == 0 {
			//容量为0,不添加,直接返回
			return
		}
		if len(this.key2Node) >= this.capacity {
			//容量已满,删除最少使用次数的尾部节点
			minEndPoint := this.times2Endpoint[this.minTimes]
			tail := minEndPoint[1]
			delNode := tail.prev
			delete(this.key2Node, delNode.key)
			tail.prev = tail.prev.prev
			tail.prev.next = tail
		}
		//添加
		n = &node{
			key:   key,
			val:   value,
			times: 0,
		}
		this.key2Node[key] = n
		this.addOneTimes(n)
		//设置最少访问次数为1
		this.minTimes = 1
	}
}

//增加一次访问次数
func (this *LFUCache) addOneTimes(n *node) {
	n.times++
	endpoint, ok := this.times2Endpoint[n.times]
	if !ok {
		this.initEndpoint(n.times)
		endpoint = this.times2Endpoint[n.times]
	}
	head := endpoint[0]
	//移到头部
	n.next = head.next
	n.prev = head
	head.next.prev = n
	head.next = n
}

//从访问次数的链表中删除
func (this *LFUCache) removeTimesLink(n *node) {
	//从原来的次数链表删除
	n.prev.next = n.next
	n.next.prev = n.prev
}

//初始化次数链表的虚拟头尾节点
func (this *LFUCache) initEndpoint(times int) {
	//新的链表不存在,先初始化
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head
	this.times2Endpoint[times] = [2]*node{head, tail}
}
