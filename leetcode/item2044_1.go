package leetcode

func countMaxOrSubsets_1(nums []int) int {
	ans := 0
	mx := 0
	for i := 1; i < 1<<len(nums); i++ {
		mark := i
		or := 0
		for j := 0; j < len(nums); j++ {
			if mark&(1<<j) > 0 {
				or |= nums[j]
			}
		}
		if or > mx {
			ans = 1
			mx = or
		} else if or == mx {
			ans++
		}
	}
	return ans
}
