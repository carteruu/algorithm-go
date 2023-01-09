package leetcode_2

import "sort"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	start := strs[0]
	end := strs[len(strs)-1]
	for i := 0; i < len(start); i++ {
		if len(end) <= i || end[i] != start[i] {
			return start[:i]
		}
	}
	return start
}
