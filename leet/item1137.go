package leet

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	dp1, dp2, dp3 := 0, 1, 1
	for i := 3; i <= n; i++ {
		dp := dp1 + dp2 + dp3
		dp1, dp2, dp3 = dp2, dp3, dp
	}
	return dp3
}
