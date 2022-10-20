package leetcode

type listNode struct {
	key  int
	val  int
	prev *listNode
	next *listNode
}

type LRUCache struct {
	size     int
	cacheMap map[int]*listNode
	head     *listNode
	tail     *listNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		size:     capacity,
		cacheMap: make(map[int]*listNode),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cacheMap[key]; ok {
		if node == this.tail {
			return node.val
		}

		if node == this.head {
			this.head = this.head.next
		}
		// move to tail
		if node.prev != nil {
			node.prev.next = node.next
		}
		if node.next != nil {
			node.next.prev = node.prev
		}
		node.prev = this.tail
		this.tail.next = node
		this.tail = node
		this.tail.next = nil

		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cacheMap[key]; ok {
		node.val = value
		this.Get(key)
	} else {
		node := &listNode{key: key, val: value}

		if this.head == nil {
			// head
			this.head = node
			this.tail = node
		} else {
			// add to tail
			node.prev = this.tail
			this.tail.next = node
			this.tail = node
		}

		this.cacheMap[key] = node

		// remove head if over capacity
		if len(this.cacheMap) > this.size {
			delete(this.cacheMap, this.head.key)

			this.head = this.head.next
			this.head.prev = nil
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
