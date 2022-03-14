package leetcode

func numIslands1(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	neighbor := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		for _, nb := range neighbor {
			dfs(i+nb[0], j+nb[1])
		}
		return
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
