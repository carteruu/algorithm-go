package leetcode

import "math"

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left := -1
	right := -1

	min := math.MaxInt64
	max := math.MinInt64
	for i := n - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			left = i
		}
	}
	if left == -1 {
		return 0
	}

	for i := 0; i < n; i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			right = i
		}
	}
	return right - left + 1
}
