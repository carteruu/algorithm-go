package leetcode_2

//暴力
func partitionDisjoint(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		lMax := nums[0]
		rMin := nums[len(nums)-1]
		for j := 0; j <= i; j++ {
			lMax = maxInt(lMax, nums[j])
		}
		for j := len(nums) - 1; j > i; j-- {
			rMin = minInt(rMin, nums[j])
		}
		if lMax <= rMin {
			return i + 1
		}
	}
	return -1
}

func partitionDisjoint1(nums []int) int {
	lMax := nums[0]
	rMin := make([]int, 1, len(nums)-1)
	rMin[0] = nums[len(nums)-1]
	for i := len(nums) - 2; i > 0; i-- {
		rMin = append(rMin, minInt(rMin[len(rMin)-1], nums[i]))
	}
	for i := 1; i < len(nums); i++ {
		if rMin[len(rMin)-i] >= lMax {
			return i
		}
		lMax = maxInt(lMax, nums[i])
	}
	return -1
}
