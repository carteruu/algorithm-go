package leetcode_1

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			idx++
			nums[idx] = nums[i]
		}
	}
	return idx + 1
}
