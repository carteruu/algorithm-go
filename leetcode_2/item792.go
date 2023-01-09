package leetcode_2

func numMatchingSubseq2(s string, words []string) int {
	ans := 0
	idx := make(map[int][26]int, len(s)+1)
	ii := [26]int{}
	for i := 0; i < 26; i++ {
		ii[i] = -1
	}
	idx[len(s)-1] = ii
	for i := len(s) - 1; i >= 0; i-- {
		ii := idx[i]
		ii[s[i]-'a'] = i
		idx[i-1] = ii
	}
	for _, word := range words {
		pos := -1
		i := 0
		for ; i < len(word); i++ {
			pos = idx[pos][word[i]-'a']
			if pos == -1 {
				break
			}
		}
		if i == len(word) {
			ans++
		}
	}
	return ans
}
func numMatchingSubseq1(s string, words []string) int {
	ans := 0
	idx := make([][26]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		idx[i] = idx[i+1]
		idx[i][s[i]-'a'] = i + 1
	}
	for _, word := range words {
		pos := 0
		i := 0
		for ; i < len(word); i++ {
			pos = idx[pos][word[i]-'a']
			if pos == 0 {
				break
			}
		}
		if i == len(word) {
			ans++
		}
	}
	return ans
}
func numMatchingSubseq(s string, words []string) int {
	ans := 0
	for _, word := range words {
		idx := 0
		i := 0
		for i < len(word) && idx < len(s) {
			for ; idx < len(s); idx++ {
				if s[idx] == word[i] {
					i++
					idx++
					break
				}
			}
		}
		if i == len(word) {
			ans++
		}
	}
	return ans
}
