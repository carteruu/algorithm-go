package leetcode_2

func minOperations_1(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	ans := 0
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= pre {
			ans += pre + 1 - nums[i]
			pre = pre + 1
		}
	}
	return ans
}
func minOperations(nums []int) int {
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			ans += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return ans
}
