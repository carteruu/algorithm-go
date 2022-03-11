package leet

func countKDifference(nums []int, k int) int {
	ans := 0
	cnt := make([]int, 300)
	for _, num := range nums {
		//num+k的取值范围为[2,199]
		ans += cnt[num+k+100]
		//num-k的取值范围为[-98,99]
		ans += cnt[num-k+100]
		cnt[num+100]++
	}
	return ans
}
func countKDifference1(nums []int, k int) int {
	ans := 0
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			x := num - nums[j]
			if x < 0 {
				x = -x
			}
			if x == k {
				ans++
			}
		}
	}
	return ans
}
