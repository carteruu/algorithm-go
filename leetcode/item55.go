package leetcode

func canJump(nums []int) bool {
	max := 0
	idx := 0
	for idx <= max {
		if idx+nums[idx] > max {
			max = idx + nums[idx]
			if max >= len(nums)-1 {
				return true
			}
		}
		idx++
	}
	return false
}
