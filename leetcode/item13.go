package leetcode

func romanToInt(s string) int {
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := 0
	for i, c := range s {
		if i+1 < len(s) {
			next := s[i+1]
			if (c == 'I' && (next == 'V' || next == 'X')) ||
				(c == 'X' && (next == 'L' || next == 'C')) ||
				(c == 'C' && (next == 'D' || next == 'M')) {
				ans -= m[c]
				continue
			}
		}
		ans += m[c]
	}
	return ans
}
