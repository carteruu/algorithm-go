package leet

func maxProfit(prices []int, fee int) int {
	//第一天,不持有股票状态金额为0
	dp := 0
	//第一天,持有股票状态金额为负的第一天股票价格
	dp1 := -prices[0]
	for i := 1; i < len(prices); i++ {
		//当天不持有股票状态:保持不持有股票状态不变,或从持有股票状态卖出股票
		t := max(dp, dp1+prices[i]-fee)
		//当天持有股票状态:保持持有股票状态不变,或从不持有股票状态买入股票
		t1 := max(dp1, dp-prices[i])
		dp = t
		dp1 = t1
	}
	//因为只考虑现金,所以取最后一天不持有股票的金额
	return dp
}
