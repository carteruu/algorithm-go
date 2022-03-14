package interview

import (
	"math"
)

func getMaxMatrix(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	max := math.MinInt32
	var ans []int
	//先计算出每一列的前缀和
	preSum := make([][]int, m+1)
	preSum[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		preSum[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			//前缀和
			preSum[i][j] = matrix[i-1][j-1] + preSum[i-1][j]
		}
	}
	//上边界,包含
	for t := 0; t < m; t++ {
		//下边界,包含
		for b := t; b < m; b++ {
			//双指针
			l := 0
			sum := 0
			for r := 0; r < n; r++ {
				//使用前缀和快速计算出当前列的和,累加
				sum += preSum[b+1][r+1] - preSum[t][r+1]
				if sum > max {
					//更新答案
					max = sum
					ans = []int{t, l, b, r}
				}
				if sum < 0 {
					//如果和小于0,抛弃前面的元素,慢指针从下一个元素开始
					l = r + 1
					sum = 0
				}
			}
		}
	}
	return ans
}
func getMaxMatrix1(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	//计算出二维前缀和,preSum[i][j]为从matrix[0][0]到matrix[i][j]的和
	preSum := make([][]int, m+1)
	preSum[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		preSum[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			//左边的矩阵+右边的矩阵-重复计算部分
			preSum[i][j] = matrix[i-1][j-1] + preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1]
		}
	}
	max := math.MinInt32
	var ans []int
	for t := 0; t < m; t++ {
		for b := t; b < m; b++ {
			//双指针
			l := 0
			for r := 0; r < n; r++ {
				//计算当前矩阵的和
				lm := preSum[b+1][r+1] - preSum[b+1][l] - preSum[t][r+1] + preSum[t][l]
				if lm > max {
					//更新答案
					max = lm
					ans = []int{t, l, b, r}
				}
				if lm < 0 {
					//如果和小于0,抛弃前面部分,重新开始
					l = r + 1
				}
			}
		}
	}
	return ans
}

func getMaxMatrix2(matrix [][]int) []int {
	res := make([]int, 4)
	maxSum := math.MinInt32
	m, n := len(matrix), len(matrix[0])
	for l := 0; l < n; l++ {
		rowSum := make([]int, m)
		for r := l; r < n; r++ {
			for i := 0; i < m; i++ {
				rowSum[i] += matrix[i][r]
			}
			up := 0
			sum := 0
			for down, v := range rowSum {
				if sum < 0 {
					up = down
					sum = 0
				}
				sum += v
				if maxSum < sum {
					maxSum = sum
					res = []int{up, l, down, r}
				}
			}
		}
	}
	return res
}
