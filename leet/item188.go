package leet

func maxProfit188(k int, prices []int) int {
	if prices == nil || len(prices) == 0 || k == 0 {
		return 0
	}
	dp := make([][2]int, k+1)
	for i := 1; i <= k; i++ {
		dp[i][1] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
		}
	}
	return dp[k][0]
}
