package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j]
		}
	}
}
func rotate1(matrix [][]int) {
	n := len(matrix)
	temp := make([][]int, n)
	for i := 0; i < n; i++ {
		temp[i] = make([]int, n)
		copy(temp[i], matrix[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][n-1-i] = temp[i][j]
		}
	}
}

func rotate2(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			tmp[j][n-1-i] = v
		}
	}
	copy(matrix, tmp) // 拷贝 tmp 矩阵每行的引用
}

func rotate3(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
