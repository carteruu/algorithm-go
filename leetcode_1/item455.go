package leetcode_1

import "sort"

/**
贪心
*/
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			//当前小孩能被满足
			i++
		}
		//饼干被吃了或饼干小了，跳过当前饼干
		j++
	}
	return i
}
func findContentChildren_1(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	sIdx := 0
	ans := 0
	for _, gg := range g {
		for sIdx < len(s) && gg > s[sIdx] {
			sIdx++
		}
		if sIdx == len(s) {
			break
		}
		sIdx++
		ans++
	}
	return ans
}
