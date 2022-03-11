package leet

func majorityElement(nums []int) int {
	cnt := 1
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			ans = nums[i]
			cnt++
		} else {
			if ans == nums[i] {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return ans
}
