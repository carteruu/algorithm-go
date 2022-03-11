package leet

func isValidSudoku11(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		rowMap := make([]bool, 9)
		colMap := make([]bool, 9)
		m := make([]bool, 9)
		for j := 0; j < 9; j++ {
			c1 := board[i][j]
			if c1 != '.' {
				if rowMap[c1-'1'] {
					return false
				} else {
					rowMap[c1-'1'] = true
				}
			}
			c2 := board[j][i]
			if c2 != '.' {
				if colMap[c2-'1'] {
					return false
				} else {
					colMap[c2-'1'] = true
				}
			}
			c := board[i/3*3+j/3][i%3*3+j%3]
			if c != '.' {
				if m[c-'1'] {
					return false
				} else {
					m[c-'1'] = true
				}
			}
		}
	}
	return true
}
func isValidSudoku1(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		rowMap := make(map[byte]struct{}, 9)
		colMap := make(map[byte]struct{}, 9)
		cellMap := make(map[byte]struct{}, 9)
		for j := 0; j < 9; j++ {
			c1 := board[i][j]
			if _, ok := rowMap[c1]; c1 != '.' && ok {
				return false
			}
			c2 := board[j][i]
			if _, ok := colMap[c2]; c2 != '.' && ok {
				return false
			}
			c3 := board[i%3+j/3][(i%3)*3+j%3]
			if _, ok := cellMap[c3]; c3 != '.' && ok {
				return false
			}
			rowMap[c1] = struct{}{}
			colMap[c2] = struct{}{}
			cellMap[c3] = struct{}{}
		}
	}
	return true
}
