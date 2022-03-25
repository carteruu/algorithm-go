package leetcode_1

func climbStairs_3(n int) int {
	pp, p := 0, 1
	for i := 1; i <= n; i++ {
		pp, p = p, pp+p
	}
	return p
}
func climbStairs_2(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs_1(n int) int {
	cache := make([]int, n+1)
	cache[0] = 1
	cache[1] = 1
	return dfs(n, cache)
}
func dfs(n int, cache []int) int {
	if cache[n] == 0 {
		cache[n] = dfs(n-1, cache) + dfs(n-2, cache)
	}
	return cache[n]
}
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
