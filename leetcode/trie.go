package leetcode

type Trie struct {
	isEnd    bool
	children [26]*Trie
}
