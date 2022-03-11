package leet

func integerBreak(n int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], j*dp[i-j])
			dp[i] = max(dp[i], j*(i-j))
		}
	}
	return dp[n]
}
func integerBreak1(n int) int {
	if n < 2 {
		return -1
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	var m = make([]int, 59)
	m[1] = 1
	m[2] = 2
	var backtrack func(r int) int
	backtrack = func(r int) int {
		if r < 1 {
			return 0
		}
		if m[r] > 0 {
			return m[r]
		}
		max := r
		for i := 2; i <= r/2; i++ {
			num := i * backtrack(r-i)
			if num > max {
				max = num
			}
		}
		m[r] = max
		return max
	}
	return backtrack(n)
}
