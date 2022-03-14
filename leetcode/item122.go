package leetcode

func maxProfitItem122(prices []int) int {
	//第一天,不持有股票状态金额为0
	dp := 0
	//第一天,持有股票状态金额为负的第一天股票价格
	dp1 := -prices[0]
	for i := 1; i < len(prices); i++ {
		//当天不持有股票状态:保持不持有股票状态不变,或从持有股票状态卖出股票
		t := max(dp, dp1+prices[i])
		//当天持有股票状态:保持持有股票状态不变,或从不持有股票状态买入股票
		t1 := max(dp1, dp-prices[i])
		dp = t
		dp1 = t1
	}
	//因为只考虑现金,所以取最后一天不持有股票的金额
	return dp
}
func maxProfit1(prices []int) int {
	dp := [2]map[int]int{make(map[int]int, len(prices)), make(map[int]int, len(prices))}
	var dfs func(day int, bought int) int
	dfs = func(day int, bought int) int {
		if day == len(prices) {
			return 0
		}
		if val, ok := dp[bought][day]; ok {
			return val
		}
		//不操作
		val := dfs(day+1, bought)
		if bought == 1 {
			//已买入状态
			//卖出
			val = max(val, dfs(day+1, 0)+prices[day])
		} else {
			//卖出状态
			//买入
			val = max(val, dfs(day+1, 1)-prices[day])
		}
		dp[bought][day] = val
		return val
	}
	return dfs(0, 0)
}
