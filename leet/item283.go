package leet

func moveZeroes(nums []int) []int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		}
	}
	return nums
}
