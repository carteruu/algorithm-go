package leetcode

func wordBreak1(s string, wordDict []string) bool {
	wordSet := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = true
	}
	n := len(s)
	//dp为以 i-1 下标的字符结尾的子串,即s[0,i)能否完全拆分为字典中的字符
	dp := make([]bool, n+1)
	//默认空串可以满足条件
	dp[0] = true
	//从第一个字符开始遍历s[0,1)
	for i := 1; i <= n; i++ {
		//j为子串的起始下标
		for j := 0; j <= i; j++ {
			//状态转移:以j-1下标结尾的子串(dp[j])满足条件,且s[j,i)存在单词的列表中,则s[0,i)满足条件
			//dp[j]是s[j,i)子串的前一个字符是否满足条件
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
