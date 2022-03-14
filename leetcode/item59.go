package leetcode

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	top, bottom, left, right := 0, n-1, 0, n-1
	row, col := 0, 0
	steps := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	stepIdx := 0
	for i := 1; i <= n*n; i++ {
		ans[row][col] = i
		nextRow := row + steps[stepIdx][0]
		nextCol := col + steps[stepIdx][1]
		if nextRow > bottom {
			//向下,越界
			stepIdx++
			stepIdx = stepIdx % 4
			right--
		} else if nextCol > right {
			//向右,越界
			stepIdx++
			stepIdx = stepIdx % 4
			top++
		} else if nextRow < top {
			//向上,越界
			stepIdx++
			stepIdx = stepIdx % 4
			left++
		} else if nextCol < left {
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
