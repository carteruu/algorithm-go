package leetcode

import "fmt"

func simplifiedFractions(n int) []string {
	gcd := func(a, b int) int {
		if a < b {
			a, b = b, a
		}
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	var ans []string
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				ans = append(ans, fmt.Sprintf("%v/%v", j, i))
			}
		}
	}
	return ans
}
func simplifiedFractions1(n int) []string {
	var ans []string
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			find := true
			for k := 2; k <= j; k++ {
				if i%k == 0 && j%k == 0 {
					find = false
					break
				}
			}
			if find {
				ans = append(ans, fmt.Sprintf("%v/%v", j, i))
			}
		}
	}
	return ans
}
