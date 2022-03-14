package leetcode

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
func change1(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		temp := make([]int, amount+1)
		for i := 0; i <= amount; i++ {
			temp[i] = dp[i]
			if i >= coin {
				temp[i] += temp[i-coin]
			}
		}
		dp = temp
	}
	return dp[amount]
}
