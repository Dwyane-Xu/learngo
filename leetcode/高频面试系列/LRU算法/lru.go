package main

import "fmt"

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	cap    int
	header *Node
	tail   *Node
	lruMap map[int]*Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:    capacity,
		header: &Node{},
		tail:   &Node{},
		lruMap: make(map[int]*Node, capacity),
	}
	cache.header.next = cache.tail
	cache.tail.prev = cache.header
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.lruMap[key]; ok {
		this.removeNode(node)
		this.addNode(node)
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.lruMap[key]; ok {
		node.val = value
		this.removeNode(node)
		this.addNode(node)
	} else {
		if len(this.lruMap) >= this.cap {
			delete(this.lruMap, this.tail.prev.key)
			this.removeNode(this.tail.prev)
		}
		newNode := &Node{key: key, val: value}
		this.addNode(newNode)
		this.lruMap[key] = newNode
	}
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addNode(node *Node) {
	node.prev = this.header
	node.next = this.header.next
	// this.header.next.prev = node
	// this.header.next = node
	node.prev.next = node
	node.next.prev = node
}

func main() {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	fmt.Println(lruCache.Get(1))
}
