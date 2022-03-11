package leet

func nextPermutation(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}
	ordIdx := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			ordIdx = i - 1
			break
		}
	}
	if ordIdx >= 0 {
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] > nums[ordIdx] {
				nums[ordIdx], nums[i] = nums[i], nums[ordIdx]
				break
			}
		}
	}
	rev1(nums, ordIdx+1)
}

func rev1(nums []int, idx int) {
	for i := 0; i < (len(nums)-idx)/2; i++ {
		nums[idx+i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[idx+i]
	}
}
