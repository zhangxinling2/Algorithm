package main

type Node struct {
	key, value int
	pre, next  *Node
}

func createNode(k, v int) *Node {
	return &Node{
		key:   k,
		value: v,
	}
}

type LinkedList struct {
	head, tail *Node
	size       int
}

func NewLinkdedList(size int) *LinkedList {
	dll := &LinkedList{
		head: createNode(0, 0),
		tail: createNode(0, 0),
		size: 0,
	}
	dll.head.next = dll.tail
	dll.tail.pre = dll.head
	return dll
}
func (l *LinkedList) addLast(n *Node) {
	n.pre = l.tail.pre
	n.next = l.tail
	l.tail.pre.next = n
	l.tail.pre = n
	l.size++
}
func (l *LinkedList) remove(n *Node) {
	n.pre.next = n.next
	n.next.pre = n.pre
	n.pre = nil
	n.next = nil
	l.size--
}
func (l *LinkedList) removeFirst() *Node {
	if l.head.next == l.tail {
		return nil
	}
	node := l.head.next
	l.remove(l.head.next)
	return node
}

type LRUCache struct {
	capacity int
	list     *LinkedList
	index    map[int]*Node
}

func LRUConstructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     NewLinkdedList(capacity),
		index:    map[int]*Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.index[key]; ok {
		this.list.remove(val)
		this.list.addLast(val)
		return val.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.index[key]; ok {
		val.value = value
		this.list.remove(val)
		this.list.addLast(val)
		return
	}
	node := createNode(key, value)
	if this.list.size >= this.capacity {
		tmp := this.list.removeFirst()
		delete(this.index, tmp.key)

	}
	this.list.addLast(node)
	this.index[key] = node
}
