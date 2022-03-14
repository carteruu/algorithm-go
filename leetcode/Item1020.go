package leetcode

import "algorithm/pkg"

//BFS
func numEnclaves3(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	seen := make([][]bool, m)
	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
	}
	var q [][]int
	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			q = append(q, []int{i, 0})
		}
		if grid[i][n-1] == 1 {
			q = append(q, []int{i, n - 1})
		}
	}
	for i := 0; i < n; i++ {
		if grid[0][i] == 1 {
			q = append(q, []int{0, i})
		}
		if grid[m-1][i] == 1 {
			q = append(q, []int{m - 1, i})
		}
	}
	for len(q) > 0 {
		last := len(q) - 1
		poll := q[last]
		q = q[:last]
		seen[poll[0]][poll[1]] = true
		for _, step := range steps {
			row := poll[0] + step[0]
			col := poll[1] + step[1]
			if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] == 0 || seen[row][col] {
				continue
			}
			q = append(q, []int{row, col})
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !seen[i][j] {
				ans++
			}
		}
	}
	return ans
}

//DFS
func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(row, col int) (int, bool)
	dfs = func(row, col int) (int, bool) {
		if row < 0 || row >= m || col < 0 || col >= n {
			return 0, false
		}
		if grid[row][col] == 0 {
			return 0, true
		}
		grid[row][col] = 0
		ans := 1
		isNeed := true
		for _, step := range steps {
			cnt, need := dfs(row+step[0], col+step[1])
			isNeed = isNeed && need
			ans += cnt
		}
		return ans, isNeed
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt, need := dfs(i, j)
			if need {
				ans += cnt
			}
		}
	}
	return ans
}

//DFS
func numEnclaves2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	seen := make([][]bool, m)
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] == 0 || seen[row][col] {
			return
		}
		seen[row][col] = true
		for _, step := range steps {
			dfs(row+step[0], col+step[1])
		}
	}
	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !seen[i][j] {
				ans++
			}
		}
	}
	return ans
}

//并查集
func numEnclaves1(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	stepRow := []int{-1, 0, 1, 0}
	stepCol := []int{0, -1, 0, 1}
	uf := pkg.NewUnionFind(m*n + 1)
	leave := m * n
	//遍历所有单元格,先将所有相邻的陆地连通,如果陆地位于边界(4条边)则与leave连通
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				//陆地
				pos := i*n + j
				for k := 0; k < len(stepRow); k++ {
					row := i + stepRow[k]
					col := j + stepCol[k]
					if row >= 0 && row < m && col >= 0 && col < n && grid[row][col] == 1 {
						adj := row*n + col
						uf.Union(pos, adj)
					}
				}
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					//边界
					uf.Union(pos, leave)
				}
			}
		}
	}
	ans := 0
	//遍历所有单元格,如果是陆地,判断是否与leave相连,否则答案+1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !uf.Connected(i*n+j, leave) {
				ans++
			}
		}
	}
	return ans
}
