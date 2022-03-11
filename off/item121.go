package off

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ans := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > ans {
			ans = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return ans
}
func maxProfit1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	mins := make([]int, n)
	maxs := make([]int, n)
	mins[0] = prices[0]
	maxs[n-1] = prices[n-1]
	for i := 1; i < n; i++ {
		if prices[i] < mins[i-1] {
			mins[i] = prices[i]
		} else {
			mins[i] = mins[i-1]
		}

		curIdx := n - 1 - i
		if prices[curIdx] > maxs[n-i] {
			maxs[curIdx] = prices[curIdx]
		} else {
			maxs[curIdx] = maxs[n-i]
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		profit := maxs[i] - mins[i]
		if profit > ans {
			ans = profit
		}
	}
	return ans
}
