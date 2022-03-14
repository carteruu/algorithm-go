package leetcode

import (
	"container/list"
	"runtime/debug"
)

// github.com/EndlessCheng/codeforces-go
func init() { debug.SetGCPercent(-1) }

type lruEntry struct {
	key, value int
}

type LRUCache2 struct {
	capacity int
	cache    map[int]*list.Element
	lst      *list.List
}

func ConstructorLRUCache2(capacity int) LRUCache2 {
	return LRUCache2{capacity, map[int]*list.Element{}, list.New()}
}

func (c *LRUCache2) Get(key int) int {
	e := c.cache[key]
	if e == nil {
		return -1
	}
	c.lst.MoveToFront(e) // 刷新缓存
	return e.Value.(lruEntry).value
}

func (c *LRUCache2) Put(key, value int) {
	if e := c.cache[key]; e != nil {
		e.Value = lruEntry{key, value}
		c.lst.MoveToFront(e) // 刷新缓存
		return
	}
	c.cache[key] = c.lst.PushFront(lruEntry{key, value})
	if len(c.cache) > c.capacity {
		delete(c.cache, c.lst.Remove(c.lst.Back()).(lruEntry).key)
	}
}
