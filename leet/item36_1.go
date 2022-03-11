package leet

func isValidSudoku(board [][]byte) bool {
	acrossCache := [9][9]bool{}
	verticalCache := [9][9]bool{}
	cellCache := [9][9]bool{}
	for i := 0; i < 9; i++ {
		acrossCache[i] = [9]bool{}
		verticalCache[i] = [9]bool{}
		cellCache[i] = [9]bool{}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				if acrossCache[i][num] {
					return false
				}

				acrossCache[i][num] = true
				if verticalCache[j][num] {
					return false
				}
				verticalCache[j][num] = true

				if cellCache[j/3*3+i/3][num] {
					return false
				}
				cellCache[j/3*3+i/3][num] = true
			}
		}
	}
	return true
}
