package leet

func numIslands(grid [][]byte) int {
	var sign func(k, l int)
	n := len(grid)
	bytes := grid[0]
	m := len(bytes)
	sign = func(k, l int) {
		if k < 0 || k >= n || l < 0 || l >= m || grid[k][l] != '1' {
			return
		}
		grid[k][l] = '0'
		sign(k-1, l)
		sign(k+1, l)
		sign(k, l-1)
		sign(k, l+1)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				sign(i, j)
				res++
			}
		}
	}
	return res
}
