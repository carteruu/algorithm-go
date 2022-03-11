package leet

func totalNQueens(n int) int {
	ans := 0
	cache := make([]bool, n)
	path := make([]int, n)
	var check func(row, col int) bool
	check = func(row, col int) bool {
		//同一行和同一列都不能有Q
		colSigh := 1 << col
		sigh := 1
		for i := 0; i < n; i++ {
			sigh <<= 1
			if path[i]&colSigh > 0 || path[row]&sigh > 0 {
				return false
			}
		}
		//对角线-左上
		for i := 1; i <= row && i <= col; i++ {
			if path[row-i]&(colSigh>>i) > 0 {
				return false
			}
		}
		//对角线-右上
		for i := 1; i <= row && i < n-col; i++ {
			if path[row-i]&(colSigh<<i) > 0 {
				return false
			}
		}
		return true
	}
	var backtrack func(num, row int)
	backtrack = func(num, row int) {
		if num == n {
			ans++
			return
		}
		for j := 0; j < n; j++ {
			if cache[j] || !check(row, j) {
				continue
			}
			path[row] |= 1 << j
			cache[j] = true
			backtrack(num+1, row+1)
			path[row] ^= 1 << j
			cache[j] = false
		}
	}
	backtrack(0, 0)
	return ans
}
