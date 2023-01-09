package leetcode_2

func minOperations1(s string) int {
	a := 0
	for i := 0; i < len(s); i++ {
		if (i & 1) != int(s[i]-'0') {
			a++
		}
	}
	if a <= len(s)-a {
		return a
	}
	return len(s) - a
}
func minOperations0(s string) int {
	a, b := 0, 0
	for i := 0; i < len(s); i++ {
		if (i&1 == 0 && s[i] == '1') || (i&1 == 1 && s[i] == '0') {
			a++
		}
		if (i&1 == 0 && s[i] == '0') || (i&1 == 1 && s[i] == '1') {
			b++
		}
	}
	if a < b {
		return a
	}
	return b
}
