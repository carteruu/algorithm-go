package leet

import "math"

func maxProfit123(prices []int) int {
	s1, b1 := 0, -prices[0]
	s2, b2 := 0, math.MinInt64
	for i := 1; i < len(prices); i++ {
		s1 = max(s1, b1+prices[i])
		b1 = max(b1, -prices[i])

		s2 = max(s2, b2+prices[i])
		b2 = max(b2, s1-prices[i])
	}
	return max(s1, s2)
}
