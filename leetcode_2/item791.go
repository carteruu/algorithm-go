package leetcode_2

import "sort"

func customSortString1(order string, s string) string {
	charExist := [26]bool{}
	for _, c := range order {
		charExist[c-'a'] = true
	}
	ans := make([]byte, len(s))
	idx := 0
	cs := [26]int{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !charExist[c-'a'] {
			ans[idx] = c
			idx++
			continue
		}
		cs[c-'a']++
	}
	for i := 0; i < len(order); i++ {
		c := order[i]
		for j := 0; j < cs[c-'a']; j++ {
			ans[idx] = c
			idx++
		}
	}
	return string(ans)
}
func customSortString(order string, s string) string {
	idx := [26]int{}
	for i, c := range order {
		idx[c-'a'] = i
	}
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return idx[bytes[i]-'a'] < idx[bytes[j]-'a']
	})
	return string(bytes)
}
