package leetcode

import "sort"

func longestWord(words []string) string {
	root := &Trie{isEnd: true}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) < len(words[j])
	})
	var ans string
	for _, word := range words {
		node := root
		is := true
		for _, c := range word {
			if !node.isEnd {
				is = false
				break
			}
			idx := c - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &Trie{}
			}
			node = node.children[idx]
		}
		if is {
			node.isEnd = true
			if len(word) > len(ans) {
				ans = word
			}
		}
	}
	return ans
}
