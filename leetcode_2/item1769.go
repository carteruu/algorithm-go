package leetcode_2

func minOperations1769_1(boxes string) []int {
	ans := make([]int, len(boxes)+1)
	cntL := 0
	for i := 0; i < len(boxes); i++ {
		ans[i+1] += ans[i] + cntL
		cntL += int(boxes[i] - '0')
	}
	cntR := 0
	pre := 0
	for i := len(boxes) - 1; i >= 0; i-- {
		pre += cntR
		cntR += int(boxes[i] - '0')

		ans[i+1] += pre
	}
	return ans[1:]
}
func minOperations1769(boxes string) (ans []int) {
	for i := 0; i < len(boxes); i++ {
		op := 0
		for j, c := range boxes {
			if c == '1' {
				op += absInt(i - j)
			}
		}
		ans = append(ans, op)
	}
	return ans
}
