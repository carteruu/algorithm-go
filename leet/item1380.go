package leet

func luckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	rowMin := make([]int, m)
	colMax := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowMin[i] == 0 || matrix[i][j] < rowMin[i] {
				rowMin[i] = matrix[i][j]
			}
			if matrix[i][j] > colMax[j] {
				colMax[j] = matrix[i][j]
			}
		}
	}
	var ans []int
	for _, r := range rowMin {
		for _, c := range colMax {
			if c == r {
				ans = append(ans, c)
			}
		}
	}
	return ans
}
