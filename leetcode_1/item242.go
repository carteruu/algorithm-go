package leetcode_1

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if cnt[i] != 0 {
			return false
		}
	}
	return true
}
