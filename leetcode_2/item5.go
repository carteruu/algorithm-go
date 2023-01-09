package leetcode_2

func longestPalindrome(s string) string {
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		for j := 1; i-j >= 0 && i+j < len(s) && s[i-j] == s[i+j]; j++ {
			if 2*j > r-l {
				l, r = i-j, i+j
			}
		}
		for j := 0; i-j >= 0 && i+j+1 < len(s) && s[i-j] == s[i+j+1]; j++ {
			if 2*j+1 > r-l {
				l, r = i-j+1, i+j
			}
		}
	}
	return s[l:r]
}
