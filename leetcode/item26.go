package leetcode

func removeDuplicates12(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	slow, fast := 1, 1
	for fast < n {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func removeDuplicates11(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	slow, fast := 1, 1
	for fast < n {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
