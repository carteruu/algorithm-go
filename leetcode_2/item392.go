package leetcode_2

func isSubsequence1(s string, t string) bool {
	if t == "" {
		return s == ""
	}
	idx := make([][26]int, len(t)+1)
	for i := len(t) - 1; i >= 0; i-- {
		idx[i] = idx[i+1]
		idx[i][t[i]-'a'] = i + 1
	}

	pos := 0
	for i := 0; i < len(s); i++ {
		pos = idx[pos][s[i]-'a']
		if pos == 0 {
			return false
		}
	}
	return true
}
func isSubsequence(s string, t string) bool {
	i := 0
	for j := 0; i < len(s) && j < len(t); i++ {
		for ; j < len(t) && s[i] != t[j]; j++ {

		}
		if j == len(t) {
			return false
		}
	}
	return i == len(s)
}
