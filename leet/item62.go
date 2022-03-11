package leet

import "math/big"

func uniquePaths(m int, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
func uniquePaths1(m int, n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 1; i <= m; i++ {
		temp := make([]int, n+1)
		for j := 1; j <= n; j++ {
			temp[j] = dp[j] + temp[j-1]
		}
		dp = temp
	}
	return dp[n]
}
func uniquePaths2(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}
