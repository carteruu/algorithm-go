package leetcode_2

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans, i := 0, -1
	for _, c := range g {
		for i < len(s)-1 {
			i++
			if s[i] < c {
				continue
			} else {
				ans++
				break
			}
		}
	}
	return ans
}
