package main

type Trie struct {
	Children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, bt := range word {
		bt -= 'a'
		if node.Children[bt] == nil {
			node.Children[bt] = &Trie{}
		}
		node = node.Children[bt]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, bt := range word {
		bt -= 'a'
		if node.Children[bt] == nil {
			return false
		}
		node = node.Children[bt]
	}
	if node.isEnd {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, bt := range prefix {
		bt -= 'a'
		if node.Children[bt] == nil {
			return false
		}
		node = node.Children[bt]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
