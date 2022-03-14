package leetcode

func maxProfit121_2(prices []int) int {
	ans := 0
	//s := 0
	b := -prices[0]
	for _, price := range prices {
		if price-b > ans {
			ans = price - b
		}
		if price < b {
			b = price
		}
	}
	return ans
}
