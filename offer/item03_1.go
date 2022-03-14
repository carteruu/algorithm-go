package offer

func findRepeatNumber1(nums []int) int {
	m := make([]bool, len(nums))
	for _, num := range nums {
		if m[num] {
			return num
		}
		m[num] = true
	}
	return -1
}

func findRepeatNumber2(nums []int) int {
	for i, num := range nums {
		if i == num {
			continue
		}
		if num == nums[num] {
			return num
		}
		nums[num], nums[i] = nums[i], nums[num]
	}
	return -1
}
func findRepeatNumber(nums []int) int {
	for idx, num := range nums {
		if idx == num {
			continue
		}
		if num == nums[num] {
			return num
		}
		nums[num], nums[idx] = nums[idx], nums[num]
	}
	return 0
}
