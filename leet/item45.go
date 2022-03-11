package leet

import "math"

func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func jump2(nums []int) int {
	position := len(nums) - 1
	steps := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}

func jump1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt64
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= nums[i] && j < n-i; j++ {
			tIdx := i + j
			if dp[i]+1 < dp[tIdx] {
				dp[tIdx] = dp[i] + 1
			}
		}
	}
	return dp[n-1]
}
