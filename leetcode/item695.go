package leetcode

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	seen := make([][]bool, m)
	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		if seen[i][j] || grid[i][j] == 0 {
			return 0
		}
		seen[i][j] = true
		return 1 + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !seen[i][j] {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}
