package linklist

import "fmt"

type doubleListNode struct {
	key  int
	val  int
	pre  *doubleListNode
	next *doubleListNode
}

type doubleList struct {
	len  int
	cap  int
	next *doubleListNode
	tail *doubleListNode
}

func (this *doubleList) moveToHead(node *doubleListNode) {
	if node == nil || node.pre == nil {
		return
	}
	// get node out of it
	pre := node.pre
	next := node.next
	if next != nil {
		next.pre = pre
	}
	pre.next = next
	// self
	node.pre = nil
	node.next = this.next
	// insert to head
	this.next.pre = node
	this.next = node
	// change tail, 自己写的代码就忘了，WA
	if node == this.tail {
		this.tail = pre
	}
}

func (this *doubleList) insertNodeToHead(key, value int) *doubleListNode {
	node := &doubleListNode{
		key: key,
		val: value,
	}
	node.pre = nil
	node.next = this.next
	if this.next == nil {
		this.next = node
		this.tail = node
		this.len = 1
		return node
	}
	this.next.pre = node
	this.next = node
	this.len++
	return node
}

func (this *doubleList) deleteTail() *doubleListNode {
	tail := this.tail
	if this.len == 1 {
		this.next, this.tail = nil, nil
		this.len = 0
		return tail
	}
	fmt.Println(this.len)
	tail.pre.next = nil
	this.tail = tail.pre
	this.len--
	return tail
}

type LRUCache struct {
	mapKV    map[int]*doubleListNode
	listHead doubleList
}

func Constructor(capacity int) LRUCache {
	var cache LRUCache
	cache.mapKV = make(map[int]*doubleListNode)
	cache.listHead = doubleList{
		len:  0,
		cap:  capacity,
		next: nil,
		tail: nil,
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.mapKV[key]; ok {
		this.listHead.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.mapKV[key]; ok {
		this.listHead.moveToHead(node)
		node.val = value
		return
	}
	var node *doubleListNode
	if this.listHead.cap > this.listHead.len {
		node = this.listHead.insertNodeToHead(key, value)
	} else {
		node = this.listHead.deleteTail()
		delete(this.mapKV, node.key)
		node = this.listHead.insertNodeToHead(key, value)
	}
	fmt.Printf("after put, len %d\n", this.listHead.len)
	this.mapKV[key] = node
}
