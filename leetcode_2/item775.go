package leetcode_2

func isIdealPermutation1(nums []int) bool {
	return false
}
func isIdealPermutation(nums []int) bool {
	p, a := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			p++
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				a++
			}
		}
	}
	return p == a
}
