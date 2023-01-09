package leetcode_2

import "sort"

func splitArraySameAverage(nums []int) bool {
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return dfs(nums, 0, 0, 0, 0, 0)
}

func dfs(nums []int, idx, cnt1, sum1, cnt2, sum2 int) bool {
	if idx == len(nums) {
		if cnt1 == 0 || cnt2 == 0 {
			return false
		}
		return sum1/cnt1 == sum2/cnt2
	}
	return dfs(nums, idx+1, cnt1+1, sum1+nums[idx], cnt2, sum2) || dfs(nums, idx+1, cnt1, sum1, cnt2+1, sum2+nums[idx])
}
