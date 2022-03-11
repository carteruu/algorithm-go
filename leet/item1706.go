package leet

func findBall(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		col := i
		for row := 0; row < m; row++ {
			//当前挡板的方向
			dir := grid[row][col]
			//按挡板方向移动小球后，列的位置（左移或右移）
			col += dir
			//如果不能移动（超过列的范围），或新位置的挡板方向和原位置挡板方向不同，则不能到达底部
			if col < 0 || col >= n || grid[row][col] != dir {
				col = -1
				break
			}
		}
		if 0 <= col && col < n {
			ans[i] = col
		} else {
			ans[i] = -1
		}
	}
	return ans
}
func findBall1(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	ans := make([]int, n)
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row == m {
			if 0 <= col && col < n {
				return col
			}
			return -1
		}
		if col < 0 || col >= n ||
			(col < n-1 && grid[row][col] == 1 && grid[row][col+1] == -1) ||
			(col > 0 && grid[row][col] == -1 && grid[row][col-1] == 1) {
			return -1
		}
		return dfs(row+1, col+grid[row][col])
	}
	for i := 0; i < n; i++ {
		ans[i] = dfs(0, i)
	}
	return ans
}
