package leetcode_2

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ans := 0
	l := prices[0]
	for i := 1; i < len(prices); i++ {
		ans = maxInt(ans, prices[i]-l)
		l = minInt(l, prices[i])
	}
	return ans
}
