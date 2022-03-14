package leetcode

func titleToNumber(columnTitle string) int {
	res := 0
	for i := 0; i < len(columnTitle); i++ {
		res *= 26
		res += int(columnTitle[i]-'A') + 1
	}
	return res
}
