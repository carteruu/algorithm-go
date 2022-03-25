package leetcode_1

func moveZeroes(nums []int) {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[p] = nums[i]
			p++
		}
	}
	for ; p < len(nums); p++ {
		nums[p] = 0
	}
}
