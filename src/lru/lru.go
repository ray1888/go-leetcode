package lru

import (
	"errors"
)

type DNode struct {
	val  int
	key  int
	prev *DNode
	next *DNode
}

type LRUCache struct {
	record map[int]*DNode
	head   *DNode
}

func NewLRUCache(cap int) (*LRUCache, error) {
	if cap <= 0 {
		return nil, errors.New("cap must >= 0")
	}
	lru := new(LRUCache)
	lru.record = make(map[int]*DNode)
	head := &DNode{-1, -1, nil, nil}
	cur := head
	for i := 0; i < cap-1; i++ {
		cur.next = &DNode{-1, -1, cur, nil}
		cur = cur.next
	}
	cur.next = head
	head.prev = cur
	lru.head = head
	return lru, nil
}

func (c *LRUCache) move2head(node *DNode) {
	if node == c.head {
		c.head = c.head.prev
		return
	}
	// detach this node
	node.prev.next = node.next
	node.next.prev = node.prev
	// attach this node to head
	node.next = c.head.next
	node.next.prev = node
	c.head.next = node
	node.prev = c.head
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.record[key]; ok {
		val := node.val
		c.move2head(node)
		return val
	} else {
		return -1
	}
}

func (c *LRUCache) Put(key, value int) {
	/*
		three condition needed to think of
		1. put same key
		2. put key not exist , but key length doesn't meet limit
		3. put key not exist , but key length meet limit
	*/
	if node, ok := c.record[key]; ok {
		node.val = value
		c.move2head(node)
	} else {
		if c.head.val != -1 {
			delete(c.record, c.head.val)
		}
		c.head.val = value
		c.head.key = key
		c.record[key] = c.head
		c.head = c.head.prev
	}
}
