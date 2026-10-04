package mycache

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	cap int
	m   map[string]*list.Element
	l   *list.List
	mu  sync.Mutex
}

type entry struct {
	key   string
	value any
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		cap: cap,
		m:   make(map[string]*list.Element),
		l:   list.New(),
	}
}

// Get 获取，命中则移到头部
func (c *LRUCache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ele, ok := c.m[key]
	if !ok {
		return nil, false
	}
	c.l.MoveToFront(ele)
	return ele.Value.(*entry).value, true
}

// Put 写入，存在就更新并移到头部；不存在新增，满了淘汰尾部
func (c *LRUCache) Put(key string, val any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 已存在，更新value，移到头部
	if ele, ok := c.m[key]; ok {
		ele.Value.(*entry).value = val
		c.l.MoveToFront(ele)
		return
	}
	// 新增节点
	newEle := c.l.PushFront(&entry{key: key, value: val})
	c.m[key] = newEle

	// 超过容量，淘汰尾部
	if c.l.Len() > c.cap {
		delEle := c.l.Back()
		if delEle != nil {
			c.l.Remove(delEle)
			delKey := delEle.Value.(*entry).key
			delete(c.m, delKey)
		}
	}
}
