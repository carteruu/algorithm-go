package leetcode

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	dp1, dp2 := 1, 1
	for i := 1; i < len(s); i++ {
		dp3 := 0
		if s[i] != '0' {
			dp3 = dp2
		}
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '0' && s[i] <= '6') {
			dp3 += dp1
		}
		dp1 = dp2
		dp2 = dp3
	}
	return dp2
}
func numDecodings2(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] >= '0' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
func numDecodings1(s string) int {
	n := len(s)
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if idx == n {
			return 1
		}
		if idx > n || s[idx] == '0' {
			return 0
		}
		c := s[idx]
		sum := dfs(idx + 1)
		if idx < n-1 && (c >= '1' || (c <= '2' && s[idx+1] <= '6')) {
			sum += dfs(idx + 2)
		}
		return sum
	}
	return dfs(0)
}
