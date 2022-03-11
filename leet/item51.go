package leet

func solveNQueens(n int) [][]string {
	var ans [][]string
	cache := make([]bool, n)
	path := make([][]byte, n)
	for i := 0; i < n; i++ {
		path[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			path[i][j] = '.'
		}
	}
	var check func(row, col int) bool
	check = func(row, col int) bool {
		//同一行和同一列都不能有Q
		for i := 0; i < n; i++ {
			if i != col && path[row][i] == 'Q' {
				return false
			}
			if i != row && path[i][col] == 'Q' {
				return false
			}
		}
		//对角线-左上
		for i := 1; i <= row && i <= col; i++ {
			if path[row-i][col-i] == 'Q' {
				return false
			}
		}
		//对角线-右上
		for i := 1; i <= row && i < n-col; i++ {
			if path[row-i][col+i] == 'Q' {
				return false
			}
		}
		return true
	}
	var backtrack func(num, row int)
	backtrack = func(num, row int) {
		if num == n {
			aa := make([]string, n)
			d := ""
			for i := 0; i < n; i++ {
				aa[i] = string(path[i])
				d += aa[i]
			}
			ans = append(ans, aa)
			return
		}
		for j := 0; j < n; j++ {
			if cache[j] || path[row][j] == 'Q' {
				continue
			}
			if check(row, j) {
				path[row][j] = 'Q'
				cache[j] = true
				backtrack(num+1, row+1)
				path[row][j] = '.'
				cache[j] = false
			}
		}
	}
	backtrack(0, 0)
	return ans
}
