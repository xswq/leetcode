package leetcode

type Trie struct {
	isEnd    bool
	children []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		isEnd:    false,
		children: make([]*Trie, 26),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, s := range word {
		index := s - 'a'
		if node.children[index] == nil {
			temp := Constructor()
			node.children[index] = &temp
		}
		node = node.children[index]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, s := range word {
		index := s - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, s := range prefix {
		index := s - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true
}
