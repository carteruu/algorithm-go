package leet

func maxProfit121_1(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	ans := 0
	s := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > s {
			s = prices[i]
		} else {
			if s-prices[i] > ans {
				ans = s - prices[i]
			}
		}
	}
	return ans
}
func maxProfit121(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-minPrice > max {
			max = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return max
}
