package offer_1

import "sort"

func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

func findRepeatNumber1(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		}
		m[num] = struct{}{}
	}
	return -1
}
func findRepeatNumber2(nums []int) int {
	for i := range nums {
		for i != nums[i] {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
func findRepeatNumber3(nums []int) int {
	m := make([]bool, len(nums))
	for _, num := range nums {
		if m[num] {
			return num
		}
		m[num] = true
	}
	return -1
}
