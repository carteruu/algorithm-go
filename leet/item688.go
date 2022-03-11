package leet

import "math"

var steps = [][]int{{-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}
var dirs = []struct{ i, j int }{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}

func knightProbability(n, k, row, column int) float64 {
	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					for _, d := range dirs {
						if x, y := i+d.i, j+d.j; 0 <= x && x < n && 0 <= y && y < n {
							dp[step][i][j] += dp[step-1][x][y] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}

//DP
func knightProbability3(n int, k int, row int, column int) float64 {
	if row < 0 || row >= n || column < 0 || column >= n {
		return 0
	}
	dp := make([][][]float64, k+1)
	dp[0] = make([][]float64, n)
	for i := 0; i < n; i++ {
		dp[0][i] = make([]float64, n)
	}
	dp[0][row][column] = 1
	for s := 1; s <= k; s++ {
		dp[s] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[s][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				for _, step := range steps {
					r := i - step[0]
					c := j - step[1]
					if r >= 0 && r < n && c >= 0 && c < n {
						dp[s][i][j] += dp[s-1][r][c]
					}
				}
			}
		}
	}
	var in float64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			in += dp[k][i][j]
		}
	}
	return in / math.Pow(float64(len(steps)), float64(k))
}

//DFS
func knightProbability2(n int, k int, row int, column int) float64 {
	var dfs func(k int, row int, column int) int
	dfs = func(k int, row int, column int) int {
		ans := 0
		if row < 0 || row >= n || column < 0 || column >= n {
			return 0
		}
		if k == 0 {
			return 1
		}
		for _, step := range steps {
			r := row + step[0]
			c := column + step[1]
			ans += dfs(k-1, r, c)
		}
		return ans
	}
	return float64(dfs(k, row, column)) / math.Pow(float64(len(steps)), float64(k))
}

//BFS
func knightProbability1(n int, k int, row int, column int) float64 {
	var q [][]int
	q = append(q, []int{row, column})
	for i := 0; i < k; i++ {
		size := len(q)
		for j := 0; j < size; j++ {
			poll := q[len(q)-1]
			q = q[:len(q)-1]
			for _, step := range steps {
				r := poll[0] + step[0]
				c := poll[1] + step[1]
				if r < 0 || r >= n || c < 0 || c >= n {
					continue
				}
				q = append(q, []int{r, c})
			}
		}
	}
	return float64(len(q)) / math.Pow(float64(len(steps)), float64(k))
}
