package leetcode_2

func minimumBoxes(n int) int {
	var dp []int
	dp = append(dp, 0)
	for next := dp[len(dp)-1] + len(dp); next <= n; next = dp[len(dp)-1] + len(dp) {
		dp = append(dp, next)
		n -= next
	}
	ans := dp[len(dp)-1]
	for n > 0 {
		is := true
		for i := len(dp) - 1; i >= 0; i-- {
			if is {
				ans += dp[i]
				is = false
			}
			n -= dp[i]
		}
	}
	return ans
}
