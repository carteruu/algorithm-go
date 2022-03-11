package leet

import "math"

func numSquares(n int) int {
	l := int(math.Pow(float64(n), 0.5))
	if l*l == n {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for j := 1; j <= n; j++ {
		dp[j] = math.MaxInt64
	}
	for i := 1; i <= l; i++ {
		num := i * i
		for j := num; j <= n; j++ {
			if dp[j-num]+1 < dp[j] {
				dp[j] = dp[j-num] + 1
			}
		}
	}
	return dp[n]
}
func numSquares2(n int) int {
	l := int(math.Pow(float64(n), 0.5))
	if l*l == n {
		return 1
	}
	arr := make([]int, l)
	for i := 1; i < len(arr); i++ {
		arr[i] = i * i
	}

	dp := make([]int, n+1)
	for j := 1; j <= n; j++ {
		dp[j] = math.MaxInt64
	}
	dp[0] = 0
	for i := 1; i <= l; i++ {
		temp := make([]int, n+1)
		for j := 0; j <= n; j++ {
			temp[j] = dp[j]
			if j >= arr[i] && temp[j-arr[i]] < temp[j] {
				temp[j] = temp[j-arr[i]] + 1
			}
		}
		dp = temp
	}
	return dp[n]
}
func numSquares3(n int) int {
	l := int(math.Pow(float64(n), 0.5))
	if l*l == n {
		return 1
	}
	arr := make([]int, l)
	for i := 1; i < len(arr); i++ {
		arr[i] = i * i
	}

	dp := make([][]int, l)
	for i := 0; i <= l; i++ {
		dp[i] = make([]int, n+1)
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = math.MaxInt64
	}
	dp[0][0] = 0
	for i := 1; i < l; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= arr[i] && dp[i][j-arr[i]] < dp[i][j] {
				dp[i][j] = dp[i][j-arr[i]] + 1
			}
		}
	}
	return dp[l-1][n]
}
