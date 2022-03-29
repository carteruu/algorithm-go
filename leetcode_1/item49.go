package leetcode_1

import "sort"

//方法二，按字母计算每个单词的hash值，按hash分组
func groupAnagrams(strs []string) (ans [][]string) {
	m := make(map[string][]string)
	for _, str := range strs {
		//因为每个单词的长度不超过100，所以这里可以用byte
		bucket := make([]byte, 26)
		for _, c := range str {
			bucket[c-'a']++
		}
		key := string(bucket)
		m[key] = append(m[key], str)
	}
	for _, ss := range m {
		ans = append(ans, ss)
	}
	return
}
func groupAnagrams_2(strs []string) (ans [][]string) {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		var bucket [26]int
		for _, c := range str {
			bucket[c-'a']++
		}
		m[bucket] = append(m[bucket], str)
	}
	for _, ss := range m {
		ans = append(ans, ss)
	}
	return
}

//方法一，每个单词都排序，按排序后的单词分组
func groupAnagrams_1(strs []string) (ans [][]string) {
	s := make([]string, len(strs))
	for i, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		s[i] = string(bytes)
	}
	m := make(map[string][]string)
	for i, str := range s {
		m[str] = append(m[str], strs[i])
	}
	for _, ss := range m {
		ans = append(ans, ss)
	}
	return
}
