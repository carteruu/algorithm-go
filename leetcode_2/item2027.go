package leetcode_2

func minimumMoves(s string) int {
	ans := 0
	for i := 0; i < len(s); {
		if s[i] == 'O' {
			i++
			continue
		}
		ans++
		i += 3
	}
	return ans
}
