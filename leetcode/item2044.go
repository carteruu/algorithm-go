package leetcode

import (
	"math"
)

// 还能继续优化：
//1. 最大值必然是所有数字与的结果；
//2. 当回溯到某个数字的与结果等于最大值时，不需要继续递归，直接将2^(n-i)加到答案
func countMaxOrSubsets(nums []int) int {
	maxNum := math.MinInt64
	cnt := 0
	var dfs func(idx, or int)
	dfs = func(idx, or int) {
		if idx == len(nums) {
			if or > maxNum {
				maxNum = or
				cnt = 1
			} else if or == maxNum {
				cnt++
			}
			return
		}
		//不选择
		dfs(idx+1, or)
		//选择
		dfs(idx+1, or|nums[idx])
	}
	dfs(0, 0)
	return cnt
}
