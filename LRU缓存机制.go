package leetcode

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	Data map[int]*Node
	Head *Node
	Tail *Node
	Cap  int
}

func Constructor(capacity int) LRUCache {
	h := &Node{
		Key: -1,
		Val: -1,
	}
	t := &Node{
		Key: -1,
		Val: -1,
	}
	h.Next = t
	t.Prev = h
	return LRUCache{
		Data: make(map[int]*Node),
		Cap:  capacity,
		Head: h,
		Tail: t,
	}
}

func (l *LRUCache) insert(n *Node) {
	n.Next = l.Head.Next
	n.Next.Prev = n
	l.Head.Next = n
	n.Prev = l.Head
}

func (l *LRUCache) remove(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.Data[key]; !ok {
		return -1
	} else {
		l.remove(node)
		l.insert(node)
		return node.Val
	}
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.Data[key]; !ok {
		// add
		add := &Node{
			Key: key,
			Val: value,
		}
		if len(l.Data) == l.Cap {
			last := l.Tail.Prev
			l.remove(last)
			delete(l.Data, last.Key)
		}
		l.insert(add)
		l.Data[key] = add
	} else {
		// modify
		l.remove(node)
		node.Val = value
		l.insert(node)
		l.Data[key] = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
