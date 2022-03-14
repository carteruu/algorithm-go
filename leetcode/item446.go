package leetcode

func numberOfArithmeticSlices111(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	res := 0
	d, t := nums[1]-nums[0], 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == d {
			t++
			res += t
		} else {
			d = nums[i] - nums[i-1]
			t = 0
		}
	}
	return res
}
