package leetcode_1

/**
动态规划
*/
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p1 := 1
	p2 := 1
	p3 := 1
	for i := 2; i <= n; i++ {
		num1 := dp[p1] * 2
		num2 := dp[p2] * 3
		num3 := dp[p3] * 5
		dp[i] = min(min(num1, num2), num3)
		//num1、num2、num3可能会相同
		if dp[i] == num1 {
			p1++
		}
		if dp[i] == num2 {
			p2++
		}
		if dp[i] == num3 {
			p3++
		}
	}
	return dp[n]
}
