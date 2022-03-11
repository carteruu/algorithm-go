package leet

import "strings"

func removeKdigits(num string, k int) string {
	n := len(num)
	if k >= n {
		return "0"
	}
	stack := make([]byte, 0, n)
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	ans := string(stack[:len(stack)-k])
	ans = strings.TrimLeft(ans, "0")
	if ans == "" {
		return "0"
	}
	return ans
}
