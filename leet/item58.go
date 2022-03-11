package leet

func lengthOfLastWord(s string) int {
	idx := len(s) - 1
	ans := 0
	for ; idx >= 0 && s[idx] == ' '; idx-- {
	}
	for ; idx >= 0 && s[idx] != ' '; idx-- {
		ans++
	}
	return ans
}
func lengthOfLastWord1(s string) int {
	l := 0
	restart := false
	for _, c := range s {
		if c == ' ' {
			restart = true
			continue
		}
		if restart {
			l = 0
			restart = false
		}
		l++
	}
	return l
}
