package leet

import "algorithm/pkg"

func maxNumberOfBalloons(text string) int {
	cnt := make([]int, 26)
	for _, c := range text {
		cnt[c-'a']++
	}
	ans := cnt['a'-'a']
	ans = pkg.Min(ans, cnt['b'-'a'])
	ans = pkg.Min(ans, cnt['l'-'a']/2)
	ans = pkg.Min(ans, cnt['o'-'a']/2)
	ans = pkg.Min(ans, cnt['n'-'a'])
	return ans
}
