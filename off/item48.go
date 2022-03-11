package off

import "math"

func lengthOfLongestSubstring(s string) int {
	var ans, slow int
	cnt := make([]int, math.MaxUint8)
	for fast := 0; fast < len(s); fast++ {
		c := s[fast]
		cnt[c]++
		for ; cnt[c] > 1; slow++ {
			cnt[s[slow]]--
		}
		if fast-slow+1 > ans {
			ans = fast - slow + 1
		}
	}
	return ans
}
