package leet

func longestCommonPrefix(strs []string) string {
	str0 := strs[0]
	for i := 0; i < len(str0); i++ {
		c := str0[i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				return str0[:i]
			}
		}
	}
	return str0
}
