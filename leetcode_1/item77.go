package leetcode_1

import (
	"math"
	"math/bits"
)

/**
用位标识每一个数，是否应该选择
*/
func combine(n int, k int) [][]int {
	l := 1
	for i := 0; i < k; i++ {
		l = l * (n - i) / (i + 1)
	}
	ans := make([][]int, 0, l)
	for i := 0; i < 1<<n; i++ {
		if bits.OnesCount32(uint32(i)) != k {
			continue
		}
		item := make([]int, 0, k)
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				item = append(item, j+1)
			}
		}
		ans = append(ans, item)
	}
	return ans
}
func combine_1(n int, k int) [][]int {
	ans := make([][]int, 0, int(math.Pow(float64(n), float64(k))))
	var dfs func(idx int, path []int)
	dfs = func(num int, path []int) {
		if len(path) == k {
			item := make([]int, k)
			copy(item, path)
			ans = append(ans, item)
			return
		}
		if n-num+1 < k-len(path) {
			//剩余的数，不够了
			return
		}
		//选
		path = append(path, num)
		dfs(num+1, path)
		path = path[:len(path)-1]
		//不选
		dfs(num+1, path)
	}
	dfs(1, make([]int, 0, k))
	return ans
}
