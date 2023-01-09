package leetcode_2

func check(nums []int) bool {
	d := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			d++
		}
	}
	if d == 1 && nums[0] >= nums[len(nums)-1] {
		return true
	}
	return d == 0
}
