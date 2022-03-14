package leetcode

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	neighbor := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == 'X' || board[i][j] == 'A' {
			return
		}
		board[i][j] = 'A'
		for _, nb := range neighbor {
			dfs(i+nb[0], j+nb[1])
		}
		return
	}

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(m-1, j)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}
func solve1(board [][]byte) {
	m := len(board)
	n := len(board[0])
	notChange := make([][]bool, m)
	for i := 0; i < m; i++ {
		notChange[i] = make([]bool, n)
	}
	neighbor := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == 'X' || notChange[i][j] {
			return
		}
		notChange[i][j] = true
		for _, nb := range neighbor {
			dfs(i+nb[0], j+nb[1])
		}
		return
	}

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(m-1, j)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !notChange[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
