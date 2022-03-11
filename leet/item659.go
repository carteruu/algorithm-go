package leet

func maxAreaOfIsland11(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	neighbor := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		area := 1
		for _, nb := range neighbor {
			area += dfs(i+nb[0], j+nb[1])
		}
		return area
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := dfs(i, j)
				if area > ans {
					ans = area
				}
			}
		}
	}
	return ans
}
