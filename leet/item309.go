package leet

//func maxProfit309_1(prices []int) int {
//	//第一天,不持有股票状态金额为0
//	sold := 0
//	soldDay := -1
//	//第一天,持有股票状态金额为负的第一天股票价格
//	bought := -prices[0]
//
//	for i := 1; i < len(prices); i++ {
//		//当天不持有股票状态:保持不持有股票状态不变,或从持有股票状态卖出股票
//		t := max(sold, bought+prices[i])
//		if bought+prices[i]>sold
//		//当天持有股票状态:保持持有股票状态不变,或从不持有股票状态买入股票
//		t1 := max(bought, sold-prices[i])
//		sold = t
//		bought = t1
//	}
//	//因为只考虑现金,所以取最后一天不持有股票的金额
//	return sold
//}

func maxProfit309(prices []int) int {
	dp := [2][2]map[int]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			dp[i][j] = make(map[int]int, len(prices))
		}
	}
	var dfs func(day int, bought, preSell int) int
	dfs = func(day int, bought, preSell int) int {
		if day == len(prices) {
			return 0
		}
		if val, ok := dp[bought][preSell][day]; ok {
			return val
		}
		//不操作
		val := dfs(day+1, bought, 0)
		if bought == 1 {
			//已买入状态
			//卖出
			val = max(val, dfs(day+1, 0, 1)+prices[day])
		}
		if bought == 0 && preSell == 0 {
			//卖出状态
			//买入
			val = max(val, dfs(day+1, 1, 0)-prices[day])
		}
		dp[bought][preSell][day] = val
		return val
	}
	return dfs(0, 0, 0)
}

func maxProfit11(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		nf0 := max(f0, f2-prices[i])
		nf1 := f0 + prices[i]
		nf2 := max(f1, f2)
		f0, f1, f2 = nf0, nf1, nf2
	}
	return max(f1, f2)
}
