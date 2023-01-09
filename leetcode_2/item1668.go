package leetcode_2

func maxRepeating(sequence string, word string) int {
	dp := make([]int, len(sequence)+1)
	ans := 0
	for i := len(word); i <= len(sequence); i++ {
		if sequence[i-len(word):i] == word {
			dp[i] = dp[i-len(word)] + 1
			ans = maxInt(ans, dp[i])
		}
	}
	return ans
}
func maxRepeating1(sequence string, word string) int {
	ans := 0
	for i := range sequence {
		for j := i; j <= len(sequence); j++ {
			if j == len(sequence) || sequence[j] != word[(j-i)%len(word)] {
				ans = maxInt(ans, (j-i)/len(word))
				break
			}
		}
	}
	return ans
}
