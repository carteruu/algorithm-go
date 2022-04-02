package leetcode_1

func solveNQueens_2(n int) [][]string {
	//记录各行是否存在皇后
	rowState := 0
	//记录各列是否存在皇后
	colState := 0
	//记录各条对角线是否存在皇后
	dia1State := 0
	dia2State := 0
	var ans [][]string
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	var dfs func(idx, cnt int)
	dfs = func(idx, cnt int) {
		row := idx / n
		col := idx % n
		if row > cnt {
			//剪枝：当某一行已走完，且没有放皇后时，不符合要求
			return
		}
		if idx == n*n {
			if cnt == n {
				b := make([]string, n)
				for i := 0; i < n; i++ {
					b[i] = string(board[i])
				}
				ans = append(ans, b)
			}
			return
		}
		//对角线的斜率为 1，row = col + k => k = row - col，为了避免 k 小于0，上移 n 个位置
		k1 := row - col + n
		//对角线的斜率为 -1，row = -col + k
		k2 := row + col
		rowMark := 1 << row
		colMark := 1 << col
		k1Mark := 1 << k1
		k2Mark := 1 << k2
		if (rowState&rowMark)|(colState&colMark)|(dia1State&k1Mark)|(dia2State&k2Mark) == 0 {
			//当前位置可以放皇后
			rowState |= rowMark
			colState |= colMark
			dia1State |= k1Mark
			dia2State |= k2Mark
			board[row][col] = 'Q'
			dfs(idx+1, cnt+1)
			board[row][col] = '.'
			rowState ^= rowMark
			colState ^= colMark
			dia1State ^= k1Mark
			dia2State ^= k2Mark
		}
		dfs(idx+1, cnt)
	}
	dfs(0, 0)
	return ans
}

/**
对每个位置进行检查，可以放皇后的话，有两种选择：放皇后或不放；否则只能选择不放
*/
func solveNQueens(n int) [][]string {
	var ans [][]string
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	var dfs func(idx, cnt int)
	var check func(row, col int) bool
	dfs = func(idx, cnt int) {
		if idx == n*n {
			if cnt == n {
				b := make([]string, n)
				for i := 0; i < n; i++ {
					b[i] = string(board[i])
				}
				ans = append(ans, b)
			}
			return
		}
		row := idx / n
		col := idx % n
		if check(row, col) {
			board[row][col] = 'Q'
			dfs(idx+1, cnt+1)
			board[row][col] = '.'
		}
		dfs(idx+1, cnt)
	}
	check = func(row, col int) bool {
		//因为是row和col都是从0到n移动的，所以只需要考虑前面
		//行,列
		for i := 0; i < col; i++ {
			if board[row][i] == 'Q' {
				return false
			}
		}
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}
		//对角线
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		return true
	}
	dfs(0, 0)
	return ans
}
