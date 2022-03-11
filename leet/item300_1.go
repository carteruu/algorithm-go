package leet

func lengthOfLIS1(nums []int) int {
	dp := make([]int, len(nums))
	ans := 0
	for i, num := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < num {
				if dp[j] >= dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
