package leetcode

func longestPalindrome(s string) string {
	isPalindrome := func(start, end int) (int, int) {
		for start >= 0 && end < len(s) && s[start] == s[end] {
			start--
			end++
		}
		//前闭后开
		return start + 1, end
	}
	l, r := 0, 1
	for i := 0; i < len(s); i++ {
		nl, nr := isPalindrome(i, i)
		if nr-nl > r-l {
			l, r = nl, nr
		}
		nl, nr = isPalindrome(i, i+1)
		if nr-nl > r-l {
			l, r = nl, nr
		}
	}
	return s[l:r]
}
