package leetcode_1

/**
贪心：从第二天开始，如果当天的价格比昨天高，则昨天买入，今天卖出，赚取差价
*/
func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

/**
动态规划
*/
func maxProfit_1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	bought, sold := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		//状态转换：卖出 -> 买入
		bought = max(bought, sold-prices[i])
		//状态转换：买入 -> 卖出
		sold = max(sold, bought+prices[i])
	}
	return sold
}
