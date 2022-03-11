package off

func findNumberIn2DArray(matrix [][]int, target int) bool {
	row, col := len(matrix)-1, 0
	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			row--
		} else {
			col++
		}
	}
	return false
}
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var find func(top, bottom, left, right int) bool
	find = func(top, bottom, left, right int) bool {
		if top > bottom || left > right {
			return false
		}
		midRow, midCol := (top+bottom)/2, (left+right)/2
		if target == matrix[midRow][midCol] {
			return true
		} else if target < matrix[midRow][midCol] {
			return find(top, midRow-1, left, right) || find(midRow, bottom, left, midCol-1)
		} else {
			return find(midRow+1, bottom, left, right) || find(top, midRow, midCol+1, right)
		}
	}
	return find(0, len(matrix)-1, 0, len(matrix[0])-1)
}
