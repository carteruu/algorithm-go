package leet

func lengthOfLongestSubstring(s string) int {
	cntMap := make(map[byte]int)
	slow, ans := 0, 0
	for fast := 0; fast < len(s); fast++ {
		cntMap[s[fast]]++
		for cntMap[s[fast]] > 1 {
			cntMap[s[slow]]--
			slow++
		}
		if fast-slow+1 > ans {
			ans = fast - slow + 1
		}
	}
	return ans
}
