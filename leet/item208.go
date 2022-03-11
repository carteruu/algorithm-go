package leet

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

func ConstructorTrie() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}
	node := this
	for _, c := range word {
		c -= 'a'
		if node.children[c] == nil {
			node.children[c] = &Trie{}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.searchStart(word)
	return node != nil && node.isEnd
}

func (this *Trie) searchStart(word string) *Trie {
	node := this
	for _, c := range word {
		c -= 'a'
		if node.children[c] == nil {
			return nil
		}
		node = node.children[c]
	}
	return node
}

func (this *Trie) StartsWith(prefix string) bool {
	this.searchStart(prefix)
	return true
}
