package leet

import "math"

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
func minPathSum1(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	cache := make([][]int, m)
	seen := make([][]bool, m)
	cache[m-1][n-1] = grid[m-1][n-1]
	seen[m-1][n-1] = true
	for i := range grid {
		cache[i] = make([]int, n)
		seen[i] = make([]bool, n)
	}
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row >= m || col >= n {
			return math.MaxInt64
		}
		if seen[row][col] {
			return cache[row][col]
		}
		val := min(dfs(row+1, col), dfs(row, col+1)) + grid[row][col]
		cache[row][col] = val
		seen[row][col] = true
		return val
	}
	return dfs(0, 0)
}
