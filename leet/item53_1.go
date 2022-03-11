package leet

import "math"

func maxSubArray(nums []int) int {
	ans := math.MinInt64
	sum := 0
	for _, num := range nums {
		sum += num
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}
