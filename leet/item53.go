package leet

import "math"

func maxSubArray11(nums []int) int {
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
func maxSubArray12(nums []int) int {
	ans := math.MinInt64
	temp := 0
	slow, fast := 0, 0
	for fast < len(nums) {
		temp += nums[fast]
		if temp > ans {
			ans = temp
		}
		for slow <= fast && temp < 0 {
			temp -= nums[slow]
			slow++
		}
		fast++
	}
	return ans
}
