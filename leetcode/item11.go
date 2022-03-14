package leetcode

import "algorithm/pkg"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		ans = pkg.Max(ans, (r-l)*min(height[l], height[r]))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return ans
}
func maxArea1(height []int) int {
	ans := 0
	for i, val := range height {
		for j := i + 1; j < len(height); j++ {
			ans = pkg.Max(ans, (j-i)*min(val, height[j]))
		}
	}
	return ans
}
