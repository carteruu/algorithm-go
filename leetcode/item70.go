package leetcode

func climbStairs(n int) int {
	m := make([]int, n+1)
	var dfs func(n int) int
	dfs = func(n int) int {
		if n <= 2 {
			return n
		}
		if m[n] != 0 {
			return m[n]
		}
		m[n] = dfs(n-1) + dfs(n-2)
		return m[n]
	}
	return dfs(n)
}
