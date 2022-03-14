package leetcode

import (
	"math"
)

func countMaxOrSubsets(nums []int) int {
	maxNum := math.MinInt64
	cnt := 0
	var dfs func(idx, or, numCnt int)
	dfs = func(idx, or, numCnt int) {
		if idx == len(nums) {
			if numCnt > 0 {
				if or > maxNum {
					maxNum = or
					cnt = 0
				}
				if or == maxNum {
					cnt++
				}
			}
			return
		}
		//不选择
		dfs(idx+1, or, numCnt)
		//选择
		dfs(idx+1, or|nums[idx], numCnt+1)
	}
	dfs(0, 0, 0)
	return cnt
}
