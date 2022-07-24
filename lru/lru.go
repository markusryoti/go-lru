package lru

type LRU struct {
	capacity  int
	length    int
	list      *doublyLinkedList
	searchMap map[string]*node
}

type doublyLinkedList struct {
	first *node
	last  *node
}

type node struct {
	value    int
	next     *node
	previous *node
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		list:      &doublyLinkedList{},
		searchMap: make(map[string]*node, capacity),
		capacity:  capacity,
		length:    0,
	}
}

func (l *LRU) Get(key string) int {
	n := l.searchMap[key]
	if n == nil {
		return -1
	}

	currentFirst := l.list.first
	l.updateNode(n, currentFirst, n.value)
	return n.value
}

func (l *LRU) Put(key string, value int) {
	existingNode := l.searchMap[key]
	currentFirst := l.list.first

	if existingNode != nil {
		l.updateNode(existingNode, currentFirst, value)
		return
	}

	if l.length == l.capacity {
		l.removeLast()
	}

	l.addNewNode(key, value, currentFirst)
}

func (l *LRU) updateNode(existing, currentFirst *node, value int) {
	existing.value = value
	currentNext := existing.next

	if currentNext != nil {
		currentNext.previous = existing.previous
	}

	existing.previous = nil
	existing.next = currentFirst
	currentFirst.previous = existing
	l.list.first = existing
}

func (l *LRU) removeLast() {
	currentLast := l.list.last
	if currentLast.previous != nil {
		currentLast.previous.next = nil
	}
	l.list.last = currentLast.previous
	l.length--
}

func (l *LRU) addNewNode(key string, value int, currentFirst *node) {
	l.searchMap[key] = &node{
		value: value,
	}

	n := l.searchMap[key]

	if currentFirst == nil {
		l.list.first = n
		l.list.last = n
		l.length++
	} else {
		n.next = currentFirst
		currentFirst.previous = n
		l.list.first = n
		l.length++
	}
}
