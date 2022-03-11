package leet

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	maxSize := 1
	var maxIdx int
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[i] <= dp[j] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize = dp[i]
			maxIdx = i
		}
	}

	maxVal := nums[maxIdx]
	ans := make([]int, maxSize)
	for i := n - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			maxSize -= 1
			ans[maxSize] = nums[i]
			maxVal = nums[i]
		}
	}
	return ans
}

func largestDivisibleSubset1(nums []int) (res []int) {
	sort.Ints(nums)

	// 第 1 步：动态规划找出最大子集的个数、最大子集中的最大整数
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxVal := 1, 1
	for i := 1; i < n; i++ {
		for j, v := range nums[:i] {
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}

	if maxSize == 1 {
		return []int{nums[0]}
	}

	// 第 2 步：倒推获得最大子集
	for i := n - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return
}

func largestDivisibleSubset2(nums []int) []int {
	n, maxEndIndex, maxLen := len(nums), 0, 0
	dp_table := make([]int, n)
	ans := []int{}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		dp_table[i] = 1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp_table[i] = max(dp_table[j]+1, dp_table[i])
			}
		}
		if dp_table[i] > maxLen {
			maxEndIndex = i
			maxLen = dp_table[i]
		}
	}
	maxEnd := nums[maxEndIndex]
	for i := maxEndIndex; i >= 0; i-- {
		if maxEnd%nums[i] == 0 && maxLen == dp_table[i] {
			maxLen--
			maxEnd = nums[i]
			ans = append(ans, nums[i])
		}
	}
	return ans
}
