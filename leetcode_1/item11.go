package leetcode_1

func maxArea(height []int) int {
	l, r, ans := 0, len(height)-1, 0
	for l < r {
		var h int
		if height[l] > height[r] {
			h = height[r]
			r--
		} else {
			h = height[l]
			l++
		}
		w := h * (r - l + 1)
		if w > ans {
			ans = w
		}
	}
	return ans
}
