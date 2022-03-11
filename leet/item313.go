package leet

import "math"

func nthSuperUglyNumber11(n int, primes []int) int {
	m := len(primes)
	dp := make([]int, n+1)
	dp[1] = 1
	idxs := make([]int, m)
	for i := 0; i < m; i++ {
		idxs[i] = 1
	}
	temp := make([]int, m)
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt64
		for j := 0; j < m; j++ {
			val := dp[idxs[j]] * primes[j]
			temp[j] = val
			if val < dp[i] {
				dp[i] = val
			}
		}
		for j, val := range temp {
			if val == dp[i] {
				idxs[j]++
			}
		}
	}
	return dp[n]
}

func nthSuperUglyNumber(n int, primes []int) int {
	res := make([]int, n)
	res[0] = 1
	index := make([]int, len(primes))
	for i := 1; i < n; i++ {
		min := int(math.MaxInt32)
		for j, v := range primes {
			if min > v*res[index[j]] {
				min = v * res[index[j]]
			}
		}
		for j, v := range primes {
			if v*res[index[j]] == min {
				index[j]++
			}
		}
		res[i] = min
	}
	return res[n-1]
}
