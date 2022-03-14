package leetcode

import "strings"

func removeKdigits1(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	for i := 0; i < k; i++ {
		var flag bool
		for j := 0; j < len(num)-1; j++ {
			if num[j] != '0' && ((j == 0 && num[j+1] == '0') || (num[j] >= num[j+1])) {
				num = num[:j] + num[j+1:]
				flag = true
				break
			}
		}
		if !flag {
			num = num[:len(num)-1]
		}
	}
	idx := 0
	for ; idx < len(num); idx++ {
		if num[idx] != '0' {
			break
		}
	}
	if idx == len(num) {
		return "0"
	}
	return num[idx:]
}

func removeKdigits11(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
