package leet

import "algorithm/pkg"

func longestCommonSubsequence1(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = pkg.Max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, len1+1)
	dp[0] = make([]int, len2+1)
	for idx1, str1 := range text1 {
		dp[idx1+1] = make([]int, len2+1)
		for idx2, str2 := range text2 {
			if str1 == str2 {
				dp[idx1+1][idx2+1] = dp[idx1][idx2] + 1
			} else {
				dp[idx1+1][idx2+1] = pkg.Max(dp[idx1+1][idx2], dp[idx1][idx2+1])
			}
		}
	}
	return dp[len1][len2]
}
