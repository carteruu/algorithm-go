package leet

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0, m*n)
	top, bottom, left, right := 0, m-1, 0, n-1
	row, col := 0, 0
	steps := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	stepIdx := 0
	for i := 0; i < m*n; i++ {
		ans = append(ans, matrix[row][col])
		if row+steps[stepIdx][0] > bottom {
			//向下,越界
			stepIdx++
			stepIdx = stepIdx % 4
			right--
		} else if col+steps[stepIdx][1] > right {
			//向右,越界
			stepIdx++
			stepIdx = stepIdx % 4
			top++
		} else if row+steps[stepIdx][0] < top {
			//向上,越界
			stepIdx++
			stepIdx = stepIdx % 4
			left++
		} else if col+steps[stepIdx][1] < left {
			//向左,越界
			stepIdx++
			stepIdx = stepIdx % 4
			bottom--
		}
		row += steps[stepIdx][0]
		col += steps[stepIdx][1]
	}
	return ans
}
