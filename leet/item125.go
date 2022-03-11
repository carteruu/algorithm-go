package leet

func isPalindrome3(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for ; l < r && skip(s[l]); l++ {

		}
		for ; l < r && skip(s[r]); r-- {

		}
		if l >= r {
			break
		}
		if isDigit(s[l]) && s[l] != s[r] {
			return false
		}
		if isLower(s[l]) && s[l] != s[r] && s[l]-'a' != s[r]-'A' {
			return false
		}
		if isUpper(s[l]) && s[l] != s[r] && s[l]-'A' != s[r]-'a' {
			return false
		}
		l++
		r--
	}
	return true
}

func skip(b byte) bool {
	return !isDigit(b) && !isUpper(b) && !isLower(b)
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
func isUpper(b byte) bool {
	return 'A' <= b && b <= 'Z'
}
func isLower(b byte) bool {
	return 'a' <= b && b <= 'z'
}
