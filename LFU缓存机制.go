package leetcode

type Node struct {
	Key   int
	Value int
	Feq   int
	Prev  *Node
	Next  *Node
	List  *LinkList
}

type LinkList struct {
	Fre  int
	Prev *LinkList
	Next *LinkList
	Head *Node
	Tail *Node
}

type LFUCache struct {
	Data    map[int]*Node
	MaxList *LinkList
	MinList *LinkList
	Size    int
	Cap     int
}

func NewLinkList() *LinkList {
	h := &Node{}
	t := &Node{}
	h.Next = t
	t.Prev = h
	return &LinkList{
		Head: h,
		Tail: t,
	}
}

func Constructor(cap int) LFUCache {
	max := NewLinkList()
	min := NewLinkList()
	max.Next = min
	min.Prev = max
	return LFUCache{
		Data:    make(map[int]*Node),
		MaxList: max,
		MinList: min,
		Size:    0,
		Cap:     cap,
	}
}

func (l *LinkList) addNode(n *Node) {
	n.Next = l.Head.Next
	l.Head.Next.Prev = n
	l.Head.Next = n
	n.Prev = l.Head
	n.List = l
}

func (l *LinkList) removeNode(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

func (l *LFUCache) Get(key int) int {
	if node, ok := l.Data[key]; !ok {
		return -1
	} else {
		l.feqInc(node)
		return node.Value
	}
}

func (l *LFUCache) Put(key, value int) {
	if l.Cap == 0 {
		return
	}
	if node, ok := l.Data[key]; ok {
		node.Value = value
		l.feqInc(node)
	} else {
		if l.Size == l.Cap {
			delNode := l.MinList.Prev.Tail.Prev
			l.MinList.Prev.removeNode(delNode)
			delete(l.Data, delNode.Key)
			l.Size--
			if l.MinList.Prev.Head.Next == l.MinList.Prev.Tail {
				l.MinList.Prev.removeList()
			}
		}
		node := &Node{
			Key:   key,
			Value: value,
			Feq:   1,
		}
		l.Data[key] = node
		if l.MinList.Prev.Fre != 1 {
			newList := NewLinkList()
			newList.Fre = 1
			l.MinList.Prev.addList(newList)
			newList.addNode(node)
		} else {
			l.MinList.Prev.addNode(node)
		}
		l.Size++
	}
}

func (l *LFUCache) feqInc(n *Node) {
	preList := n.List.Prev
	n.List.removeNode(n)
	if n.List.Head.Next == n.List.Tail {
		n.List.removeList()
	}
	n.Feq++
	if preList.Fre != n.Feq {
		newList := NewLinkList()
		newList.Fre = n.Feq
		preList.addList(newList)
		newList.addNode(n)
	} else {
		preList.addNode(n)
	}
}

func (l *LinkList) removeList() {
	l.Prev.Next = l.Next
	l.Next.Prev = l.Prev
}

func (l *LinkList) addList(list *LinkList) {
	list.Next = l.Next
	l.Next.Prev = list
	l.Next = list
	list.Prev = l
}
