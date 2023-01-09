package leetcode_2

import "math"

func beautySum(s string) int {
	var ans int
	for i := range s {
		var cnts [26]int
		maxFreq := 0
		for _, c := range s[i:] {
			cnts[c-'a']++
			maxFreq = maxInt(maxFreq, cnts[c-'a'])
			minFreq := math.MaxInt
			for _, cnt := range cnts {
				if cnt > 0 {
					minFreq = minInt(minFreq, cnt)
				}
			}
			ans += maxFreq - minFreq
		}
	}
	return ans
}
